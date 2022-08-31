package api

import (
	"fmt"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func writeLabel(img image.Image, label string, x, y int, fontColor color.Color, size float64, fontPath string) (image.Image, error) {
	// 原图边界
	bound := img.Bounds()
	// 创建一个新的图片
	rgba := image.NewRGBA(image.Rect(0, 0, bound.Dx(), bound.Dy()))
	// 读取字体
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return rgba, err
	}
	myFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return rgba, err
	}

	draw.Draw(rgba, rgba.Bounds(), img, bound.Min, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(myFont)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	uni := image.NewUniform(fontColor)
	c.SetSrc(uni)
	c.SetHinting(font.HintingNone)

	// 在指定的位置显示
	pt := freetype.Pt(x, y+int(c.PointToFixed(size)>>6))
	if _, err := c.DrawString(label, pt); err != nil {
		return rgba, err
	}

	return rgba, nil
}

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
	// pp := color.Palette{color.RGBA{0x00, 0x00, 0x00, 0x00}}
	//  Plan9是一个256色的调色板，它划分了24位RGB空间
	//p9 := color.Palette(palette.Plan9)
	// Bounds返回At可以返回非零颜色的域
	b := m.Bounds()
	// black := false
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			// At返回像素At (x, y)的颜色。
			c := m.At(x, y)
			// Convert返回欧氏R,G,B空间中最接近c的调色板颜色。
			// cc := p9.Convert(c)
			//if cc == p9[0] {
			//	black = true
			//}
			// 不同颜色
			if isInPalette(p, c) == -1 {
				//p = append(p, cc)
				p = append(p, c)
			}
		}
	}
	// 256色
	//if len(p) < 256 && black == true {
	//	//透明的
	//	p[0] = color.RGBA{0x00, 0x00, 0x00, 0x00}
	//	p = append(p, p9[0])
	//	pp = append(pp, p9[0])
	//}
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
	//imageMap := make(map[int]*image.Paletted, 500)
	var imageMap sync.Map
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
			// 写入文字 600x600 : 32x32   gold   234, 178, 16, 255    green  720x720:35x35  rgb :28, 223, 143
			color_rgba := color.RGBA{28, 223, 143, 255}
			if cardColor == "gold" {
				color_rgba = color.RGBA{234, 178, 16, 255}
			} else if cardColor == "color" {
				color_rgba = color.RGBA{187, 46, 115, 255}
			}
			writeimage, err := writeLabel(g, label, 35, 35, color_rgba, 30, "./PingFang Heavy.ttf")
			if err != nil {
				fmt.Println(err)
			}
			cp := getPalette(g)
			// 返回一个给定宽度，高度和面板。
			p := image.NewPaletted(image.Rect(0, 0, 720, 720), cp)
			draw.Draw(p, p.Bounds(), writeimage, image.ZP, draw.Src)
			//images = append(images, p)
			//imageMap[i] = p
			imageMap.Store(i, p)
			//time.Sleep(time.Second * 1)
		}(i, work)
	}

	log.Println("waiting...")
	work.Wait()
	log.Println("done")
	for i := 1; i < 251; i++ {
		img, _ := imageMap.Load(i)
		if v, ok := img.(*image.Paletted); ok {
			images = append(images, v)
		} else {
		}
		delays = append(delays, 4)
		// 透明图片:DisposalBackground（背景为纯色时设置 不然有重影） 覆盖:DisposalPrevious  不处理:DisposalNone
		disposals = append(disposals, gif.DisposalNone)
	}
	g := &gif.GIF{
		Image: images,
		Delay: delays,
		//LoopCount为0表示永远循环
		LoopCount: 0,
		//gif背景颜色和gif动画处理方法
		Disposal: disposals,
	}
	f, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = f.Close() }()
	encode_err := gif.EncodeAll(f, g)
	if encode_err != nil {
		log.Println(fmt.Sprintf("%s encodegif err:%s", label, encode_err))
	}

}
