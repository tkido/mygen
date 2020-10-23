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
	return VariationManager{
		Map: map[base.Type]map[part.Type][]Variation{},
	}
}

func (vm *VariationManager) Init() {
	re := regexp.MustCompile(`_p(\d+)`)
	for _, bt := range base.Types {
		vm.Map[bt] = map[part.Type][]Variation{}
		for _, pt := range part.Types {
			files := vm.Get(bt, pt)
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
			log.Println(vs)
			vm.Map[bt][pt] = vs
		}
	}
}

func (vm *VariationManager) Get(bt base.Type, pt part.Type) []string {
	pattern := filepath.Join(
		rootPath,
		"Variation",
		bt.String(),
		fmt.Sprintf("icon_%s_*.png", pt))
	return g.GlobManager.Get(pattern)
}
