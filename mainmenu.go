package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/font"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/ui"
)

type MainMenu struct {
	*MenuBase
	CursorBox *ui.Box
	Data      []string
}

func NewMainMenu(w, h, col, row int) *MainMenu {
	m := &MainMenu{
		NewMenuBase(w, h, col, row, color.White),
		nil,
		[]string{},
	}
	m.Self = m

	m.Self.SetKeyCallback(ebiten.KeyLeft, func(el ui.Element) {
		m.MoveCursor(-1, 0)
	})
	m.Self.SetKeyCallback(ebiten.KeyRight, func(el ui.Element) {
		m.MoveCursor(1, 0)
	})
	m.Self.SetKeyCallback(ebiten.KeyUp, func(el ui.Element) {
		m.MoveCursor(0, -1)
	})
	m.Self.SetKeyCallback(ebiten.KeyDown, func(el ui.Element) {
		m.MoveCursor(0, 1)
	})
	return m
}

func (m *MainMenu) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("MainMenu MoveCursor")
	exit = m.MenuBase.MoveCursor(dX, dY)
	m.SetCursor(m.Cursor)
	return exit
}

func (m *MainMenu) SetCursor(index int) {
	log.Printf("MainMenu SetCursor")
	m.MenuBase.SetCursor(index)
	const offset = 2 // num of appended menu item
	if m.Cursor >= offset {
		g.PartMenu.Part = part.Types[m.Cursor-offset]
		g.PartMenu.Update()
		g.PaletteMenu.Update()
	}
	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	log.Printf("SetCursor x: %d, y: %d", x, y)
	m.CursorBox.Move(x*m.W, y*m.H)
}

func (m *MainMenu) Update() {
	m.Data = []string{}
	m.Data = append(m.Data, fmt.Sprintf("%04d", g.Character.Id))
	m.Data = append(m.Data, g.Character.Base.String())
	for _, pt := range part.Types {
		m.Data = append(m.Data, pt.String())
	}
	m.Limit = len(m.Data) - 1

	m.Clear()
	for i, text := range m.Data {
		label := ui.NewLabel(m.W, m.H, text, font.Regular, font.Small, ui.Center, color.Black, color.White)

		caption := text
		index := i
		clicked := func(el ui.Element) {
			log.Printf("%s clicked\n", caption)
			m.SetCursor(index)
		}
		label.SetMouseCallback(
			ui.LeftClick,
			clicked,
		)
		x, y := i%m.Col, i/m.Col
		m.Add(x*m.W, y*m.H, label)
	}
	m.CursorBox = ui.NewBox(m.W, m.H, ui.Color("ff0a"))
	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.Add(x*m.W, y*m.H, m.CursorBox)

	m.Dirty()
}

func (m *MainMenu) Reflesh() {
	log.Println("MainMenu.Reflesh")
	m.Image.Fill(ui.Color("aa0"))
	// f := g.FontManager.Face(font.Regular, font.XSmall)
	// fHeight := f.Metrics().Height.Ceil()

	// for i, _ := range m.Data {
	// 	x := i % m.Col
	// 	y := i / m.Col
	// 	text.Draw(m.Canvas, s, f, x*m.W, y*m.H+fHeight, color.Black)
	// 	if m.Cursor == i {
	// 		m.CursorImg.Fill(g.View.GetFocusColor(m))
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
	// 		m.Canvas.DrawImage(m.CursorImg, op)
	// 	}
	// }
}
