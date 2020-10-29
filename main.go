package main

import (
	"encoding/json"
	"fmt"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
	ModeMenu    *ModeMenu
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

	g.Root.Add(100*2+64*12+100, 0, g.ColorMenu)
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
	file := filepath.Join(".", "_savedata", fmt.Sprintf("%04d.json", g.Character.Id))
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bs, &g.Character); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded!!")

}
func (g *Game) Save(el ui.Element) {
	if !ebiten.IsKeyPressed(ebiten.KeyControl) {
		return
	}
	bs, err := json.Marshal(g.Character)
	if err != nil {
		log.Fatal(err)
	}
	file := filepath.Join(".", "_savedata", fmt.Sprintf("%04d.json", g.Character.Id))
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write(bs); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved!!")
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
