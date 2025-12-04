package day04

import (
	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
)

var directions = []geom.Direction{
	geom.UpLeft,
	geom.Up,
	geom.UpRight,
	geom.Left,
	geom.Right,
	geom.DownLeft,
	geom.Down,
	geom.DownRight,
}

func PartOne(lines []string, extras ...any) any {
	grid := createGrid(lines)
	return countRollsToRemove(grid, false)
}

func PartTwo(lines []string, extras ...any) any {
	grid := createGrid(lines)
	rolls := 0
	for {
		removed := countRollsToRemove(grid, true)
		if removed == 0 {
			break
		}
		rolls += removed
	}
	return rolls
}

func createGrid(lines []string) [][]byte {
	grid := ds.Allocate2DSlice[byte](len(lines[0]), len(lines))
	for y, line := range lines {
		copy(grid[y], line)
	}
	return grid
}

func countRollsToRemove(grid [][]byte, remove bool) int {
	rolls := 0
	height, width := len(grid), len(grid[0])
	for y, line := range grid {
		for x, char := range line {
			if char != '@' {
				continue
			}

			adjacentItems := 0
			for _, dir := range directions {
				xn := x + dir.X
				yn := y + dir.Y
				if 0 <= yn && yn < height && 0 <= xn && xn < width && grid[yn][xn] == '@' {
					adjacentItems++
				}
			}
			if adjacentItems < 4 {
				rolls++
				if remove {
					grid[y][x] = '.'
				}
			}
		}
	}

	return rolls
}
