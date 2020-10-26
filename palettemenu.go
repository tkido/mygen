package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/font"
	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/ui"
)

type PaletteMenu struct {
	*MenuBase
	CursorBox *ui.Box
	Data      []palette.Type
}

func NewPaletteMenu(w, h, col, row int) *PaletteMenu {
	m := &PaletteMenu{
		NewMenuBase(w, h, col, row, color.White),
		ui.NewBox(w, h, ui.Color("ff0a")),
		[]palette.Type{},
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
		log.Printf("PaletteMenu GotFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("f00a"))
	})
	m.SetUiCallback(ui.LostFocus, func(el ui.Element) {
		log.Printf("PaletteMenu LostFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("ff0a"))
	})

	return m
}

func (m *PaletteMenu) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("PaletteMenu MoveCursor")
	exit = m.MenuBase.MoveCursor(dX, dY)
	m.SetCursor(m.Cursor)
	return exit
}

func (m *PaletteMenu) SetCursor(index int) {
	m.MenuBase.SetCursor(index)
	g.ColorMenu.Update()
	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.CursorBox.Move(x*m.W, y*m.H)
}

func (m *PaletteMenu) Update() {
	m.Data = []palette.Type{}
	pt := g.PartMenu.Part
	ps := palette.Map[pt]
	for _, p := range ps {
		m.Data = append(m.Data, p)
	}
	m.Limit = len(m.Data) - 1

	m.Clear()
	var headerLen int
	for i, pt := range m.Data {
		bgColor := ui.Color("fff")
		if i%2 == 1 {
			bgColor = ui.Color("e3ebf1")
		}
		caption := pt.String()
		if i == 0 {
			headerLen = len(caption)
		} else {
			caption = caption[headerLen:]
		}
		label := ui.NewLabel(m.W, m.H, caption, font.Regular, font.Small, ui.Center, color.Black, bgColor)

		index := i
		caption = pt.String()
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

	m.SetCursor(0)
	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.Add(x*m.W, y*m.H, m.CursorBox)

	m.Dirty()
}

func (m *PaletteMenu) Reflesh() {
	log.Println("PaletteMenu.Reflesh")
	m.Image.Fill(ui.Color("0aa"))
}
