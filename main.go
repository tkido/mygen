package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type Game struct {
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
	game = Game{}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		updateFace()
		updateMenu()
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
