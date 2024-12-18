package day18

import (
	"math"
	"slices"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
	"github.com/chwallen/advent-of-code/internal/util"
)

type tile struct {
	corrupted bool
	score     int
}

type mazePath struct {
	pos   geom.Point
	tile  *tile
	score int
}

const (
	startingCorruption = 1024

	gridHeight = 71
	gridWidth  = 71
)

var start = geom.Point{X: 0, Y: 0}

func PartOne(lines []string, extras ...any) any {
	height, width, corruption := getParameters(extras)
	grid := createTileGrid(
		lines,
		height,
		width,
		corruption,
	)
	return getFewestStepsToTraverse(grid)
}

func PartTwo(lines []string, extras ...any) any {
	height, width, _ := getParameters(extras)
	grid := createTileGrid(
		lines,
		height,
		width,
		// Start backwards by maximizing corruption
		len(lines),
	)

	var cutoffTile string
	for i, line := range slices.Backward(lines) {
		if getFewestStepsToTraverse(grid) < math.MaxInt {
			cutoffIndex := min(i+1, len(lines))
			cutoffTile = lines[cutoffIndex]
			break
		}

		for y := range height {
			for x := range width {
				grid.Get(x, y).score = math.MaxInt
			}
		}
		setTileCorruption(grid, line, false)
	}

	return cutoffTile
}

func getParameters(extras []any) (height, width, corruption int) {
	height = gridHeight
	width = gridWidth
	corruption = startingCorruption
	switch len(extras) {
	case 3:
		corruption = extras[2].(int)
		fallthrough
	case 2:
		width = extras[1].(int)
		fallthrough
	case 1:
		height = extras[0].(int)
	}
	return
}

func createTileGrid(
	lines []string,
	height int,
	width int,
	corruptionLimit int,
) *ds.Grid[tile] {
	grid := ds.NewGrid[tile](width, height)
	for y := range height {
		for x := range width {
			grid.Get(x, y).score = math.MaxInt
		}
	}
	for i := range corruptionLimit {
		setTileCorruption(grid, lines[i], true)
	}

	return grid
}

func setTileCorruption(grid *ds.Grid[tile], line string, corrupt bool) {
	x, y := util.CutToInts(line, ",")
	grid.Get(x, y).corrupted = corrupt
}

func getFewestStepsToTraverse(grid *ds.Grid[tile]) int {
	queue := ds.Queue[mazePath]{}
	startTile := grid.Get(start.XY())
	startTile.score = 0

	queue = queue.Push(mazePath{start, startTile, 0})
	fewestSteps := math.MaxInt
	var item mazePath
	for !queue.IsEmpty() {
		item, queue = queue.Pop()
		if item.score >= fewestSteps {
			break
		} else if item.pos.X == grid.Width-1 && item.pos.Y == grid.Height-1 {
			fewestSteps = item.score
		} else {
			steps := item.score + 1
			for _, neighbor := range item.pos.Neighbors() {
				if !grid.IsWithinBounds(neighbor.XY()) {
					continue
				}

				nextTile := grid.Get(neighbor.XY())
				if nextTile.corrupted || steps >= nextTile.score {
					continue
				}
				nextTile.score = steps
				queue = queue.Push(mazePath{neighbor, nextTile, steps})
			}
		}
	}

	return fewestSteps
}
