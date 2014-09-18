package robo

type GridType int

const (
	Open GridType = iota
	Pit
	Wall
	Conveyor
	Pusher
	Gear
	Laser
	Flag
	Repair
)

type Space interface {
	Deadly() bool
	Moves(p Player, g Game) Location
}

type Board struct {
	Spaces [][]Space
}
