package main

import (
	"fmt"
	"testing"

	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/sprite"
)

func TestGlob(t *testing.T) {
	var files []string
	files = globVariations(base.Male, part.Wing)
	fmt.Println(files)
	files = globParts(sprite.SV, base.Male, part.Wing)
	fmt.Println(files)
}
