package ds

// A Grid represents a 2D slice of items with a fixed size.
type Grid[V any] struct {
	Data   [][]V
	Height int
	Width  int
}

// Allocate2DSlice allocates a 2D slice as one large chunk of memory instead of
// multiple smaller ones which is faster to allocate and more cache efficient.
func Allocate2DSlice[T any](width, height int) [][]T {
	rows := make([][]T, height)
	columns := make([]T, height*width)
	for i := range rows {
		rows[i], columns = columns[:width], columns[width:]
	}
	return rows
}

func NewGrid[V any](width, height int) *Grid[V] {
	return &Grid[V]{
		Data:   Allocate2DSlice[V](width, height),
		Height: height,
		Width:  width,
	}
}

// Clone creates a shallow copy of g.
func (g *Grid[V]) Clone() *Grid[V] {
	clone := NewGrid[V](g.Width, g.Height)
	clone.CopyFrom(g)
	return clone
}

// Get retrieves the reference to the data at x,y.
func (g *Grid[V]) Get(x, y int) *V {
	return &g.Data[y][x]
}

func (g *Grid[V]) IsWithinBounds(x, y int) bool {
	return 0 <= x && x < g.Width && 0 <= y && y < g.Height
}

// CopyFrom shallowly copies all data from src into g.
// Panics if g or src have no rows or if g and src differ in size.
func (g *Grid[V]) CopyFrom(src *Grid[V]) {
	n := cap(g.Data[0])
	if n != cap(src.Data[0]) {
		panic("cannot copy src into g as dimensions differ")
	}
	// Since all rows are actually one contiguous memory space, we can copy all
	// data in a single call.
	copy(g.Data[0][:n], src.Data[0][:n])
}
