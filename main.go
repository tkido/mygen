package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
}

var (
	game     Game
	gradient *ebiten.Image
	img1     *ebiten.Image
	img2     *ebiten.Image
)

func init() {
	var err error
	gradient, _, err = ebitenutil.NewImageFromFile("generator/gradients.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	game = Game{}

}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		screen.Fill(color.RGBA{0, 0, 0xff, 0xff})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(200, 200)
		screen.DrawImage(gradient, op)
		update(screen)
	}
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
