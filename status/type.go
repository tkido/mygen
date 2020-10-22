package status

//go:generate stringer -type=Type
type Type int

const (
	Human Type = iota
	Vampire
	Woke
	Zombie
	HumanNaked
	VampireNaked
	WokeNaked
	ZombieNaked
)
