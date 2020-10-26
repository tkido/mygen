package main

import (
	"image/color"
	"log"

	"github.com/tkido/mygen/ui"
)

type MenuBase struct {
	*ui.Box
	W, H, Col, Row int
	Cursor         int
	Limit          int
}

func NewMenuBase(w, h, col, row int, bg color.Color) *MenuBase {
	m := &MenuBase{
		Box:    ui.NewBox(w*col, h*row, bg),
		W:      w,
		H:      h,
		Col:    col,
		Row:    row,
		Cursor: 0,
		Limit:  0,
	}
	m.Self = m

	return m
}

func (m *MenuBase) MoveCursor(dX, dY int) (exit bool) {
	log.Printf("MoveCursor dX:%d dY:%d", dX, dY)
	x := m.Cursor % m.Col
	y := m.Cursor / m.Col
	newCursor := (y+dY)*m.Col + (x + dX)
	switch {
	case dX == -1 && (x == 0 || newCursor < 0):
		return true // 左脱出
	case dX == 1 && (x == m.Col-1 || newCursor > m.Limit):
		return true // 右脱出
	case dY == -1 && (y == 0 || newCursor < 0):
		return true // 上脱出
	case dY == 1 && (y == m.Row-1 || newCursor > m.Limit):
		return true // 下脱出
	default:
		// if mb, ok := m.Self.(*MenuBase); ok {
		// 	mb.SetCursor(newCursor)
		// }
		m.SetCursor(newCursor)
		return false
	}
}

func (m *MenuBase) SetCursor(index int) {
	log.Printf("MenuBase SetCursor")
	m.Cursor = index
}

type Menu interface {
	SetData(data []interface{})
	Reflesh()
	MoveCursor(dX, dY int) (exit bool)
	SetCursor(index int)
}
