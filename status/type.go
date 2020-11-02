package status

//go:generate stringer -type=Type
type Type int

const (
	Human Type = iota
	HumanNaked
	Vampire
	VampireNaked
	Woke
	WokeNaked
	Zombie
	ZombieNaked
)

var Types = []Type{
	Human,
	HumanNaked,
	Vampire,
	VampireNaked,
	Woke,
	WokeNaked,
	Zombie,
	ZombieNaked,
}
