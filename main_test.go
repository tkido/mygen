package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"testing"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
)

func TestMakeVariationMap(t *testing.T) {
	re := regexp.MustCompile(`_p(\d+)`)
	for _, bt := range base.Types {
		variationMap[bt] = map[part.Type][]Variation{}
		for _, pt := range part.Types {
			files := globVariations(bt, pt)
			vs := []Variation{}
			for _, file := range files {
				fmt.Println(file)
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
			fmt.Println(vs)
			variationMap[bt][pt] = vs
		}
	}
	// parts := globParts(sprite.Face, base.Male, layer.Face, "01")
	// for _, part := range parts {
	// 	fmt.Println(part)
	// }
	// files = globParts(sprite.SV, base.Male, layer.Wing1)
	// fmt.Println(files)
}