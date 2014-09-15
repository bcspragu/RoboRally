package robo

type Bot int

// The eight RoboRally characters
const (
	ZoomBot Bot = iota
	Twitch
	HammerBot
	SquashBot
	SpinBot
	HulkX90
	Twonky
	TrundleBot
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Location struct {
	Orientation Direction
	X           int
	Y           int
}

type Player struct {
	Character   Bot
	Place       Location
	LifeTokens  int
	PoweredDown bool
	Damage      int
	Hand        []ProgramCard
}
