package main

import (
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
)

const rootPath = "generator"

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
