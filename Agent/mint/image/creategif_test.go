package image

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
	"sync"
	"testing"
)

func isInPalette(p color.Palette, c color.Color) int {
	ret := -1
	for i, v := range p {
		if v == c {
			return i
		}
	}
	return ret
}

// 获取每张图片的调色板
func getPalette(m image.Image) color.Palette {
	p := color.Palette{color.RGBA{0x00, 0x00, 0x00, 0x00}}
	p9 := color.Palette(palette.Plan9)
	// Bounds返回At可以返回非零颜色的域
	b := m.Bounds()
	black := false
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			// At返回像素At (x, y)的颜色。
			c := m.At(x, y)
			// Convert返回欧氏R,G,B空间中最接近c的调色板颜色。
			cc := p9.Convert(c)
			if cc == p9[0] {
				black = true
			}
			if isInPalette(p, cc) == -1 {
				p = append(p, cc)
			}
		}
	}
	if len(p) < 256 && black == true {
		p[0] = color.RGBA{0x00, 0x00, 0x00, 0x00} // transparent
		p = append(p, p9[0])
	}
	return p
}

// pool

type WaitGroup struct {
	workChan chan int
	// WaitGroup等待一组goroutines完成。
	//主goroutine调用Add来设置
	// goroutines等待。然后是每一个goroutine
	//在完成时运行并调用Done。与此同时，
	//等待可以用来阻塞，直到所有goroutines完成
	wg sync.WaitGroup
}

func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}
func (ap *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		// Add将增量添加到WaitGroup计数器，增量可以是负数。
		// 如果计数器变为0，所有在Wait阻塞的goroutines将被释放。
		ap.wg.Add(1)
	}
}
func (ap *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
	}
	ap.wg.Done()
}

func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}

func CreateGif(label string, cardColor string, savePath string) {
	var disposals []byte
	var images []*image.Paletted
	var delays []int
	imageMap := make(map[int]*image.Paletted, 250)
	work := NewPool(100)
	for i := 1; i < 251; i++ {
		work.Add(1)
		go func(i int, wg *WaitGroup) {
			defer wg.Done()
			var src string
			src = fmt.Sprintf("./%s/图层 %d.png", cardColor, i)
			file, err := os.Open(src)
			if err != nil {
				fmt.Println(err)
			}
			defer func() { _ = file.Close() }()
			g, err := png.Decode(file)
			// 写入文字
			writeimage, err := writeLabel(g, label, 32, 32, color.RGBA{234, 178, 16, 255}, 30, "./PingFang Heavy.ttf")
			if err != nil {
				fmt.Println(err)
			}
			cp := getPalette(writeimage)
			// cp := append(palette.WebSafe, color.Transparent)
			// 透明图片需要设置
			disposals = append(disposals, gif.DisposalBackground)
			// 返回一个给定宽度，高度和面板。
			p := image.NewPaletted(image.Rect(0, 0, 600, 600), cp)
			draw.Draw(p, p.Bounds(), writeimage, image.ZP, draw.Src)
			//images = append(images, p)
			imageMap[i] = p
			delays = append(delays, 4)
			//time.Sleep(time.Second * 1)
		}(i, work)
	}

	log.Println("waiting...")
	work.Wait()
	log.Println("done")
	for i := 1; i < 251; i++ {
		images = append(images, imageMap[i])
	}
	g := &gif.GIF{
		Image: images,
		Delay: delays,
		//LoopCount为0表示永远循环
		LoopCount: 0,
		//处理是连续的处理方法，每帧一个
		Disposal: disposals,
	}
	f, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = f.Close() }()
	gif.EncodeAll(f, g)
}

func TestG(t *testing.T) {
	CreateGif("NO.2672", "color", "./color_2672.gif")
}
