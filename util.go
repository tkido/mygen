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

func globParts(sp sprite.Type, base base.Type, layer layer.Type) []string {
	header := sp.String()
	if sp == sprite.Face {
		header = "FG"
	}
	path := filepath.Join(
		rootPath,
		sp.String(),
		base.String(),
		fmt.Sprintf("%s_%s_*.png", header, layer))
	fmt.Println(path)
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
