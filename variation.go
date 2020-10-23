package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
)

type Variation struct {
	id    int
	label string
	file  string
}

type VariationManager struct {
	Map map[base.Type]map[part.Type][]Variation
}

func NewVariationManager() VariationManager {
	re := regexp.MustCompile(`_p(\d+)`)
	vmap := map[base.Type]map[part.Type][]Variation{}
	for _, bt := range base.Types {
		vmap[bt] = map[part.Type][]Variation{}
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
			vmap[bt][pt] = vs
		}
	}
	return VariationManager{
		Map: vmap,
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
