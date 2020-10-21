package palette

import "github.com/tkido/mygen/part"

//go:generate stringer -type=Type
type Type int

const (
	Skin Type = iota + 1
	Eyes
	Hair
	HairSub
	FacialMark
	BeastEars
	Clothing
	ClothingSub1
	ClothingSub2
	ClothingSub3
	Cloak
	CloakSub
	AccA
	AccASub1
	AccASub2
	AccB
	AccBSub1
	AccBSub2
	AccBSub3
	Glasses
	GlassesSub1
	GlassesSub2
	Tail
	Wing
)

// gradient.png の Start 行目から Number 行がパレット
type Pallete struct {
	Start  int
	Number int
}

var Map = map[part.Type][]Type{
	part.Face:       {Skin},
	part.FrontHair:  {Hair},
	part.RearHair:   {Hair, HairSub},
	part.Beard:      {Hair},
	part.Ears:       {Skin},
	part.Eyes:       {Eyes},
	part.Eyebrows:   {Hair},
	part.Nose:       {Skin},
	part.Mouth:      {Skin},
	part.FacialMark: {FacialMark},
	part.BeastEars:  {BeastEars},
	part.Tail:       {Tail},
	part.Wing:       {Wing},
	part.Clothing:   {Clothing, ClothingSub1, ClothingSub2, ClothingSub3},
	part.Cloak:      {Cloak, CloakSub},
	part.AccA:       {AccA, AccASub1, AccASub2},
	part.AccB:       {AccB, AccBSub1, AccBSub2, AccBSub3},
	part.Glasses:    {Glasses, GlassesSub1, GlassesSub2},
}
