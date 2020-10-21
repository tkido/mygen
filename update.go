package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/tkido/mygen/part"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

func updateFace() {
	for i := len(layer.FaceLayers) - 1; 0 <= i; i-- {
		layer := layer.FaceLayers[i]
		fmt.Println(layer)
		files := globParts(sprite.Face, base.Female, layer, "01")
		for i := len(files) - 1; 0 <= i; i-- {
			file := files[i]
			fmt.Println(file)
			imgSrc, _, err := ebitenutil.NewImageFromFile(file, ebiten.FilterDefault)
			if err != nil {
				log.Fatal(err)
			}
			op := &ebiten.DrawImageOptions{}
			imgFace.DrawImage(imgSrc, op)
		}
	}
	f, err := os.Create("dist/outimage.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, imgFace)
	if err != nil {
		log.Fatal(err)
	}
}

func updateMenu() {
	vs := variationMap[base.Female][part.Mouth]
	for i, v := range vs {
		src, _, err := ebitenutil.NewImageFromFile(v.file, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(64*float64(i), 0)
		imgMenu.DrawImage(src, op)
	}
}
