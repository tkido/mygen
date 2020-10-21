package main

import (
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var imgCache map[string]*ebiten.Image

func init() {
	imgCache = map[string]*ebiten.Image{}
}

func loadImage(path string) *ebiten.Image {
	img, ok := imgCache[path]
	if !ok {
		var err error
		img, _, err = ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		imgCache[path] = img
	}
	return img
}

func saveImage(path string, img *ebiten.Image) {
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

func filterImage(img *ebiten.Image) *ebiten.Image {
	w, h := img.Size()
	newImage, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			oc := img.At(x, y)
			index := colorToGradientIndex(oc)
			if index == -1 {
				continue
			}
			nc := imgGrad.At(index, 42*4)
			// 透明度は元のものを維持する
			oc1 := color.RGBAModel.Convert(oc).(color.RGBA)
			nc1 := color.RGBAModel.Convert(nc).(color.RGBA)
			nc1.A = oc1.A

			newImage.Set(x, y, nc1)
		}
	}
	return newImage
}

func colorToGradientIndex(col color.Color) int {
	c := color.RGBAModel.Convert(col).(color.RGBA)
	// transparent not convert
	if c.A == 0 {
		return -1
	}
	fr, fg, fb := float64(c.R), float64(c.G), float64(c.B)
	min := 255.0
	max := 0.0
	min = math.Min(min, fr)
	max = math.Max(max, fr)
	min = math.Min(min, fg)
	max = math.Max(max, fg)
	min = math.Min(min, fb)
	max = math.Max(max, fb)
	rst := math.Floor(255.0 - (min+max)/2)

	// rst := math.Floor(255.0 - (fr+fg+fb)/3)

	// rst := math.Floor(255.0 - math.Pow(fr*fg*fb, 1/3))

	return int(rst)
}
