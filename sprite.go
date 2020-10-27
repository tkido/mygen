package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
	"github.com/tkido/mygen/status"
	"github.com/tkido/mygen/ui"
)

type Sprites struct {
	*ui.Box
	Status status.Type
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
		status.Human,
		face,
		tv,
		tvd,
		sv,
	}
	s.Self = s
	return s
}

func (s *Sprites) Reflesh() {
	log.Println("Sprites.Reflesh")

	for j := 0; j < 6; j++ {
		for i := 0; i < 12; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*64), float64(j*64))
			s.Image.DrawImage(g.ImageManager.Bg, op)
		}
	}

	s.refleshFace()
	s.refleshTv()
}
func (s *Sprites) refleshTv() {
	s.Tv.Clear()
	for i := len(layer.TvLayers) - 1; 0 <= i; i-- { // reverse
		lay := layer.TvLayers[i]
		label := "01"
		if pt, ok := g.PartManager.LayerPartMap[lay]; ok {
			if list, ok := g.VariationManager.Map[g.Character.Base][pt]; ok {
				if index, ok := g.Character.StatusMap[s.Status].Parts[pt]; ok {
					if index == part.Null {
						continue
					}
					if index < part.Index(len(list)) {
						label = list[index].label
					}
				}
			}
		}
		files := g.PartManager.Get(sprite.Tv, g.Character.Base, lay, label)
		for i := len(files) - 1; 0 <= i; i-- { // reverse
			file := files[i]
			fmt.Println(file)
			imgSrc := g.ImageManager.LoadImage(file)
			maskFileName := file[:len(file)-4] + `_c.png`
			// mask info found
			if imgMask := g.ImageManager.LoadImage(maskFileName); imgMask != nil {
				fmt.Println(maskFileName)
				imgSrc = g.ImageManager.FilterImage2(imgSrc, imgMask)
			}
			op := &ebiten.DrawImageOptions{}
			s.Tv.DrawImage(imgSrc, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 144)
	s.Image.DrawImage(s.Tv, op)

}

func (s *Sprites) refleshFace() {
	var reDefaultColor = regexp.MustCompile(`_m(\d{3})`)

	s.Face.Clear()
	for i := len(layer.FaceLayers) - 1; 0 <= i; i-- { // reverse
		lay := layer.FaceLayers[i]
		// log.Printf("Set %s...\n", lay)
		label := "01"
		if pt, ok := g.PartManager.LayerPartMap[lay]; ok {
			if list, ok := g.VariationManager.Map[g.Character.Base][pt]; ok {
				if index, ok := g.Character.StatusMap[s.Status].Parts[pt]; ok {
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
				row, ok := g.Character.StatusMap[status.Human].Colors[p]
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
