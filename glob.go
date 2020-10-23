package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/layer"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
)

var variationMap = map[base.Type]map[part.Type][]Variation{}

type Variation struct {
	id    int
	label string
	file  string
}

func init() {
	// make variationMap
	re := regexp.MustCompile(`_p(\d+)`)
	for _, bt := range base.Types {
		variationMap[bt] = map[part.Type][]Variation{}
		for _, pt := range part.Types {
			files := g.GlobManager.Variations(bt, pt)
			vs := []Variation{}
			for _, file := range files {
				fileName := filepath.Base(file)
				ms := re.FindStringSubmatch(fileName)
				if len(ms) < 2 {
					log.Fatalf("part number not found")
				}
				label := ms[1]
				id, err := strconv.Atoi(label)
				if err != nil {
					log.Fatal(err)
				}
				v := Variation{id, label, file}
				vs = append(vs, v)
			}
			sort.Slice(vs, func(i, j int) bool {
				return vs[i].id < vs[j].id
			})
			variationMap[bt][pt] = vs
		}
	}
}

type GlobManager struct {
	Cache map[string][]string
}

func NewGlobManager() GlobManager {
	return GlobManager{
		Cache: map[string][]string{},
	}
}

func (gm *GlobManager) Variations(base base.Type, part part.Type) []string {
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
