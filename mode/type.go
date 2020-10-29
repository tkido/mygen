package mode

//go:generate stringer -type=Type
type Type int

const (
	Normal Type = iota
	Simple
	Animation
)

var Types = []Type{
	Normal,
	Simple,
	Animation,
}
