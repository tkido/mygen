package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/base"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	Controller
	View
	Character
	ImageManager
	GlobManager
	VariationManager
}

var (
	g Game
)

func init() {
	g = Game{
		Controller:       NewController(),
		View:             NewView(),
		Character:        NewCharacter(0, base.Female),
		ImageManager:     NewImageManager(),
		GlobManager:      NewGlobManager(),
		VariationManager: NewVariationManager(),
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return g.Controller.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.View.Draw(screen)
	g.DebugPrint(screen)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Charactor Generator")
	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowFloating(true)
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
