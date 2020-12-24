package layer

import (
	"github.com/tkido/mygen/sprite"
)

//go:generate stringer -type=Type
type Type int

const (
	AccA Type = iota
	AccB
	Beard
	Beard1
	Beard2
	BeastEars
	Bitten
	Body
	Cloak
	Cloak1
	Cloak2
	Clothing
	Clothing1
	Clothing2
	Ears
	Emotion
	Eyebrows
	Eyes
	Face
	FacialMark
	FrontHair
	FrontHair1
	FrontHair2
	Glasses
	Head
	Mouth
	Nose
	RearHair
	RearHair1
	RearHair2
	Scar
	Soil
	Tail
	Tail1
	Tail2
	Wing
	Wing1
	Wing2
)

var Map = map[sprite.Type][]Type{
	sprite.Face: []Type{
		AccB, Glasses, FrontHair, Cloak1, AccA, BeastEars,
		Beard, Clothing1, Ears, RearHair1, Eyebrows, Eyes,
		Soil, FacialMark, Emotion, Bitten, Scar, Nose, Mouth, Face, Clothing2, Body, Cloak2,
		RearHair2,
	},
	sprite.Tv: []Type{
		Wing1, AccB, FrontHair1, AccA, RearHair1, Glasses,
		Ears, BeastEars, Cloak1, Tail1, Clothing1, Beard1, Clothing2,
		RearHair2, Soil, FacialMark, Emotion, Bitten, Scar,
		Head, Body, Beard2, FrontHair2, Tail2,
		Cloak2, Wing2,
	},
	sprite.Tvd: []Type{
		AccB, FrontHair, AccA, Eyes, Ears, BeastEars, RearHair, Wing,
		Glasses, Cloak, Tail, Beard, Clothing, Soil, FacialMark, Emotion, Bitten, Scar, Head,
	},
	sprite.Sv: []Type{
		AccB, FrontHair, AccA, Glasses, Ears, BeastEars, Cloak1,
		Clothing1, Beard, Clothing2, RearHair1, Soil, FacialMark, Emotion, Bitten, Scar, Head, Body,
		Tail, Cloak2, Wing,
	},
}
