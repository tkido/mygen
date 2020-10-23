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

type MenuBase struct {
	X, Y, W, H, Col, Row int
	Canvas               *ebiten.Image
	Cursor               *ebiten.Image
	Dirty                bool
}

type Menu interface {
	Update()
	Reflesh()
	MoveCursor(dX, dY int)
	SetFocus(index int)
	IsFocused() bool
	GetFocus()
	LoseFocus()
}

type MainMenu struct {
	MenuBase
	Data []string
}

func NewMainMenu(w, h, col, row int) MainMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursor, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	cursor.Fill(color.RGBA{255, 255, 0, 64})

	menu := MainMenu{
		MenuBase: MenuBase{
			X:      0,
			Y:      0,
			W:      w,
			H:      h,
			Col:    col,
			Row:    row,
			Canvas: canvas,
			Cursor: cursor,
			Dirty:  true,
		},
		Data: []string{},
	}
	menu.Update()
	return menu
}

func (m *MainMenu) Update() {
	m.Data = []string{}
	m.Data = append(m.Data, fmt.Sprintf("%04d", g.Character.Id))
	m.Data = append(m.Data, g.Character.Base.String())
	for _, pt := range part.Types {
		m.Data = append(m.Data, pt.String())
	}
	m.Data = append(m.Data, "Load")
	m.Data = append(m.Data, "Save")
	m.Data = append(m.Data, "New")
}

func (m *MainMenu) Reflesh() {
	log.Println("Menu.Reflesh")
	f := g.FontManager.Face(font.Regular, font.XSmall)
	fHeight := f.Metrics().Height.Ceil()

	m.Canvas.Fill(color.Black)

	for i, s := range m.Data {
		x := i % m.Col
		y := i / m.Col
		text.Draw(m.Canvas, s, f, x*m.W, y*m.H+fHeight, color.White)
		if x == m.X && y == m.Y {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
			m.Canvas.DrawImage(m.Cursor, op)
		}
	}
}
