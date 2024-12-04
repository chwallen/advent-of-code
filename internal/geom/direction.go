package geom

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
