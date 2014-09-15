package robo

type Space int

const (
	Open Space = iota
	Pit
	Wall
	Conveyor
	Pusher
	Gear
	Laser
	Flag
	Repair
)
