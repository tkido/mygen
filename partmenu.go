package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/part"
)

type PartMenu struct {
	MenuBase
	Data []*ebiten.Image
}

func NewPartMenu(w, h, col, row int) *PartMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursorImg, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	cursorImg.Fill(color.RGBA{255, 255, 0, 64})

	menu := &PartMenu{
		MenuBase: MenuBase{
			W:         w,
			H:         h,
			Col:       col,
			Row:       row,
			Cursor:    0,
			Canvas:    canvas,
			CursorImg: cursorImg,
			Dirty:     true,
		},
		Data: []*ebiten.Image{},
	}
	menu.Self = menu
	return menu
}

func (m *PartMenu) SetCursor(index int) {
	m.MenuBase.SetCursor(index)
	g.Logic.UpdateFace()
}

func (m *PartMenu) Update() {
	ps := g.VariationManager.Get(g.Character.Base, part.FrontHair)
	m.Limit = len(ps) - 1
	m.Data = []*ebiten.Image{}
	for _, p := range ps {
		m.Data = append(m.Data, g.ImageManager.LoadImage(p))
	}
	m.Reflesh()
}

func (m *PartMenu) Reflesh() {
	log.Println("PartMenu.Reflesh")

	m.Canvas.Fill(color.White)

	for i, img := range m.Data {
		x := i % m.Col
		y := i / m.Col
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
		m.Canvas.DrawImage(img, op)
		if m.Cursor == i {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
			m.Canvas.DrawImage(m.CursorImg, op)
		}
	}
}
