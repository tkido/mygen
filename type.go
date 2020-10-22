package main

import (
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
)

const (
	rootPath = "generator"
	distPath = "dist"
)

var variationMap = map[base.Type]map[part.Type][]Variation{}
var partsMap = map[sprite.Type]map[base.Type]map[part.Type][]Part{}

type Variation struct {
	id    int
	label string
	file  string
}

type Part struct {
	id      int
	label   string
	file    string
	colorId int
}

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
