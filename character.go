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
	Base      base.Type
	StatusMap map[status.Type]Status
}

type Status struct {
	Parts  map[part.Type]part.Index
	Colors map[palette.Type]gradient.Row
}

func NewCharacter(id int, bt base.Type) Character {
	c := Character{
		Id:        id,
		Base:      bt,
		StatusMap: map[status.Type]Status{},
	}
	for st := status.Human; st <= status.ZombieNaked; st++ {
		s := Status{
			Parts:  part.NewSetting(bt, st),
			Colors: palette.NewSetting(),
		}
		c.StatusMap[st] = s
	}
	return c
}
