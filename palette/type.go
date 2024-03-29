package palette

import (
	"image/color"

	"github.com/tkido/mygen/gradient"
	"github.com/tkido/mygen/part"
)

//go:generate stringer -type=Type
type Type int

const (
	Null Type = iota
	Skin
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
	// added
	SkinSub // 25
	Soil    // 26
	Emotion // 27
	Bitten  // 28
	Scar    // 29
)

type Setting map[Type]gradient.Row

func NewSetting() Setting {
	return Setting{
		Null:         -1,
		Skin:         -1,
		SkinSub:      -1,
		Eyes:         -1,
		Hair:         -1,
		HairSub:      -1,
		Soil:         -1,
		FacialMark:   -1,
		Emotion:      -1,
		Bitten:       -1,
		Scar:         -1,
		BeastEars:    -1,
		Clothing:     -1,
		ClothingSub1: -1,
		ClothingSub2: -1,
		ClothingSub3: -1,
		Cloak:        -1,
		CloakSub:     -1,
		AccA:         -1,
		AccASub1:     -1,
		AccASub2:     -1,
		AccB:         -1,
		AccBSub1:     -1,
		AccBSub2:     -1,
		AccBSub3:     -1,
		Glasses:      -1,
		GlassesSub1:  -1,
		GlassesSub2:  -1,
		Tail:         -1,
		Wing:         -1,
	}
}

// gradient.png の Start 行目から Number 行がパレット
type Pallete struct {
	Start  int
	Number int
}

var Map = map[part.Type][]Type{
	part.Body:       {Skin, SkinSub},
	part.Head:       {Skin},
	part.Face:       {Skin},
	part.FrontHair:  {Hair},
	part.RearHair:   {Hair, HairSub},
	part.Beard:      {Hair},
	part.Ears:       {Skin},
	part.Eyes:       {Eyes},
	part.Eyebrows:   {Hair},
	part.Nose:       {Skin},
	part.Mouth:      {Skin},
	part.Soil:       {Soil},
	part.FacialMark: {FacialMark},
	part.Emotion:    {Emotion},
	part.Bitten:     {Bitten},
	part.Scar:       {Scar},
	part.BeastEars:  {BeastEars},
	part.Tail:       {Tail},
	part.Wing:       {Wing},
	part.Clothing:   {Clothing, ClothingSub1, ClothingSub2, ClothingSub3},
	part.Cloak:      {Cloak, CloakSub},
	part.AccA:       {AccA, AccASub1, AccASub2},
	part.AccB:       {AccB, AccBSub1, AccBSub2, AccBSub3},
	part.Glasses:    {Glasses, GlassesSub1, GlassesSub2},
}

var MaskPaletteMap = map[color.Color]Type{
	color.RGBA{249, 193, 157, 255}: Skin,
	color.RGBA{255, 145, 143, 255}: SkinSub,
	color.RGBA{44, 128, 203, 255}:  Eyes,
	color.RGBA{252, 203, 10, 255}:  Hair,
	color.RGBA{184, 146, 197, 255}: HairSub,
	color.RGBA{79, 65, 60, 255}:    Soil,    // TBD
	color.RGBA{87, 87, 85, 255}:    Emotion, // TBD
	color.RGBA{170, 175, 175, 255}: Bitten,  // TBD
	color.RGBA{56, 59, 59, 255}:    Scar,    // TBD
	color.RGBA{0, 146, 150, 255}:   FacialMark,
	color.RGBA{211, 206, 199, 255}: BeastEars,
	color.RGBA{174, 134, 130, 255}: Clothing,
	color.RGBA{254, 157, 30, 255}:  ClothingSub1,
	color.RGBA{28, 118, 208, 255}:  ClothingSub2,
	color.RGBA{217, 164, 4, 255}:   ClothingSub3,
	color.RGBA{216, 172, 0, 255}:   Cloak,
	color.RGBA{163, 7, 8, 255}:     CloakSub,
	color.RGBA{211, 206, 194, 255}: AccA,
	color.RGBA{218, 52, 110, 255}:  AccASub1,
	color.RGBA{164, 201, 17, 255}:  AccASub2,
	color.RGBA{199, 132, 7, 255}:   AccB,
	color.RGBA{192, 211, 210, 255}: AccBSub1,
	color.RGBA{65, 85, 182, 255}:   AccBSub2,
	color.RGBA{186, 59, 69, 255}:   AccBSub3,
	color.RGBA{53, 153, 153, 255}:  Glasses,
	color.RGBA{204, 186, 210, 255}: GlassesSub1,
	color.RGBA{96, 126, 75, 255}:   GlassesSub2,
	color.RGBA{230, 214, 189, 255}: Tail,
	color.RGBA{167, 214, 214, 255}: Wing,
	// nonColors
	color.RGBA{123, 66, 0, 255}:    Null,
	color.RGBA{73, 14, 18, 255}:    Null,
	color.RGBA{218, 121, 18, 255}:  Null,
	color.RGBA{211, 136, 79, 255}:  Null,
	color.RGBA{255, 252, 255, 255}: Null,
	color.RGBA{32, 29, 26, 255}:    Null,
	color.RGBA{183, 118, 38, 255}:  Null,
	color.RGBA{37, 38, 66, 255}:    Null,
}
