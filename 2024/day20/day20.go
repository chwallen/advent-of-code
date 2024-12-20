package day20

import (
	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	partOneCheatTiles = 2
	partTwoCheatTiles = 20
)

type tile struct {
	point geom.Point
	steps int
	char  rune
	next  *tile
}

func PartOne(lines []string, extras ...any) any {
	return findNumberOfGoodCheats(
		lines,
		partOneCheatTiles,
		getMinStepsToSaveByCheating(extras),
	)
}

func PartTwo(lines []string, extras ...any) any {
	return findNumberOfGoodCheats(
		lines,
		partTwoCheatTiles,
		getMinStepsToSaveByCheating(extras),
	)
}

func getMinStepsToSaveByCheating(extras []any) int {
	if len(extras) == 1 {
		return extras[0].(int)
	}
	return 100
}

func findNumberOfGoodCheats(
	lines []string,
	maxCheatDistance int,
	threshold int,
) int {
	width, height := len(lines[0]), len(lines)
	grid := ds.NewGrid[tile](width, height)
	var start *tile

	for y, line := range lines {
		for x, char := range line {
			t := grid.Get(x, y)
			t.point.X = x
			t.point.Y = y
			t.char = char
			if char == 'S' {
				start = t
			}
		}
	}
	if start == nil {
		panic("could not find start tile")
	}

	steps := traverseMaze(grid, start)

	cheats := 0
	currentTile := start

	maxHeight := height - 1
	maxWidth := width - 1

	for i := 0; i < steps-threshold; i++ {
		cheats += findCheatsAroundTile(
			grid,
			currentTile,
			maxCheatDistance,
			threshold,
			maxHeight,
			maxWidth,
		)
		currentTile = currentTile.next
	}

	return cheats
}

func traverseMaze(grid *ds.Grid[tile], start *tile) int {
	currentTile := start
	steps := 0
	currentDirection := geom.Up

	if grid.Get(currentTile.point.Add(currentDirection).XY()).char == '#' {
		for _, dir := range [3]geom.Direction{geom.Right, geom.Down, geom.Left} {
			if grid.Get(currentTile.point.Add(dir).XY()).char != '#' {
				currentDirection = dir
				break
			}
		}
	}

	for currentTile.char != 'E' {
		nextTile, nextDirection := traverseToNextTile(grid, currentTile, currentDirection)
		steps += 1
		nextTile.steps = steps
		currentTile.next = nextTile
		currentTile = nextTile
		currentDirection = nextDirection
	}

	return steps
}

func traverseToNextTile(grid *ds.Grid[tile], start *tile, dir geom.Direction) (*tile, geom.Direction) {
	next := grid.Get(start.point.Add(dir).XY())
	if next.char != '#' {
		return next, dir
	}

	right := dir.TurnRight()
	next = grid.Get(start.point.Add(right).XY())
	if next.char != '#' {
		return next, right
	}

	left := dir.TurnLeft()
	return grid.Get(start.point.Add(left).XY()), left
}

// findCheatsAroundTile finds all cheats around t0 within the manhattan distance
// specified by maxDistance that save at least threshold steps.
func findCheatsAroundTile(
	grid *ds.Grid[tile],
	t0 *tile,
	maxDistance int,
	threshold int,
	maxHeight int,
	maxWidth int,
) int {
	x0, y0 := t0.point.XY()
	steps := t0.steps
	cheats := 0

	minY := max(1, y0-maxDistance)
	maxY := min(maxHeight, y0+maxDistance)
	for y := minY; y <= maxY; y++ {
		yDist := util.Abs(y - y0)
		remainingDist := maxDistance - yDist

		minX := max(1, x0-remainingDist)
		maxX := min(maxWidth, x0+remainingDist)
		for x := minX; x <= maxX; x++ {
			xDist := util.Abs(x - x0)
			manhattanDistance := yDist + xDist

			if manhattanDistance == 0 {
				continue
			}

			t := grid.Get(x, y)
			if t.char == '#' {
				continue
			}

			savedSteps := t.steps - steps - manhattanDistance
			if savedSteps >= threshold {
				cheats += 1
			}
		}
	}
	return cheats
}
