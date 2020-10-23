package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
)

type Controller struct {
	cursorX int
	cursorY int
}

func NewController() Controller {
	return Controller{
		cursorY: 0,
		cursorX: 0,
	}
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
	updateFace()

	return nil
}

func (c *Controller) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.CursorMove(0, -1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.CursorMove(0, 1)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.CursorMove(-1, 0)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.CursorMove(1, 0)
	}
	return nil
}
