package font

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type FontType int

//go:generate stringer -type=FontType
const (
	Regular FontType = iota
	Bold
	Pixel
)

type FontSize int

//go:generate stringer -type=FontSize
const (
	XSmall FontSize = 12
	Small           = 16
	Medium          = 24
	Large           = 36
	XLarge          = 48
)

type FontManager struct {
	Fonts map[FontType]FontData
}

func NewFontManager() FontManager {
	return FontManager{
		Fonts: map[FontType]FontData{},
	}
}

type FontData struct {
	Font  *truetype.Font
	Faces map[FontSize]font.Face
}

func (fm *FontManager) Face(fontType FontType, size FontSize) font.Face {
	fd, ok := fm.Fonts[fontType]
	if !ok {
		log.Fatalf("FontManager.Face: unknown FontType %d", fontType)
	}
	if face, ok := fd.Faces[size]; ok {
		return face
	}
	face := truetype.NewFace(fd.Font, &truetype.Options{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	fd.Faces[size] = face
	return face
}

func (fm *FontManager) RegisterFont(fontType FontType, path string) {
	ttf, err := os.Open(path)
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
	fm.Fonts[fontType] = FontData{tt, map[FontSize]font.Face{}}
}
