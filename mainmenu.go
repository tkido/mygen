package main

import (
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
		log.Printf("MainMenu GotFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("f00a"))
	})
	m.SetUiCallback(ui.LostFocus, func(el ui.Element) {
		log.Printf("MainMenu LostFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("ff0a"))
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
	log.Printf("MainMenu SetCursor index = %d", index)
	m.MenuBase.SetCursor(index)
	g.PartMenu.Part = part.Types[m.Cursor]
	g.PartMenu.Update()
	g.PaletteMenu.Update()
	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.CursorBox.Move(x*m.W, y*m.H)

	g.Sprites.Dirty()
}

func (m *MainMenu) Update() {
	m.Data = []string{}
	for _, pt := range part.Types {
		m.Data = append(m.Data, pt.String())
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

func (m *MainMenu) Reflesh() {
	log.Println("MainMenu.Reflesh")
	m.Image.Fill(ui.Color("aa0"))
}
