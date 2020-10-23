package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/font"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	Controller
	View
	Character
	Logic
	ImageManager
	GlobManager
	VariationManager
	PartManager
	font.FontManager
	MainMenu *MainMenu
	PartMenu *PartMenu
}

var (
	g Game
)

func init() {
	g = Game{
		Controller:       NewController(),
		View:             NewView(),
		Character:        NewCharacter(0, base.Female),
		Logic:            NewLogic(),
		ImageManager:     NewImageManager(),
		GlobManager:      NewGlobManager(),
		VariationManager: NewVariationManager(),
		PartManager:      NewPartManager(),
		FontManager:      font.NewFontManager(),
		MainMenu:         NewMainMenu(100, 20, 2, 20),
		PartMenu:         NewPartMenu(64, 64, 8, 4),
	}
	g.VariationManager.Init()
	g.FontManager.RegisterFont(font.Regular, "system/mplus-1m-regular.ttf")
	g.Controller.Focused = g.MainMenu
	// g.Controller.Focused = g.PartMenu
	g.PartMenu.Update() // TBD
}

func (g *Game) Update(screen *ebiten.Image) error {
	return g.Controller.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.View.Draw(screen)
	// g.DebugPrint(screen)
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
