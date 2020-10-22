package part

import (
	"github.com/tkido/mygen/base"
	"github.com/tkido/mygen/gradient"
	"github.com/tkido/mygen/status"
)

type Index int

const Null Index = -1

func NewSetting(bt base.Type, st status.Type) Setting {
	return Setting{
		AccA:       Null,
		AccB:       Null,
		BeastEars:  Null,
		Beard:      Null,
		Clothing:   Null,
		Cloak:      Null,
		Ears:       0,
		Eyebrows:   Null,
		Eyes:       0,
		Face:       0,
		FacialMark: Null,
		FrontHair:  Null,
		Glasses:    Null,
		Mouth:      0,
		Nose:       0,
		RearHair:   Null,
		Tail:       Null,
		Wing:       Null,
	}
}

type Setting map[Type]Index

//go:generate stringer -type=Type
type Type int

const (
	// Body
	AccA Type = iota
	AccB
	BeastEars
	Beard
	Clothing
	Cloak
	Ears
	Eyebrows
	Eyes
	Face
	FacialMark
	FrontHair
	Glasses
	Mouth
	Nose
	RearHair
	Tail
	Wing
	// Penis
)

var Types = []Type{
	Face,
	RearHair,
	FrontHair,
	Eyebrows,
	Eyes,
	Glasses,
	Ears,
	Nose,
	Beard,
	Mouth,
	FacialMark,
	BeastEars,
	AccA,
	AccB,
	Tail,
	Wing,
	Clothing,
	Cloak,
}

var GradientMap = map[Type]gradient.Type{
	Face:       gradient.Skin,
	RearHair:   gradient.Hair,
	FrontHair:  gradient.Hair,
	Eyebrows:   gradient.Hair,
	Eyes:       gradient.Eyes,
	Glasses:    gradient.Acc,
	Ears:       gradient.Skin,
	Nose:       gradient.Skin,
	Beard:      gradient.Hair,
	Mouth:      gradient.Skin,
	FacialMark: gradient.Acc,
	BeastEars:  gradient.Hair,
	AccA:       gradient.Acc,
	AccB:       gradient.Acc,
	Tail:       gradient.Hair,
	Wing:       gradient.Acc,
	Clothing:   gradient.Acc,
	Cloak:      gradient.Acc,
}
