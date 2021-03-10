package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type CaseExchange struct {
	Name  string
	Color color.Color
}

func DoExchangeColor() {
	cases := []CaseExchange{
		{"Scar", color.RGBA{56, 59, 59, 255}},
		{"Soil", color.RGBA{79, 65, 60, 255}},
		{"Emotion", color.RGBA{87, 87, 85, 255}},
		{"Bitten", color.RGBA{170, 175, 175, 255}},
	}
	for _, c := range cases {
		pattern := filepath.Join(
			rootPath,
			"*",
			"*",
			fmt.Sprintf("*_%s_*_c.png", c.Name),
		)
		paths, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}
		for _, path := range paths {
			fmt.Println(path)
			img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
			if err != nil {
				log.Println(err)
			}
			img = ExchangeColor(img, color.RGBA{0, 146, 150, 255}, c.Color)
			f, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			err = png.Encode(f, img)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func ExchangeColor(img *ebiten.Image, from, to color.Color) *ebiten.Image {
	w, h := img.Size()
	newImage, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if from == img.At(x, y) {
				newImage.Set(x, y, to)
			}
		}
	}
	return newImage
}
