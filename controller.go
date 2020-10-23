package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Controller struct {
	cursorX  int
	cursorY  int
	TabIndex int
	Menus    []Menu
}

func NewController() Controller {
	return Controller{
		cursorY: 0,
		cursorX: 0,
	}
}

func (c *Controller) Update() error {
	focused := c.Menus[c.TabIndex]
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if focused.MoveCursor(0, -1) {
			// pageUp
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if focused.MoveCursor(0, 1) {
			// pageDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if focused.MoveCursor(-1, 0) {
			// c.MoveMenu(-1)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if focused.MoveCursor(1, 0) {
			// c.MoveMenu(1)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		d := 1
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			d = -1
		}
		c.MoveMenu(d)
	}
	return nil
}
func (c *Controller) MoveMenu(d int) {
	from := c.TabIndex
	c.TabIndex += d
	if c.TabIndex < 0 {
		c.TabIndex = len(c.Menus) - 1
	} else if c.TabIndex >= len(c.Menus) {
		c.TabIndex = 0
	}
	c.Menus[from].Reflesh()
	c.Menus[c.TabIndex].Reflesh()
}
