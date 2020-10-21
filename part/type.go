package part

//go:generate stringer -type=Type
type Type int

const (
	// Body Penis
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
