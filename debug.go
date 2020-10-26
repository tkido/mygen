package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (g *Game) DebugPrint(screen *ebiten.Image) {
	msg := fmt.Sprintf(`********************************`)
	ebitenutil.DebugPrint(screen, msg)
}
