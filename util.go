package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
)

func globVariations(base base.Type, part part.Type) []string {
	path := filepath.Join(
		rootPath,
		"Variation",
		base.String(),
		fmt.Sprintf("icon_%s_*.png", part))
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

var globCache map[string][]string

func init() {
	globCache = map[string][]string{}
}

func globParts(sp sprite.Type, base base.Type, layer layer.Type, label string) []string {
	header := sp.String()
	if sp == sprite.Face {
		header = "FG"
	}
	path := filepath.Join(
		rootPath,
		sp.String(),
		base.String(),
		fmt.Sprintf("%s_%s_p%s*.png", header, layer, label))
	files, ok := globCache[path]
	if !ok {
		var err error
		files, err = filepath.Glob(path)
		if err != nil {
			log.Fatal(err)
		}
		globCache[path] = files
	}
	return files
}
