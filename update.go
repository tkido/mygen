package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

func update(screen *ebiten.Image) {
	for i := len(layer.FaceLayers) - 1; 0 <= i; i-- {
		layer := layer.FaceLayers[i]
		fmt.Println(layer)
		files := globParts(sprite.Face, base.Female, layer)
		fmt.Println(files)
		file := files[0]
		imgSrc, _, err := ebitenutil.NewImageFromFile(file, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(imgSrc, op)
	}
}
