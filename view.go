package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type View struct {
	Bg   *ebiten.Image
	Face *ebiten.Image
	Menu *ebiten.Image
}

func NewView() View {
	Bg, _, _ := ebitenutil.NewImageFromFile("system/background.png", ebiten.FilterDefault)
	Face, _ := ebiten.NewImage(144, 144, ebiten.FilterDefault)
	Menu, _ := ebiten.NewImage(64*64, 64, ebiten.FilterDefault)
	return View{
		Bg,
		Face,
		Menu,
	}
}

func (v View) Draw(screen *ebiten.Image) {
	for i := 3; i < 30; i++ {
		for j := 0; j < 17; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i)*64, float64(j)*64)
			screen.DrawImage(v.Bg, op)
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(200, 0)
	screen.DrawImage(v.Menu, op)
	op.GeoM.Translate(0, 64)
	screen.DrawImage(v.Face, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1000, 0)
	screen.DrawImage(g.ImageManager.Gradient, op)

}
