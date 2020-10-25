package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tkido/mygen/part"
)

func (g *Game) DebugPrint(screen *ebiten.Image) {
	msg := fmt.Sprintf(`********************************
cursorX: %d
cursorY: %d
genre: %s
ID: %d
Base: %s
********************************`,
		g.cursorX,
		g.cursorY,
		part.Type(g.cursorY-5),
		g.Character.Id,
		g.Character.Base,
	)
	ebitenutil.DebugPrint(screen, msg)
}
