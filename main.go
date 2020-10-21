package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	cursorX int
	cursorY int
	setting Setting
}

type Setting struct {
	Id    int
	Base  base.Type
	Parts map[part.Type]int
}

var (
	game    Game
	imgFace *ebiten.Image
	imgMenu *ebiten.Image
	imgBg   *ebiten.Image
)

func init() {
	imgFace, _ = ebiten.NewImage(144, 144, ebiten.FilterDefault)
	imgMenu, _ = ebiten.NewImage(64*64, 64, ebiten.FilterDefault)
	imgBg, _, _ = ebitenutil.NewImageFromFile("system/background.png", ebiten.FilterDefault)
	game = Game{
		cursorY: 0,
		cursorX: 0,
		setting: Setting{
			Id:    3,
			Base:  base.Female,
			Parts: map[part.Type]int{},
		},
	}
}

func (g *Game) CursorMove(x, y int) error {
	limitY := 5 + len(part.Types)
	if y != 0 {
		g.cursorY += y
		if g.cursorY == -1 {
			g.cursorY = limitY - 1
		} else if g.cursorY == limitY {
			g.cursorY = 0
		}

	}
	if x != 0 {
		switch g.cursorY {
		case 0: // Id
			g.setting.Id += x
		case 1:
			g.setting.Base += base.Type(x)
			if g.setting.Base >= base.Type(len(base.Types)) {
				g.setting.Base = base.Type(0)
			} else if g.setting.Base < base.Type(0) {
				g.setting.Base = base.Type(len(base.Types) - 1)
			}
		case 2: // Body
		case 3: // Status
		case 4: // Emotion
		default:
			pt := part.Type(g.cursorY - 5)
			if list, ok := variationMap[g.setting.Base][pt]; ok {
				g.cursorX += x
				max := len(list)
				if g.cursorX >= max {
					g.cursorX = 0
				} else if g.cursorX < 0 {
					g.cursorX = max - 1
				}
				g.setting.Parts[pt] = g.cursorX
			}
		}
	}

	return nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		updateFace()
		updateMenu()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
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

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < 30; i++ {
		for j := 0; j < 17; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i)*64, float64(j)*64)
			screen.DrawImage(imgBg, op)
		}
	}
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(imgMenu, op)
	op.GeoM.Translate(0, 64)
	screen.DrawImage(imgFace, op)

	msg := fmt.Sprintf(`********************************
cursorX: %d
cursorY: %d
genre: %s
ID: %d
Base: %s
Parts: %v
********************************`,
		g.cursorX,
		g.cursorY,
		part.Type(g.cursorY-5),
		g.setting.Id,
		g.setting.Base,
		g.setting.Parts,
	)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Charactor Generator")
	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowFloating(true)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
