package util

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/golang/freetype/truetype"
	"github.com/makotia/kanki-time/api/config"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	ft *truetype.Font
)

func init() {
	ftBinary, err := ioutil.ReadFile(config.GetConfig().Server.TTFPath)
	ft, err = truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func calcMargin(r rune, beforeHan bool, first bool) float64 {
	if beforeHan || first {
		if unicode.In(r, unicode.Han) {
			return 1
		} else if first {
			return 0.8
		} else {
			return 0.75
		}
	} else {
		return 0.8
	}
}

// GenImage is generate text
func GenImage(text string, imageType string) (id string, err error) {
	var (
		defaultFontSize float64 = 200.0
		fontSize                = defaultFontSize
		imageWidth      int     = 1920
		imageHeight     int     = 1080
		isTimeTemplate  bool    = false
		opt             truetype.Options
		img             *image.RGBA
		dr              *font.Drawer
		buf             bytes.Buffer
		file            *os.File
	)

	switch imageType {
	case "square":
		imageHeight = 1000
		imageWidth = 1000
	case "time":
		isTimeTemplate = true
	case "slide":
		break
	default:
		return "", errors.New("imageType is not valid")
	}

	id = genToken()
	opt = truetype.Options{
		Size:              defaultFontSize,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	img = image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

	if isTimeTemplate {
		draw.Draw(img, img.Bounds(), image.NewUniform(color.RGBA{14, 122, 0, 255}), image.ZP, draw.Src)
		draw.Draw(img, image.Rect(100, 50, 1820, 250), image.NewUniform(color.RGBA{140, 168, 147, 255}), image.ZP, draw.Src)
		draw.Draw(img, image.Rect(100, 300, 1820, 1030), image.NewUniform(color.RGBA{140, 168, 147, 255}), image.ZP, draw.Src)
		imageHeight += 300
	} else {
		draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)
	}

	dr = &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{247, 0, 239, 255}),
		Face: truetype.NewFace(ft, &opt),
	}

	fontSizePx := int(fontSize)

	textArr := strings.Split(text, "\n")

	y := imageHeight/2 + (fontSizePx * 3 / 4) - (fontSizePx * len(textArr) / 2)
	dy := int(fontSize)

	for _, t := range textArr {
		var (
			textWidth int  = 0
			beforeHan bool = false
		)

		for i, str := range t {
			textWidth += int(calcMargin(str, beforeHan, i == 0) * fontSize)
			beforeHan = unicode.In(str, unicode.Han)
		}

		beforeHan = false
		x := imageWidth/2 - (textWidth / 2)

		for i, str := range t {
			dr.Dot = fixed.P(x, y)
			dr.DrawString(string(str))
			x += int(calcMargin(str, beforeHan, i == 0) * fontSize)
			beforeHan = unicode.In(str, unicode.Han)
		}
		y += dy
	}
	if err = png.Encode(&buf, img); err != nil {
		return "", err
	}
	file, err = os.Create(path.Join(config.GetConfig().Server.StaticDir, id+".png"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	file.Write(buf.Bytes())
	return id, err
}
