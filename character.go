package main

import (
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/gradient"
	"github.com/tkido/mygen/palette"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
)

type Character struct {
	Id        int
	Name      string
	Base      base.Type
	StatusMap map[status.Type]Status
}

type Status struct {
	Parts  map[part.Type]part.Index
	Colors map[palette.Type]gradient.Row
}

func NewCharacter(id int, bt base.Type) *Character {
	c := &Character{
		Id:        id,
		Name:      "",
		Base:      bt,
		StatusMap: map[status.Type]Status{},
	}
	for st := status.Hum; st <= status.S63; st++ {
		s := Status{
			Parts:  part.NewSetting(),
			Colors: palette.NewSetting(),
		}
		c.StatusMap[st] = s
	}
	return c
}
