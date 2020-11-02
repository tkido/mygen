package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/ui"
)

type PartMenu struct {
	*MenuBase
	CursorBox *ui.Box
	Part      part.Type
	Data      []*ebiten.Image
}

func NewPartMenu(w, h, col, row int) *PartMenu {
	m := &PartMenu{
		NewMenuBase(w, h, col, row, color.White),
		ui.NewBox(w, h, ui.Color("ff0a")),
		part.Face,
		[]*ebiten.Image{},
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

func (m *PartMenu) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("PartMenu MoveCursor")
	exit = m.MenuBase.MoveCursor(dX, dY)
	m.SetCursor(m.Cursor)
	return exit
}

func (m *PartMenu) SetCursor(index int) {
	log.Printf("PartMenu SetCursor")
	m.MenuBase.SetCursor(index)
	g.Character.StatusMap[g.StatusMenu.Status].Parts[m.Part] = part.Index(index - 1)
	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.CursorBox.Move(x*m.W, y*m.H)
	g.Sprites.Dirty()
}

func (m *PartMenu) Update() {
	ps, ok := g.VariationManager.Map[g.Character.Base][m.Part]
	if !ok {
		log.Fatalf("not found")
	}
	m.Data = []*ebiten.Image{}
	m.Data = append(m.Data, g.ImageManager.Bg)
	for _, p := range ps {
		m.Data = append(m.Data, g.ImageManager.LoadImage(p.file))
	}

	m.Clear()
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

	// restore cursor position from data
	m.Limit = len(m.Data) - 1
	m.SetCursor(int(g.Character.StatusMap[g.StatusMenu.Status].Parts[m.Part] + 1))

	x, y := m.Cursor%m.Col, m.Cursor/m.Col
	m.Add(x*m.W, y*m.H, m.CursorBox)

	m.Dirty()
}

func (m *PartMenu) Reflesh() {
	log.Println("PartMenu.Reflesh")
	m.Image.Fill(ui.Color("0f0"))
	for i, img := range m.Data {
		x, y := i%m.Col, i/m.Col
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
		m.Image.DrawImage(g.ImageManager.Bg, op)
		m.Image.DrawImage(img, op)
	}
}
