package main

import (
	"fmt"

	"github.com/tkido/mygen/part"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

func updateFace() {
	imgFace.Clear()
	for i := len(layer.FaceLayers) - 1; 0 <= i; i-- {
		lay := layer.FaceLayers[i]
		fmt.Println(lay)

		label := "01"
		if pt, ok := layerPartMap[lay]; ok {
			if list, ok := variationMap[game.setting.Base][pt]; ok {
				if index, ok := game.setting.Parts[pt]; ok {
					if index < len(list) {
						label = list[index].label
					}
				}
			}
		}
		files := globParts(sprite.Face, game.setting.Base, lay, label)
		for i := len(files) - 1; 0 <= i; i-- {
			file := files[i]
			imgSrc := loadImage(file)
			op := &ebiten.DrawImageOptions{}
			imgFace.DrawImage(imgSrc, op)
		}
	}
	// f, err := os.Create("dist/outimage.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// err = png.Encode(f, imgFace)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func updateMenu() {
	vs := variationMap[game.setting.Base][part.Mouth]
	for i, v := range vs {
		src := loadImage(v.file)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(64*float64(i), 0)
		imgMenu.DrawImage(src, op)
	}
}
