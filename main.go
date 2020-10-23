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
	MainMenu    *MainMenu
	PartMenu    *PartMenu
	PaletteMenu *PaletteMenu
	ColorMenu   *ColorMenu
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
		PartMenu:         NewPartMenu(64, 64, 12, 7),
		PaletteMenu:      NewPaletteMenu(80, 20, 4, 1),
		ColorMenu:        NewColorMenu(32, 32, 6, 4),
	}
	g.VariationManager.Init()
	g.FontManager.RegisterFont(font.Regular, "system/mplus-1m-regular.ttf")

	g.Controller.Menus = append(g.Controller.Menus, g.MainMenu)
	g.Controller.Menus = append(g.Controller.Menus, g.PartMenu)
	g.Controller.Menus = append(g.Controller.Menus, g.PaletteMenu)
	g.Controller.Menus = append(g.Controller.Menus, g.ColorMenu)
	g.MainMenu.Update()    // TBD
	g.PartMenu.Update()    // TBD
	g.PaletteMenu.Update() // TBD
	g.ColorMenu.Update()   // TBD

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
