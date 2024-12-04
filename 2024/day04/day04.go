package day04

import (
	"github.com/chwallen/advent-of-code/internal/geom"
)

func PartOne(lines []string, extras ...any) any {
	maxY := len(lines) - 3
	maxX := len(lines[0]) - 3

	xmasAppearances := 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {
				if x < maxX {
					if isXmasSequence(lines, x, y, geom.Right) {
						xmasAppearances += 1
					}
					if y < maxY && isXmasSequence(lines, x, y, geom.DownRight) {
						xmasAppearances += 1
					}
					if y >= 3 && isXmasSequence(lines, x, y, geom.UpRight) {
						xmasAppearances += 1
					}
				}
				if x >= 3 {
					if isXmasSequence(lines, x, y, geom.Left) {
						xmasAppearances += 1
					}
					if y < maxY && isXmasSequence(lines, x, y, geom.DownLeft) {
						xmasAppearances += 1
					}
					if y >= 3 && isXmasSequence(lines, x, y, geom.UpLeft) {
						xmasAppearances += 1
					}
				}
				if y < maxY && isXmasSequence(lines, x, y, geom.Down) {
					xmasAppearances += 1
				}
				if y >= 3 && isXmasSequence(lines, x, y, geom.Up) {
					xmasAppearances += 1
				}
			}
		}
	}

	return xmasAppearances
}

func PartTwo(lines []string, extras ...any) any {
	xmasAppearances := 0
	for i, line := range lines[:len(lines)-2] {
		for j := range line[:len(line)-2] {
			if lines[i+1][j+1] == 'A' {
				topLeft := line[j]
				topRight := line[j+2]
				bottomLeft := lines[i+2][j]
				bottomRight := lines[i+2][j+2]

				if isMasOrSam(topLeft, bottomRight) && isMasOrSam(topRight, bottomLeft) {
					xmasAppearances += 1
				}
			}
		}
	}

	return xmasAppearances
}

func isXmasSequence(lines []string, x, y int, dir geom.Direction) bool {
	dx, dy := dir.X, dir.Y
	return lines[y+dy][x+dx] == 'M' &&
		lines[y+2*dy][x+2*dx] == 'A' &&
		lines[y+3*dy][x+3*dx] == 'S'
}

func isMasOrSam(a, b byte) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}
