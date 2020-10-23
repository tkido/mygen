package main

import (
	"github.com/hajimehoshi/ebiten"
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

func (m *MenuBase) MoveCursor(dX, dY int) (exit bool) {
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
		m.Self.SetCursor(newCursor)
		return false
	}
}

func (m *MenuBase) SetCursor(index int) {
	m.Cursor = index
}

type Menu interface {
	Update()
	Reflesh()
	MoveCursor(dX, dY int) (exit bool)
	SetCursor(index int)
}
