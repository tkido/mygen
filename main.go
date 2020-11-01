package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/flag"
	"github.com/tkido/mygen/ui"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

var (
	g Game
)

type Game struct {
	Root      ui.Element
	Character *Character

	ImageManager
	GlobManager
	VariationManager
	PartManager
	SaveManager

	MainMenu    *MainMenu
	PartMenu    *PartMenu
	PaletteMenu *PaletteMenu
	ColorMenu   *ColorMenu
	Tabs        []ui.Element
	TabIndex    int
	ModeMenu    *ModeMenu

	Sprites *Sprites
	Sample  *Sample
}

func init() {
	sm := NewSaveManager("_savedata")
	char := NewCharacter(flag.Id, flag.Base)
	if sm.Exists(sm.FileName(char.Id)) {
		char = sm.Load(char.Id)
	}
	g = Game{
		Root:             ui.NewRoot(screenWidth, screenHeight, ui.Color("ff0000")),
		Character:        char,
		ImageManager:     NewImageManager(),
		GlobManager:      NewGlobManager(),
		VariationManager: NewVariationManager(),
		PartManager:      NewPartManager(),
		SaveManager:      sm,
		MainMenu:         NewMainMenu(100, 20, 2, 20),
		PartMenu:         NewPartMenu(64, 64, 12, 7),
		PaletteMenu:      NewPaletteMenu(100, 20, 1, 4),
		ColorMenu:        NewColorMenu(32, 32, 6, 4),
		ModeMenu:         NewModeMenu(100, 20, 1, 3),
		Tabs:             []ui.Element{},
		TabIndex:         0,
		Sprites:          NewSprites(),
		Sample:           nil,
	}
	g.VariationManager.Init()

	g.Sample = NewSample(g.Sprites)

	g.MainMenu.SetFocus()
	g.MainMenu.Update()
	g.MainMenu.SetCursor(0)
	g.Root.Add(0, 0, g.MainMenu)
	g.Tabs = append(g.Tabs, g.MainMenu)

	g.Root.Add(100*2, 0, g.PartMenu)
	g.Tabs = append(g.Tabs, g.PartMenu)

	g.Root.Add(100*2+64*12, 0, g.PaletteMenu)
	g.Tabs = append(g.Tabs, g.PaletteMenu)

	g.Root.Add(100*2+64*12+100, 4, g.ColorMenu)
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

	g.ModeMenu.Update()
	g.Root.Add(0, 64*8, g.ModeMenu)

	g.Root.Add(100, 64*8, g.Sprites)
	g.Root.Add(900, 64*8, g.Sample)

	g.Root.SetKeyCallback(ebiten.KeyS, g.Save)
	g.Root.SetKeyCallback(ebiten.KeyL, g.Load)
}

func (g *Game) Load(el ui.Element) {
	if !ebiten.IsKeyPressed(ebiten.KeyControl) {
		return
	}
	g.Character = g.SaveManager.Load(g.Character.Id)
}

func (g *Game) Save(el ui.Element) {
	if !ebiten.IsKeyPressed(ebiten.KeyControl) {
		return
	}
	g.SaveManager.Save()
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
