package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/status"
	"github.com/tkido/mygen/ui"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tkido/mygen/gradient"
)

type CachedImage struct {
	Image *ebiten.Image
	Frame int
}

type ImageManager struct {
	Bg       *ebiten.Image
	Gradient *ebiten.Image
	Cache    map[string]*CachedImage
}

func NewImageManager() ImageManager {
	Bg, _, _ := ebitenutil.NewImageFromFile("system/background.png", ebiten.FilterDefault)
	Gradient, _, _ := ebitenutil.NewImageFromFile("generator/gradients.png", ebiten.FilterDefault)
	return ImageManager{
		Bg,
		Gradient,
		map[string]*CachedImage{},
	}
}

func (m *ImageManager) Gc() {
	const limit = 480
	if len(m.Cache) < limit {
		return
	}
	log.Printf("ImageManager Gc Start")
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Alloc : %v\n", humanize.Bytes(mem.Alloc))
	fmt.Printf("Cache size: %d\n", len(m.Cache))
	const sampleSize = 16
	i := 0
	sum := 0
	for _, v := range m.Cache {
		if i == sampleSize {
			break
		}
		i++
		sum += v.Frame
	}
	floor := sum / sampleSize
	for k, v := range m.Cache {
		if v.Frame < floor {
			delete(m.Cache, k)
		}
	}
	fmt.Printf("Alloc : %v\n", humanize.Bytes(mem.Alloc))
	fmt.Printf("Cache size: %d\n", len(m.Cache))
}

func (m *ImageManager) LoadImage(path string) *ebiten.Image {
	if cached, ok := m.Cache[path]; ok {
		cached.Frame = ui.Now()
		return cached.Image
	}
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		log.Println(err)
	}
	m.Cache[path] = &CachedImage{img, ui.Now()}
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

func (m *ImageManager) FilterImage2(img, mask *ebiten.Image) *ebiten.Image {
	w, h := img.Size()
	newImage, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			oc := img.At(x, y)
			mc := mask.At(x, y)
			pt, ok := palette.MaskPaletteMap[mc]
			if !ok {
				newImage.Set(x, y, oc)
				continue
			}
			index := m.ColorToGradientIndex(oc)
			if index == -1 {
				continue
			}
			row, ok := g.Character.StatusMap[status.Human].Colors[pt]
			if !ok || row == gradient.Null {
				newImage.Set(x, y, oc)
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
	return int(rst)
}
