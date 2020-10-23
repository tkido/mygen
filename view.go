package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type View struct {
	Bg   *ebiten.Image
	Face *ebiten.Image
}

func NewView() View {
	Bg, _, _ := ebitenutil.NewImageFromFile("system/background.png", ebiten.FilterDefault)
	Face, _ := ebiten.NewImage(144, 144, ebiten.FilterDefault)
	return View{
		Bg:   Bg,
		Face: Face,
	}
}

func (v View) Draw(screen *ebiten.Image) {
	for i := 0; i < 30; i++ {
		for j := 0; j < 17; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i)*64, float64(j)*64)
			screen.DrawImage(v.Bg, op)
		}
	}

	var op *ebiten.DrawImageOptions

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 0)
	screen.DrawImage(g.PartMenu.Canvas, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 64*4)
	screen.DrawImage(v.Face, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1000, 0)
	screen.DrawImage(g.ImageManager.Gradient, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(g.MainMenu.Canvas, op)
}
