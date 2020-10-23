package main

import (
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tkido/mygen/gradient"
)

type ImageManager struct {
	Gradient *ebiten.Image
	Cache    map[string]*ebiten.Image
}

func NewImageManager() ImageManager {
	Gradient, _, _ := ebitenutil.NewImageFromFile("generator/gradients.png", ebiten.FilterDefault)
	Cache := map[string]*ebiten.Image{}
	return ImageManager{
		Gradient,
		Cache,
	}
}

func (m *ImageManager) LoadImage(path string) *ebiten.Image {
	img, ok := m.Cache[path]
	if !ok {
		var err error
		img, _, err = ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		m.Cache[path] = img
	}
	return img
}

func (m *ImageManager) SaveImage(path string, img *ebiten.Image) {
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

func (m *ImageManager) FilterImage(img *ebiten.Image, row gradient.Row) *ebiten.Image {
	w, h := img.Size()
	newImage, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			oc := img.At(x, y)
			index := m.ColorToGradientIndex(oc)
			if index == -1 {
				continue
			}
			nc := m.Gradient.At(index, int(row)*4)
			// 透明度は元のものを維持する
			oc1 := color.RGBAModel.Convert(oc).(color.RGBA)
			nc1 := color.RGBAModel.Convert(nc).(color.RGBA)
			nc1.A = oc1.A

			newImage.Set(x, y, nc1)
		}
	}
	return newImage
}

func (m *ImageManager) ColorToGradientIndex(col color.Color) int {
	c := color.RGBAModel.Convert(col).(color.RGBA)
	// transparent pixel need not convert
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
