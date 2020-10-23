package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/sprite"
)

type GlobManager struct {
	Cache map[string][]string
}

func NewGlobManager() GlobManager {
	return GlobManager{
		Cache: map[string][]string{},
	}
}

func (gm *GlobManager) Parts(sp sprite.Type, base base.Type, layer layer.Type, label string) []string {
	header := sp.String()
	if sp == sprite.Face {
		header = "FG"
	}
	path := filepath.Join(
		rootPath,
		sp.String(),
		base.String(),
		fmt.Sprintf("%s_%s_p%s_*.png", header, layer, label))
	files, ok := gm.Cache[path]
	if !ok {
		var err error
		files, err = filepath.Glob(path)
		if err != nil {
			log.Fatal(err)
		}
		gm.Cache[path] = files
	}
	return files
}
