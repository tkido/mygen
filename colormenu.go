package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/gradient"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
	"github.com/tkido/mygen/ui"
)

type ColorMenu struct {
	*MenuBase
	CursorBox *ui.Box
	Part      part.Type
	Data      []gradient.Row
}

func NewColorMenu(w, h, col, row int) *ColorMenu {
	m := &ColorMenu{
		NewMenuBase(w, h, col, row, color.White),
		ui.NewBox(w, h, ui.Color("ff0a")),
		part.Face,
		[]gradient.Row{},
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
		log.Printf("PartMenu GotFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("f00a"))
	})
	m.SetUiCallback(ui.LostFocus, func(el ui.Element) {
		log.Printf("PartMenu LostFocus")
		m.CursorBox.SetBackgroundColor(ui.Color("ff0a"))
	})

	return m
}

func (m *ColorMenu) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("ColorMenu MoveCursor")
	exit = m.MenuBase.MoveCursor(dX, dY)
	m.SetCursor(m.Cursor)
	return exit
}

func (m *ColorMenu) SetCursor(index int) {
	log.Printf("ColorMenu SetCursor")
	m.MenuBase.SetCursor(index)

	pt := g.PaletteMenu.Data[g.PaletteMenu.Cursor]
	g.Character.StatusMap[status.Human].Colors[pt] = m.Data[m.Cursor]

	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.CursorBox.Move(x*m.W, y*m.H)
}

func (m *ColorMenu) Update() {
	m.Clear()
	m.Data = []gradient.Row{-1}
	newCursor := 0

	if gt, ok := part.GradientMap[g.PartMenu.Part]; ok {
		pt := g.PaletteMenu.Data[g.PaletteMenu.Cursor]
		currentRow := g.Character.StatusMap[status.Human].Colors[pt]

		g := gradient.Map[gt]
		for row := g.Start; row <= g.Start+g.Number; row++ {
			if currentRow == row {
				newCursor = len(m.Data)
			}
			m.Data = append(m.Data, row)
		}
	}
	m.Limit = len(m.Data) - 1
	m.SetCursor(newCursor)

	for i, _ := range m.Data {
		box := ui.NewBox(m.W, m.H, nil)
		index := i
		clicked := func(el ui.Element) {
			m.SetFocus()
			m.SetCursor(index)
		}
		box.SetMouseCallback(ui.LeftClick, clicked)
		x, y := i%m.Col, i/m.Col
		m.Add(x*m.W, y*m.H, box)
	}

	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.Add(x*m.W, y*m.H, m.CursorBox)

	m.Dirty()
}

func (m *ColorMenu) Reflesh() {
	log.Println("ColorMenu.Reflesh")
	m.Image.Fill(ui.Color("00f"))

	for i, row := range m.Data {
		x, y := i%m.Col, i/m.Col
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.25, 8)
		op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
		r := int(row)
		rect := image.Rect(64, r*4, 256-64, (r+1)*4)
		m.Image.DrawImage(g.ImageManager.Gradient.SubImage(rect).(*ebiten.Image), op)
	}
}
