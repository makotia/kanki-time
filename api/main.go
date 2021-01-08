package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func calcMargin(r rune, beforeHan bool, first bool) int {
	if beforeHan || first {
		if unicode.In(r, unicode.Han) {
			return 200
		} else if first {
			return 160
		} else {
			return 150
		}
	} else {
		return 160
	}
}

// Usage: ./main.out [font.ttf] [dist.png] [text]
func main() {
	ftBinary, err := ioutil.ReadFile(os.Args[1])
	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	opt := truetype.Options{
		Size:              200,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	imageWidth := 1920
	imageHeight := 1080
	text := os.Args[3]

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	dr := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{247, 0, 239, 255}),
		Face: truetype.NewFace(ft, &opt),
	}

	fontSizePx := int(opt.Size)

	textArr := strings.Split(text, "\n")

	y := imageHeight/2 + (fontSizePx * 3 / 4) - (fontSizePx * len(textArr) / 2)
	dy := int(opt.Size)

	for _, t := range textArr {
		textWidth := 0
		beforeHan := false
		for i, str := range t {
			textWidth += calcMargin(str, beforeHan, i == 0)
			beforeHan = unicode.In(str, unicode.Han)
		}

		beforeHan = false
		x := imageWidth/2 - (textWidth / 2)

		for i, str := range t {
			dr.Dot = fixed.P(x, y)
			dr.DrawString(string(str))
			x += calcMargin(str, beforeHan, i == 0)
			beforeHan = unicode.In(str, unicode.Han)
		}
		y += dy
	}

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
