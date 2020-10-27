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

/*
Sprite_Actor.MOTIONS = {
    walk:     { index: 0,  loop: true  },
    wait:     { index: 1,  loop: true  },
    chant:    { index: 2,  loop: true  },
    guard:    { index: 3,  loop: true  },
    damage:   { index: 4,  loop: false },
    evade:    { index: 5,  loop: false },
    thrust:   { index: 6,  loop: false },
    swing:    { index: 7,  loop: false },
    missile:  { index: 8,  loop: false },
    skill:    { index: 9,  loop: false },
    spell:    { index: 10, loop: false },
    item:     { index: 11, loop: false },
    escape:   { index: 12, loop: true  },
    victory:  { index: 13, loop: true  },
    dying:    { index: 14, loop: true  },
    abnormal: { index: 15, loop: true  },
    sleep:    { index: 16, loop: true  },
    dead:     { index: 17, loop: true  }
};
*/
