package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/ui"
)

type Sample struct {
	*ui.Box
	src *Sprites
}

func NewSample(src *Sprites) *Sample {
	s := &Sample{
		ui.NewBox(64*6, 144+64*3, nil),
		src,
	}
	s.Self = s
	return s
}

func (s *Sample) Update() {
	now := ui.Now()
	if now%30 != 0 {
		return
	}
	s.Dirty()
	log.Println("Sample.Update")

}

var loopCycle = [4]int{1, 2, 1, 0}

func (s *Sample) Reflesh() {
	cycle := loopCycle[(ui.Now()/30)%4]
	log.Println("Sample.Reflesh")
	// BG
	for j := 0; j < 6; j++ {
		for i := 0; i < 6; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*64), float64(j*64))
			s.Image.DrawImage(g.ImageManager.Bg, op)
		}
	}
	// Face
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	s.Image.DrawImage(s.src.Face, op)
	// TV
	for i := 0; i < 4; i++ {
		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((3+i)*48), 2*48)
		r := image.Rect(cycle*48, i*48, (cycle+1)*48, (i+1)*48)
		s.Image.DrawImage(s.src.Tv.SubImage(r).(*ebiten.Image), op)
	}
	// TVD
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64((3+4)*48), 2*48)
	r := image.Rect(cycle*48, 0, (cycle+1)*48, 48)
	s.Image.DrawImage(s.src.Tvd.SubImage(r).(*ebiten.Image), op)
	// SV
	for i := 0; i < 18; i++ {
		sx, sy := i/6, i%6
		r := image.Rect((sx*3+cycle)*64, sy*64, (sx*3+cycle+1)*64, (sy+1)*64)
		op = &ebiten.DrawImageOptions{}
		tx, ty := i/3, i%3
		op.GeoM.Translate(float64(tx*64), float64(ty*64+144))
		s.Image.DrawImage(s.src.Sv.SubImage(r).(*ebiten.Image), op)
	}
}
