package font

import (
	"io/ioutil"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/tkido/mygen/assets"
	"github.com/tkido/mygen/ui"
)

type FontType int

const (
	Regular ui.FontType = iota
	Bold
	Pixel
)

type FontSize int

const (
	XSmall ui.FontSize = 12
	Small              = 16
	Medium             = 24
	Large              = 36
	XLarge             = 48
)

func init() {
	register(Regular, "mplus-1m-regular.ttf")
	// register(Bold, "mplus-1p-bold.ttf")
	// register(Pixel, "PixelMplus12-Regular.ttf")
}

func register(fontType ui.FontType, path string) {
	ttf, err := assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer ttf.Close()
	bs, err := ioutil.ReadAll(ttf)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(bs)
	if err != nil {
		log.Fatal(err)
	}
	ui.RegisterFont(fontType, tt)
}
