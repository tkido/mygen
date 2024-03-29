package main

import (
	"fmt"
	"image"
	"log"
	"regexp"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/mode"
	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
	"github.com/tkido/mygen/status"
	"github.com/tkido/mygen/ui"
)

type Sprites struct {
	*ui.Box
	Status status.Type
	Mode   mode.Type
	Face   *ebiten.Image
	Tv     *ebiten.Image
	Tvd    *ebiten.Image
	Sv     *ebiten.Image
}

func NewSprites() *Sprites {
	face, _ := ebiten.NewImage(144, 144, ebiten.FilterDefault)
	tv, _ := ebiten.NewImage(48*3, 48*4, ebiten.FilterDefault)
	tvd, _ := ebiten.NewImage(48*3, 48*1, ebiten.FilterDefault)
	sv, _ := ebiten.NewImage(64*9, 64*6, ebiten.FilterDefault)
	s := &Sprites{
		ui.NewBox(64*12, 64*6, nil),
		status.Hum,
		mode.Normal,
		face,
		tv,
		tvd,
		sv,
	}
	s.Self = s
	return s
}

func (s *Sprites) Reflesh() {
	g.ImageManager.Gc()
	log.Println("Sprites.Reflesh")

	for j := 0; j < 6; j++ {
		for i := 0; i < 12; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*64), float64(j*64))
			s.Image.DrawImage(g.ImageManager.Bg, op)
		}
	}

	for _, st := range sprite.Types {
		s.reflesh(st)
	}
}

func (s *Sprites) reflesh(st sprite.Type) {
	if st == sprite.Face {
		s.refleshFace()
		return
	}
	if g.StatusMenu.Status > status.ZombN {
		return
	}
	var target *ebiten.Image
	op := &ebiten.DrawImageOptions{}
	var clip image.Rectangle
	switch st {
	case sprite.Tv:
		target = s.Tv
		op.GeoM.Translate(0, 144)
		if s.Mode == mode.Simple {
			clip = image.Rect(0, 0, 48, 48)
		}
	case sprite.Tvd:
		target = s.Tvd
		op.GeoM.Translate(0, 144+48*4)
		if s.Mode == mode.Simple {
			clip = image.Rect(0, 0, 48, 48)
		}
	case sprite.Sv:
		target = s.Sv
		op.GeoM.Translate(144, 0)
		if s.Mode == mode.Simple {
			clip = image.Rect(0, 0, 64, 64)
		}
	}
	target.Clear()

	layers := layer.Map[st]
	for i := len(layers) - 1; 0 <= i; i-- { // reverse
		lay := layers[i]
		label := "01"
		if pt, ok := g.PartManager.LayerPartMap[lay]; ok {
			if list, ok := g.VariationManager.Map[g.Character.Base][pt]; ok {
				if index, ok := g.Character.StatusMap[g.StatusMenu.Status].Parts[pt]; ok {
					if index == part.Null {
						continue
					}
					if index < part.Index(len(list)) {
						label = list[index].label
					}
				}
			}
		}
		files := g.PartManager.Get(st, g.Character.Base, lay, label)
		for i := len(files) - 1; 0 <= i; i-- { // reverse
			file := files[i]
			fmt.Println(file)
			imgSrc := g.ImageManager.LoadImage(file)
			if s.Mode == mode.Simple {
				imgSrc = imgSrc.SubImage(clip).(*ebiten.Image)
			}
			maskFileName := file[:len(file)-4] + `_c.png`
			// mask info found
			if imgMask := g.ImageManager.LoadImage(maskFileName); imgMask != nil {
				if s.Mode == mode.Simple {
					imgMask = imgMask.SubImage(clip).(*ebiten.Image)
				}
				imgSrc = g.ImageManager.FilterImage2(imgSrc, imgMask)
			}
			op := &ebiten.DrawImageOptions{}
			target.DrawImage(imgSrc, op)
		}
	}

	s.Image.DrawImage(target, op)
}

func (s *Sprites) refleshFace() {
	var reDefaultColor = regexp.MustCompile(`_m(\d{3})`)
	layers := layer.Map[sprite.Face]
	s.Face.Clear()
	for i := len(layers) - 1; 0 <= i; i-- { // reverse
		lay := layers[i]
		// log.Printf("Set %s...\n", lay)
		label := "01"
		if pt, ok := g.PartManager.LayerPartMap[lay]; ok {
			if list, ok := g.VariationManager.Map[g.Character.Base][pt]; ok {
				if index, ok := g.Character.StatusMap[g.StatusMenu.Status].Parts[pt]; ok {
					if index == part.Null {
						continue
					}
					if index < part.Index(len(list)) {
						label = list[index].label
					}
				}
			}
		}

		files := g.PartManager.Get(sprite.Face, g.Character.Base, lay, label)
		for i := len(files) - 1; 0 <= i; i-- { // reverse
			file := files[i]
			// fmt.Println(file)
			imgSrc := g.ImageManager.LoadImage(file)
			// color info found
			if ms := reDefaultColor.FindStringSubmatch(file); len(ms) >= 2 {
				label := ms[1]
				index, err := strconv.Atoi(label)
				if err != nil {
					log.Fatalf("invalid color info")
				}
				p := palette.Type(index)
				row, ok := g.Character.StatusMap[g.StatusMenu.Status].Colors[p]
				if ok && row != -1 {
					imgSrc = g.ImageManager.FilterImage(imgSrc, row)
				}
			}
			op := &ebiten.DrawImageOptions{}
			s.Face.DrawImage(imgSrc, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	s.Image.DrawImage(s.Face, op)
}
