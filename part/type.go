package part

import (
	"github.com/tkido/mygen/gradient"
)

type Index int

const Null Index = -1

func NewSetting() Setting {
	return Setting{
		AccA:       Null,
		AccB:       Null,
		BeastEars:  Null,
		Beard:      Null,
		Body:       0,
		Clothing:   Null,
		Cloak:      Null,
		Ears:       0,
		Eyebrows:   Null,
		Eyes:       0,
		Face:       0,
		Soil:       Null,
		FacialMark: Null,
		Emotion:    Null,
		Bitten:     Null,
		Scar:       Null,
		FrontHair:  Null,
		Glasses:    Null,
		Head:       0,
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
	AccA Type = iota
	AccB
	BeastEars
	Beard
	Bitten
	Body
	Clothing
	Cloak
	Ears
	Emotion
	Eyebrows
	Eyes
	Face
	FacialMark
	FrontHair
	Glasses
	Head
	Mouth
	Nose
	RearHair
	Scar
	Soil
	Tail
	Wing
)

var Types = []Type{
	Body,
	Head,
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
	Soil,
	FacialMark,
	Emotion,
	Bitten,
	Scar,
	BeastEars,
	AccA,
	AccB,
	Tail,
	Wing,
	Clothing,
	Cloak,
}

var GradientMap = map[Type]gradient.Type{
	Body:       gradient.Skin,
	Head:       gradient.Skin,
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
	Soil:       gradient.Acc,
	FacialMark: gradient.Acc,
	Emotion:    gradient.Acc,
	Bitten:     gradient.Acc,
	Scar:       gradient.Acc,
	BeastEars:  gradient.Hair,
	AccA:       gradient.Acc,
	AccB:       gradient.Acc,
	Tail:       gradient.Hair,
	Wing:       gradient.Acc,
	Clothing:   gradient.Acc,
	Cloak:      gradient.Acc,
}
