package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/tkido/mygen/font"
	"github.com/tkido/mygen/part"
)

type MainMenu struct {
	MenuBase
	Data []string
}

func NewMainMenu(w, h, col, row int) *MainMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursorImg, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)

	menu := &MainMenu{
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
	menu.Update()
	return menu
}

func (m *MainMenu) SetCursor(index int) {
	m.MenuBase.SetCursor(index)
	const offset = 2 // num of appended menu item
	if m.Cursor >= offset {
		g.PartMenu.Part = part.Types[m.Cursor-offset]
		g.PartMenu.Update()
		g.PaletteMenu.Update()
	}
}

func (m *MainMenu) Update() {
	m.Data = []string{}
	m.Data = append(m.Data, fmt.Sprintf("%04d", g.Character.Id))
	m.Data = append(m.Data, g.Character.Base.String())
	for _, pt := range part.Types {
		m.Data = append(m.Data, pt.String())
	}
	m.Limit = len(m.Data) - 1
}

func (m *MainMenu) Reflesh() {
	log.Println("Menu.Reflesh")
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
