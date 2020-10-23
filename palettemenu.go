package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/tkido/mygen/font"
	"github.com/tkido/mygen/palette"
)

type PaletteMenu struct {
	MenuBase
	Data []string
}

func NewPaletteMenu(w, h, col, row int) *PaletteMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursorImg, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)

	menu := &PaletteMenu{
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
		Data: []string{},
	}
	menu.Self = menu
	return menu
}

func (m *PaletteMenu) SetCursor(index int) {
	m.MenuBase.SetCursor(index)
	// const offset = 2 // num of appended menu item
	// if m.Cursor >= offset {
	// 	g.PartMenu.Part = part.Types[m.Cursor-offset]
	// 	g.PartMenu.Update()
	// }
	m.Reflesh()
	g.Logic.UpdateFace()
}

func (m *PaletteMenu) Update() {
	m.Data = []string{}
	pt := g.PartMenu.Part
	ps := palette.Map[pt]
	for _, p := range ps {
		m.Data = append(m.Data, p.String())
	}
	m.Limit = len(m.Data) - 1
	m.SetCursor(0)
}

func (m *PaletteMenu) Reflesh() {
	log.Println("PaletteMenu.Reflesh")
	f := g.FontManager.Face(font.Regular, font.XSmall)
	fHeight := f.Metrics().Height.Ceil()
	m.Canvas.Fill(color.White)

	for i, s := range m.Data {
		x := i % m.Col
		y := i / m.Col
		text.Draw(m.Canvas, s, f, x*m.W, y*m.H+fHeight, color.Black)
		if m.Cursor == i {
			m.CursorImg.Fill(g.View.GetFocusColor(m))
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
			m.Canvas.DrawImage(m.CursorImg, op)
		}
	}
}
