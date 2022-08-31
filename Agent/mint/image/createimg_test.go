package image

import (
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"unicode"
)

func rotate90(m image.Image) image.Image {
	rotate90 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			//  设置像素点
			rotate90.Set(m.Bounds().Max.Y-x, y, m.At(y, x))
		}
	}
	return rotate90
}

// 旋转180度
func rotate180(m image.Image) image.Image {
	rotate180 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dx(), m.Bounds().Dy()))
	// 矩阵旋转
	for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
		for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
			//  设置像素点
			rotate180.Set(m.Bounds().Max.X-x, m.Bounds().Max.Y-y, m.At(x, y))
		}
	}
	return rotate180
}

// 旋转270度
func rotate270(m image.Image) image.Image {
	rotate270 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			// 设置像素点
			rotate270.Set(x, m.Bounds().Max.X-y, m.At(y, x))
		}
	}
	return rotate270

}
func writeLabel(img image.Image, label string, x, y int, fontColor color.Color, size float64, fontPath string) (image.Image, error) {
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

func CreateRootImage(cardColor string, createImg string) {

}

func UpperNumber(name string) int {
	count := 0
	for _, s := range name {
		if unicode.IsUpper(s) {
			count++
		}
	}
	return count
}
func Test(t *testing.T) {
	src, err := os.Open("./cc.png")
	if err != nil {
		log.Println(err)
		return
	}
	img, err := png.Decode(src)
	//img, err := gif.Decode(src)
	if err != nil {
		log.Println(err)
		return
	}
	//image := rotate90(img)
	// cc 270x270
	label := "Asimov123.SSSSccccccccccccccc" // 12
	//label = strings.ToLower(label)       //10 10
	count := UpperNumber(label)
	labellen := len(label)*10 + count*2
	x := (270 - labellen) / 2
	y := 130
	y1 := 0
	if x < 0 {
		sp := len(label) / 2
		label1 := label[0:sp]
		label2 := label[sp:]
		labellen1 := len(label1)*10 + count*2
		labellen2 := len(label2)*10 + count*2
		x1 := (270 - labellen1) / 2
		x2 := (270 - labellen2) / 2
		x = (x1 + x2) / 2
		y1 = y + 20
		log.Println(label1, label2)
		writeimage, err := writeLabel(img, label1, x, y, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		writeimage, err = writeLabel(writeimage, label2, x, y1, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		if err != nil {
			log.Println(err)
			return
		}
		// outimage := rotate270(writeimage)
		f, err := os.Create("test.png")
		//f, err := os.Create("test.gif")
		if err != nil {
			log.Println(err)
			return
		}
		defer func() { _ = f.Close() }()

		err = png.Encode(f, writeimage)
		//err = gif.Encode(f, outimage, nil)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		writeimage, err := writeLabel(img, label, x, y, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		//writeimage, err = writeLabel(writeimage, "XYZ", 0, 160, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		if err != nil {
			log.Println(err)
			return
		}
		// outimage := rotate270(writeimage)
		f, err := os.Create("test.png")
		//f, err := os.Create("test.gif")
		if err != nil {
			log.Println(err)
			return
		}
		defer func() { _ = f.Close() }()

		err = png.Encode(f, writeimage)
		//err = gif.Encode(f, outimage, nil)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
