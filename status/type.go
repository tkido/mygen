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
	Emotion00
	Emotion01
	Emotion02
	Emotion03
	Emotion04
	Emotion05
	Emotion06
	Emotion07
	Emotion08
	Emotion09
	Emotion10
	Emotion11
	Emotion12
	Emotion13
	Emotion14
	Emotion15
	Special00
	Special01
	Special02
	Special03
	Special04
	Special05
	Special06
	Special07
	Special08
	Special09
	Special10
	Special11
	Special12
	Special13
	Special14
	Special15
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

var FaceTypes = []Type{
	Human,
	HumanNaked,
	Vampire,
	VampireNaked,
	Woke,
	WokeNaked,
	Zombie,
	ZombieNaked,
	Emotion00,
	Emotion01,
	Emotion02,
	Emotion03,
	Emotion04,
	Emotion05,
	Emotion06,
	Emotion07,
	Emotion08,
	Emotion09,
	Emotion10,
	Emotion11,
	Emotion12,
	Emotion13,
	Emotion14,
	Emotion15,
	Special00,
	Special01,
	Special02,
	Special03,
	Special04,
	Special05,
	Special06,
	Special07,
	Special08,
	Special09,
	Special10,
	Special11,
	Special12,
	Special13,
	Special14,
	Special15,
}
