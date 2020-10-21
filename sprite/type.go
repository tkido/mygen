package sprite

//go:generate stringer -type=Type
type Type int

const (
	Face Type = iota
	SV
	TV
	TVD
)
