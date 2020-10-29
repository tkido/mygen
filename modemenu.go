package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/font"
	"github.com/tkido/mygen/mode"
	"github.com/tkido/mygen/ui"
)

type ModeMenu struct {
	*MenuBase
	CursorBox *ui.Box
	Data      []string
}

func NewModeMenu(w, h, col, row int) *ModeMenu {
	m := &ModeMenu{
		NewMenuBase(w, h, col, row, color.White),
		ui.NewBox(w, h, ui.Color("ff0a")),
		[]string{},
	}
	m.Self = m

	m.SetKeyCallback(ebiten.KeyLeft, func(el ui.Element) {
		m.MoveCursor(-1, 0)
	})
	m.SetKeyCallback(ebiten.KeyRight, func(el ui.Element) {
		m.MoveCursor(1, 0)
	})
	m.SetKeyCallback(ebiten.KeyUp, func(el ui.Element) {
		m.MoveCursor(0, -1)
	})
	m.SetKeyCallback(ebiten.KeyDown, func(el ui.Element) {
		m.MoveCursor(0, 1)
	})

	m.SetUiCallback(ui.GotFocus, func(el ui.Element) {
		log.Printf("ModeMenu GotFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("f00a"))
	})
	m.SetUiCallback(ui.LostFocus, func(el ui.Element) {
		log.Printf("ModeMenu LostFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("ff0a"))
	})

	return m
}

func (m *ModeMenu) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("ModeMenu MoveCursor")
	exit = m.MenuBase.MoveCursor(dX, dY)
	m.SetCursor(m.Cursor)
	return exit
}

func (m *ModeMenu) SetCursor(index int) {
	log.Printf("ModeMenu SetCursor index = %d", index)
	m.MenuBase.SetCursor(index)

	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.CursorBox.Move(x*m.W, y*m.H)

	g.Sprites.Mode = mode.Type(index)
	g.Sprites.Dirty()
}

func (m *ModeMenu) Update() {
	m.Data = []string{}
	for _, mt := range mode.Types {
		m.Data = append(m.Data, mt.String())
	}
	m.Limit = len(m.Data) - 1

	m.Clear()
	for i, text := range m.Data {
		bgColor := ui.Color("fff")
		if i%2 == 1 {
			bgColor = ui.Color("e3ebf1")
		}
		label := ui.NewLabel(m.W, m.H, text, font.Regular, font.Small, ui.Center, color.Black, bgColor)

		caption := text
		index := i
		clicked := func(el ui.Element) {
			log.Printf("%s clicked\n", caption)
			m.SetFocus()
			m.SetCursor(index)
		}
		label.SetMouseCallback(
			ui.LeftClick,
			clicked,
		)
		x, y := i%m.Col, i/m.Col
		m.Add(x*m.W, y*m.H, label)
	}

	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.Add(x*m.W, y*m.H, m.CursorBox)

	m.Dirty()
}

func (m *ModeMenu) Reflesh() {
	log.Println("ModeMenu.Reflesh")
	m.Image.Fill(ui.Color("aa0"))
}
