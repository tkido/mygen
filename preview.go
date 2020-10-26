package main

import (
	"github.com/tkido/mygen/status"
	"github.com/tkido/mygen/ui"
)

type Preview struct {
	*ui.Box
	Status status.Type
}

func NewPreview(s status.Type) *Preview {
	p := &Preview{
		ui.NewBox(144, 144, nil),
		s,
	}
	p.Self = p
	return p
}
