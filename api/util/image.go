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

	"github.com/golang/freetype/truetype"
	"github.com/ktnyt/go-moji"
	"github.com/makotia/kanki-time/api/config"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var ft *truetype.Font

func init() {
	ftBinary, err := ioutil.ReadFile(config.GetConfig().Server.TTFPath)
	ft, err = truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// SaveImage is save png to local
func SaveImage(img *image.RGBA) (id string, err error) {
	var (
		buf  bytes.Buffer
		file *os.File
	)
	id = genToken()
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

// GenImage is generate text
func GenImage(text string, imageType string) (img *image.RGBA, err error) {
	var (
		defaultFontSize float64 = 200.0
		fontSize                = defaultFontSize
		imageWidth      int     = 1920
		imageHeight     int     = 1005
		isTimeTemplate  bool    = false
		opt             truetype.Options
		dr              *font.Drawer
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
		return nil, errors.New("imageType is not valid")
	}
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
		draw.Draw(img, image.Rect(50, 50, 1870, 250), image.NewUniform(color.RGBA{140, 168, 147, 255}), image.ZP, draw.Src)
		draw.Draw(img, image.Rect(50, 300, 1870, 958), image.NewUniform(color.RGBA{140, 168, 147, 255}), image.ZP, draw.Src)
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
			textWidth int = 0
		)

		t := moji.Convert(t, moji.HE, moji.ZE)

		textWidth = len(t) * int(fontSize) / 3

		x := imageWidth/2 - (textWidth / 2)

		for _, str := range t {
			dr.Dot = fixed.P(x, y)
			dr.DrawString(string(str))
			x += int(fontSize)
		}
		y += dy
	}
	return img, err
}
