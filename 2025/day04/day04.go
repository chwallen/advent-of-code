package day04

import (
	"github.com/chwallen/advent-of-code/internal/ds"
)

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

func createGrid(lines []string) [][]rune {
	grid := ds.Allocate2DSlice[rune](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, char := range line {
			grid[y][x] = char
		}
	}
	return grid
}

func countRollsToRemove(grid [][]rune, remove bool) int {
	rolls := 0
	height, width := len(grid), len(grid[0])
	for y, line := range grid {
		for x, char := range line {
			if char == '.' {
				continue
			}

			adjacentItems := 0
			if y > 0 {
				if x > 0 && grid[y-1][x-1] == '@' {
					adjacentItems += 1
				}
				if grid[y-1][x] == '@' {
					adjacentItems += 1
				}
				if x < width-1 && grid[y-1][x+1] == '@' {
					adjacentItems += 1
				}
			}
			if x > 0 && grid[y][x-1] == '@' {
				adjacentItems += 1
			}
			if x < width-1 && grid[y][x+1] == '@' {
				adjacentItems += 1
			}
			if y < height-1 {
				if x > 0 && grid[y+1][x-1] == '@' {
					adjacentItems += 1
				}
				if grid[y+1][x] == '@' {
					adjacentItems += 1
				}
				if x < width-1 && grid[y+1][x+1] == '@' {
					adjacentItems += 1
				}
			}

			if adjacentItems < 4 {
				rolls += 1
				if remove {
					grid[y][x] = '.'
				}
			}
		}
	}

	return rolls
}
