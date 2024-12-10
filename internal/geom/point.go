package geom

// A Point represents a point in a 2D space.
type Point struct {
	X int
	Y int
}

// Add returns a new point with one step in direction d.
func (c Point) Add(d Direction) Point {
	return Point{X: c.X + d.X, Y: c.Y + d.Y}
}

// Neighbors returns all neighboring points.
func (c Point) Neighbors() []Point {
	return []Point{
		c.Add(Left),
		c.Add(Right),
		c.Add(Up),
		c.Add(Down),
	}
}

func (c Point) XY() (x, y int) {
	return c.X, c.Y
}

func (c Point) IsWithinBounds(minX, minY, maxX, maxY int) bool {
	return minX <= c.X && c.X < maxX && minY <= c.Y && c.Y < maxY
}
