package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

var reDefaultColor = regexp.MustCompile(`_m(\d{3})`)

func updateFace() {
	g.View.Face.Clear()
	for i := len(layer.FaceLayers) - 1; 0 <= i; i-- {
		lay := layer.FaceLayers[i]
		// log.Printf("Set %s...\n", lay)

		label := "01"
		if pt, ok := layerPartMap[lay]; ok {
			if list, ok := variationMap[g.Character.Base][pt]; ok {
				if index, ok := g.Character.StatusMap[status.Human].Parts[pt]; ok {
					if index == part.Null {
						continue
					}
					if index < part.Index(len(list)) {
						label = list[index].label
					}
				}
			}
		}

		files := globParts(sprite.Face, g.Character.Base, lay, label)
		for i := len(files) - 1; 0 <= i; i-- {
			file := files[i]
			fmt.Println(file)
			imgSrc := g.ImageManager.LoadImage(file)
			// color info found
			if ms := reDefaultColor.FindStringSubmatch(file); len(ms) >= 2 {
				label := ms[1]
				index, err := strconv.Atoi(label)
				if err != nil {
					log.Fatalf("invalid color info")
				}
				p := palette.Type(index)
				if p == palette.Skin {
					fmt.Println(p)
					// imgSrc = filterImage(imgSrc)
				}
			}
			op := &ebiten.DrawImageOptions{}
			g.View.Face.DrawImage(imgSrc, op)
		}
	}

}

func updateMenu() {
	vs := variationMap[g.Character.Base][part.Mouth]
	for i, v := range vs {
		src := g.ImageManager.LoadImage(v.file)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(64*float64(i), 0)
		g.View.Menu.DrawImage(src, op)
	}
}
