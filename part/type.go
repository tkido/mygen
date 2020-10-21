package part

import "github.com/tkido/mygen/gradient"

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

type Setting struct {
	Id     int
	Colors []int
}
