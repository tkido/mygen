package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/part"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

var reDefaultColor = regexp.MustCompile(`_m(\d{3})`)

func updateFace() {
	imgFace.Clear()
	for i := len(layer.FaceLayers) - 1; 0 <= i; i-- {
		lay := layer.FaceLayers[i]
		// log.Printf("Set %s...\n", lay)

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
			// fmt.Println(file)
			imgSrc := loadImage(file)
			// default color found
			if ms := reDefaultColor.FindStringSubmatch(file); len(ms) >= 2 {
				label := ms[1]
				index, err := strconv.Atoi(label)
				if err != nil {
					log.Fatalf("invalid default color label")
				}
				p := palette.Type(index)
				if p == palette.Skin {
					fmt.Println(p)
					imgSrc = filterImage(imgSrc)
				}
			}
			op := &ebiten.DrawImageOptions{}
			imgFace.DrawImage(imgSrc, op)
		}
	}

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
