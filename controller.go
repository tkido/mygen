package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
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
		if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
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

func (c *Controller) CursorMove(x, y int) error {
	limitY := 5 + len(part.Types)
	if y != 0 {
		c.cursorY += y
		if c.cursorY == -1 {
			c.cursorY = limitY - 1
		} else if g.cursorY == limitY {
			c.cursorY = 0
		}

	}
	if x != 0 {
		switch g.cursorY {
		case 0: // Id
			g.Character.Id += x
		case 1:
			g.Character.Base += base.Type(x)
			if g.Character.Base >= base.Type(len(base.Types)) {
				g.Character.Base = base.Type(0)
			} else if g.Character.Base < base.Type(0) {
				g.Character.Base = base.Type(len(base.Types) - 1)
			}
		case 2: // Body
		case 3: // Status
		case 4: // Emotion
		default:
			pt := part.Type(g.cursorY - 5)
			if list, ok := g.VariationManager.Map[g.Character.Base][pt]; ok {
				c.cursorX += x
				max := len(list)
				if c.cursorX >= max {
					c.cursorX = 0
				} else if g.cursorX < 0 {
					c.cursorX = max - 1
				}
				g.Character.StatusMap[status.Human].Parts[pt] = part.Index(c.cursorX)
			}
		}
	}
	// g.Logic.UpdateFace()

	return nil
}
