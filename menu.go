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
	W, H, Col, Row int
	Cursor         int
	Limit          int
	Canvas         *ebiten.Image
	CursorImg      *ebiten.Image
	Dirty          bool
	Self           Menu
}

func (m *MenuBase) MoveCursor(dX, dY int) {
	x := m.Cursor % m.Col
	y := m.Cursor / m.Col
	newCursor := (y+dY)*m.Col + (x + dX)
	switch {
	case dX == -1 && (x == 0 || newCursor < 0):
		return // 左脱出
	case dX == 1 && (x == m.Col-1 || newCursor > m.Limit):
		return // 右脱出
	case dY == -1 && (y == 0 || newCursor < 0):
		return // 上脱出
	case dY == 1 && (y == m.Row-1 || newCursor > m.Limit):
		return // 下脱出
	default:
		m.Self.SetCursor(newCursor)
	}
}

func (m *MenuBase) SetCursor(index int) {
	m.Cursor = index
}

type Menu interface {
	Update()
	Reflesh()
	MoveCursor(dX, dY int)
	SetCursor(index int)
	// SetFocus(index int)
	// IsFocused() bool
	// GetFocus()
	// LoseFocus()
}

type MainMenu struct {
	MenuBase
	Data []string
}

func NewMainMenu(w, h, col, row int) *MainMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursorImg, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	cursorImg.Fill(color.RGBA{255, 255, 0, 64})

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
	}
	g.Logic.UpdateFace()
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
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
			m.Canvas.DrawImage(m.CursorImg, op)
		}
	}
}
