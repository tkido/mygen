package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
	View
	ImageManager
	Character
	cursorX int
	cursorY int
}

var (
	game Game
)

func NewCharacter(id int, bt base.Type) Character {
	c := Character{
		Id:        id,
		Base:      bt,
		StatusMap: map[status.Type]Status{},
	}
	for st := status.Human; st <= status.ZombieNaked; st++ {
		s := Status{
			Parts:  part.NewSetting(bt, st),
			Colors: palette.NewSetting(),
		}
		c.StatusMap[st] = s
	}
	return c
}

func init() {

	game = Game{
		View:         NewView(),
		ImageManager: NewImageManager(),
		Character:    NewCharacter(0, base.Female),
		cursorY:      0,
		cursorX:      0,
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
			if list, ok := variationMap[g.Character.Base][pt]; ok {
				g.cursorX += x
				max := len(list)
				if g.cursorX >= max {
					g.cursorX = 0
				} else if g.cursorX < 0 {
					g.cursorX = max - 1
				}
				g.Character.StatusMap[status.Human].Parts[pt] = part.Index(g.cursorX)
			}
		}
	}
	updateFace()

	return nil
}

func (g *Game) Update(screen *ebiten.Image) error {
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

func (g *Game) Draw(screen *ebiten.Image) {
	g.View.Draw(screen)

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
