package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/tkido/mygen/part"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

var layerPartMap = map[layer.Type]part.Type{
	layer.AccA:      part.AccA,
	layer.AccB:      part.AccB,
	layer.Beard:     part.Beard,
	layer.Beard1:    part.Beard,
	layer.Beard2:    part.Beard,
	layer.BeastEars: part.BeastEars,
	// layer.Body:       part.Body,
	layer.Cloak:      part.Cloak,
	layer.Cloak1:     part.Cloak,
	layer.Cloak2:     part.Cloak,
	layer.Clothing:   part.Clothing,
	layer.Clothing1:  part.Clothing,
	layer.Clothing2:  part.Clothing,
	layer.Ears:       part.Ears,
	layer.Eyebrows:   part.Eyebrows,
	layer.Eyes:       part.Eyes,
	layer.Face:       part.Face,
	layer.FacialMark: part.FacialMark,
	layer.FrontHair:  part.FrontHair,
	layer.FrontHair1: part.FrontHair,
	layer.FrontHair2: part.FrontHair,
	layer.Glasses:    part.Glasses,
	layer.Mouth:      part.Mouth,
	layer.Nose:       part.Nose,
	layer.RearHair:   part.RearHair,
	layer.RearHair1:  part.RearHair,
	layer.RearHair2:  part.RearHair,
	layer.Tail:       part.Tail,
	layer.Tail1:      part.Tail,
	layer.Tail2:      part.Tail,
	layer.Wing:       part.Wing,
	layer.Wing1:      part.Wing,
	layer.Wing2:      part.Wing,
}

func updateFace() {
	imgFace, _ = ebiten.NewImage(144, 144, ebiten.FilterDefault)
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
	vs := variationMap[game.setting.Base][part.Mouth]
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
