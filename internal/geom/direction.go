package geom

import "fmt"

type Direction Point

var (
	Up        = Direction{X: 0, Y: -1}
	UpRight   = Direction{X: 1, Y: -1}
	Right     = Direction{X: 1, Y: 0}
	DownRight = Direction{X: 1, Y: 1}
	Down      = Direction{X: 0, Y: 1}
	DownLeft  = Direction{X: -1, Y: 1}
	Left      = Direction{X: -1, Y: 0}
	UpLeft    = Direction{X: -1, Y: -1}
)

// TurnRight returns a new point which is turned 90 degrees right.
func (d Direction) TurnRight() Direction {
	return Direction{-d.Y, d.X}
}

// TurnLeft returns the point which is turned 90 degrees left.
func (d Direction) TurnLeft() Direction {
	return Direction{d.Y, -d.X}
}

// GetCardinalIndex gets the int which represents the direction.
// Only works for the directions Up, Right, Down, and Left.
func (d Direction) GetCardinalIndex() int {
	switch d {
	case Up:
		return 0
	case Right:
		return 1
	case Down:
		return 2
	case Left:
		return 3
	default:
		panic(fmt.Errorf("invalid direction %d", d))
	}
}
