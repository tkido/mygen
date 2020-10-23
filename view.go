package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type View struct {
	Bg       *ebiten.Image
	Face     *ebiten.Image
	PartMenu *ebiten.Image
}

func NewView() View {
	Bg, _, _ := ebitenutil.NewImageFromFile("system/background.png", ebiten.FilterDefault)
	Face, _ := ebiten.NewImage(144, 144, ebiten.FilterDefault)
	PartMenu, _ := ebiten.NewImage(64*64, 64, ebiten.FilterDefault)
	return View{
		Bg:       Bg,
		Face:     Face,
		PartMenu: PartMenu,
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
	screen.DrawImage(v.PartMenu, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 640)
	screen.DrawImage(v.Face, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1000, 0)
	screen.DrawImage(g.ImageManager.Gradient, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(g.MainMenu.Canvas, op)
}
