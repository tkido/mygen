package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var imgCache map[string]*ebiten.Image

func init() {
	imgCache = map[string]*ebiten.Image{}
}

func loadImage(path string) *ebiten.Image {
	img, ok := imgCache[path]
	if !ok {
		var err error
		img, _, err = ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		imgCache[path] = img
	}
	return img
}
