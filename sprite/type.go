package sprite

//go:generate stringer -type=Type
type Type int

const (
	Face Type = iota
	Sv
	Tv
	Tvd
)

var Types = []Type{
	Face,
	Sv,
	Tv,
	Tvd,
}
