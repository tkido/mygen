package gradient

type Row int

//go:generate stringer -type=Type
type Type int

const (
	Hair Type = iota
	Eyes
	Skin
	Acc
)

// gradient.png の Start 行目から Number 行がそのタイプが選べる行
type Gradient struct {
	Start  Row
	Number Row
}

var Map = map[Type]Gradient{
	Hair: {0, 23},
	Eyes: {24, 12},
	Skin: {36, 17},
	Acc:  {53, 17},
}
