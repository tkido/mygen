package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/ui"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	Root ui.Element
	Character
	ImageManager
	GlobManager
	VariationManager
	PartManager

	MainMenu    *MainMenu
	PartMenu    *PartMenu
	PaletteMenu *PaletteMenu
	ColorMenu   *ColorMenu
	Tabs        []ui.Element
	TabIndex    int

	Sprites *Sprites
	Sample  *Sample
}

var (
	g Game
)

func init() {
	g = Game{
		Root:             ui.NewRoot(screenWidth, screenHeight, ui.Color("ff0000")),
		Character:        NewCharacter(0, base.Male),
		ImageManager:     NewImageManager(),
		GlobManager:      NewGlobManager(),
		VariationManager: NewVariationManager(),
		PartManager:      NewPartManager(),
		MainMenu:         NewMainMenu(100, 20, 2, 20),
		PartMenu:         NewPartMenu(64, 64, 12, 7),
		PaletteMenu:      NewPaletteMenu(80, 20, 4, 1),
		ColorMenu:        NewColorMenu(32, 32, 6, 4),
		Tabs:             []ui.Element{},
		TabIndex:         0,
		Sprites:          NewSprites(),
		Sample:           nil,
	}
	g.VariationManager.Init()

	g.Sample = NewSample(g.Sprites)

	g.MainMenu.SetFocus()
	g.MainMenu.Update()
	g.Root.Add(0, 0, g.MainMenu)
	g.Tabs = append(g.Tabs, g.MainMenu)

	g.PartMenu.Update()
	g.Root.Add(200, 0, g.PartMenu)
	g.Tabs = append(g.Tabs, g.PartMenu)

	g.Root.Add(1000, 0, g.PaletteMenu)
	g.Tabs = append(g.Tabs, g.PaletteMenu)

	g.Root.Add(1000, 20, g.ColorMenu)
	g.Tabs = append(g.Tabs, g.ColorMenu)

	changeTab := func(el ui.Element) {
		d := 1
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			d = -1
		}
		g.TabIndex += d
		num := len(g.Tabs)
		if g.TabIndex < 0 {
			g.TabIndex = num - 1
		} else if g.TabIndex >= num {
			g.TabIndex = 0
		}
		g.Tabs[g.TabIndex].SetFocus()
	}
	g.Root.SetKeyCallback(ebiten.KeyTab, changeTab)

	g.Root.Add(0, 64*8, g.Sprites)
	g.Root.Add(800, 64*8, g.Sample)
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.Sample.Update()
	return ui.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
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
