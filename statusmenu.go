package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/font"
	"github.com/tkido/mygen/status"
	"github.com/tkido/mygen/ui"
)

type StatusMenu struct {
	*MenuBase
	CursorBox *ui.Box
	Status    status.Type
	Data      []string
}

func NewStatusMenu(w, h, col, row int) *StatusMenu {
	m := &StatusMenu{
		NewMenuBase(w, h, col, row, color.White),
		ui.NewBox(w, h, ui.Color("ff0a")),
		status.Hum,
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
		log.Printf("StatusMenu GotFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("f00a"))
	})
	m.SetUiCallback(ui.LostFocus, func(el ui.Element) {
		log.Printf("StatusMenu LostFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("ff0a"))
	})

	return m
}

func (m *StatusMenu) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("StatusMenu MoveCursor")
	exit = m.MenuBase.MoveCursor(dX, dY)
	m.SetCursor(m.Cursor)
	return exit
}

func (m *StatusMenu) SetCursor(index int) {
	log.Printf("StatusMenu SetCursor index = %d", index)
	m.MenuBase.SetCursor(index)

	m.Status = status.Type(m.Cursor)

	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.CursorBox.Move(x*m.W, y*m.H)

	g.MainMenu.SetCursor(g.MainMenu.Cursor) // reset cursor position
	g.Sprites.Dirty()
}

func (m *StatusMenu) Update() {
	m.Data = []string{}
	for _, st := range status.FaceTypes {
		m.Data = append(m.Data, st.String())
	}
	m.Limit = len(m.Data) - 1

	m.Clear()
	for i, text := range m.Data {
		bgColor := ui.Color("fff")
		switch i % 8 {
		case 1, 3, 4, 6:
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

func (m *StatusMenu) Reflesh() {
	log.Println("StatusMenu.Reflesh")
	m.Image.Fill(ui.Color("aa0"))
}
