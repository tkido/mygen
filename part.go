package main

import (
	"fmt"
	"path/filepath"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
)

type Part struct {
	id    int
	label string
	file  string
}

type PartManager struct {
	LayerPartMap map[layer.Type]part.Type
}

func NewPartManager() PartManager {
	return PartManager{
		LayerPartMap: map[layer.Type]part.Type{
			layer.AccA:       part.AccA,
			layer.AccB:       part.AccB,
			layer.Beard:      part.Beard,
			layer.Beard1:     part.Beard,
			layer.Beard2:     part.Beard,
			layer.BeastEars:  part.BeastEars,
			layer.Bitten:     part.Bitten,
			layer.Body:       part.Body,
			layer.Cloak:      part.Cloak,
			layer.Cloak1:     part.Cloak,
			layer.Cloak2:     part.Cloak,
			layer.Clothing:   part.Clothing,
			layer.Clothing1:  part.Clothing,
			layer.Clothing2:  part.Clothing,
			layer.Ears:       part.Ears,
			layer.Emotion:    part.Emotion,
			layer.Eyebrows:   part.Eyebrows,
			layer.Eyes:       part.Eyes,
			layer.Face:       part.Face,
			layer.FacialMark: part.FacialMark,
			layer.FrontHair:  part.FrontHair,
			layer.FrontHair1: part.FrontHair,
			layer.FrontHair2: part.FrontHair,
			layer.Glasses:    part.Glasses,
			layer.Head:       part.Head,
			layer.Mouth:      part.Mouth,
			layer.Nose:       part.Nose,
			layer.RearHair:   part.RearHair,
			layer.RearHair1:  part.RearHair,
			layer.RearHair2:  part.RearHair,
			layer.Scar:       part.Scar,
			layer.Soil:       part.Soil,
			layer.Tail:       part.Tail,
			layer.Tail1:      part.Tail,
			layer.Tail2:      part.Tail,
			layer.Wing:       part.Wing,
			layer.Wing1:      part.Wing,
			layer.Wing2:      part.Wing,
		},
	}
}

func (pm *PartManager) Get(sp sprite.Type, base base.Type, layer layer.Type, label string) []string {
	header := sp.String()
	format := "%s_%s_p%s.png"
	if sp == sprite.Face {
		header = "FG"
		format = "%s_%s_p%s_*.png"
	}
	pattern := filepath.Join(
		rootPath,
		sp.String(),
		base.String(),
		fmt.Sprintf(format, header, layer, label))
	return g.GlobManager.Get(pattern)
}
