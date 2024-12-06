package day06

import (
	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
)

type patrolTile struct {
	char        rune
	visited     bool
	visitedDirs [4]bool
}

func PartOne(lines []string, extras ...any) any {
	grid, start := createGrid(lines)
	return len(patrolGrid(grid, start))
}

func PartTwo(lines []string, extras ...any) any {
	initialGrid, start := createGrid(lines)
	loopGrid := initialGrid.Clone()

	patrolledPath := patrolGrid(loopGrid, start)

	loops := 0
	// Skip the starting tile
	for _, point := range patrolledPath[1:] {
		loopGrid.CopyFrom(initialGrid)
		loopGrid.Get(point.XY()).char = '#'
		if isLoop(loopGrid, start) {
			loops += 1
		}
	}

	return loops
}

func createGrid(lines []string) (grid *ds.Grid[patrolTile], start geom.Point) {
	grid = ds.NewGrid[patrolTile](len(lines[0]), len(lines))

	for y, line := range lines {
		for x, char := range line {
			grid.Get(x, y).char = char
			if char == '^' {
				start = geom.Point{X: x, Y: y}
			}
		}
	}
	return grid, start
}

func patrolGrid(grid *ds.Grid[patrolTile], start geom.Point) []geom.Point {
	current := start
	dir := geom.Up
	path := make([]geom.Point, 0, 2000)

	for {
		tile := grid.Get(current.XY())
		if !tile.visited {
			tile.visited = true
			path = append(path, current)
		}

		next := current.Add(dir)
		if !grid.IsWithinBounds(next.XY()) {
			return path
		}

		if grid.Get(next.XY()).char != '#' {
			current = next
		} else {
			dir = dir.TurnRight()
		}
	}
}

func isLoop(grid *ds.Grid[patrolTile], start geom.Point) bool {
	current := start
	dir := geom.Up
	cardinalIndex := dir.GetCardinalIndex()

	for {
		tile := grid.Get(current.XY())
		if tile.visitedDirs[cardinalIndex] {
			return true
		}

		tile.visitedDirs[cardinalIndex] = true
		next := current.Add(dir)
		if !grid.IsWithinBounds(next.XY()) {
			return false
		}

		if grid.Get(next.XY()).char != '#' {
			current = next
		} else {
			dir = dir.TurnRight()
			cardinalIndex = dir.GetCardinalIndex()
		}
	}
}
