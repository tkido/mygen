package base

//go:generate stringer -type=Type
type Type int

const (
	Male Type = iota
	Female
	Kid
)
