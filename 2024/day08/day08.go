package day08

import (
	"math"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
)

func PartOne(lines []string, extras ...any) any {
	return getUniqueAntiNodesCount(lines, false)
}

func PartTwo(lines []string, extras ...any) any {
	return getUniqueAntiNodesCount(lines, true)
}

func getUniqueAntiNodesCount(lines []string, unlimitedDistance bool) int {
	antennas := make(map[rune][]geom.Point)
	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			p := geom.Point{X: x, Y: y}
			antennas[char] = append(antennas[char], p)
		}
	}

	antiNodes := ds.NewSet[geom.Point]()
	maxY := len(lines)
	maxX := len(lines[0])
	start := 1
	end := 1
	if unlimitedDistance {
		start = 0
		end = math.MaxInt
	}

	for _, points := range antennas {
		for i, pointA := range points[:len(points)-1] {
			for _, pointB := range points[i+1:] {
				ax, ay := pointA.XY()
				bx, by := pointB.XY()
				dx := ax - bx
				dy := ay - by

				for k := start; k <= end; k++ {
					antiNodeOne := geom.Point{X: ax + k*dx, Y: ay + k*dy}
					antiNodeTwo := geom.Point{X: bx - k*dx, Y: by - k*dy}

					nodeOneInBounds := antiNodeOne.IsWithinBounds(0, 0, maxX, maxY)
					nodeTwoInBounds := antiNodeTwo.IsWithinBounds(0, 0, maxX, maxY)

					if nodeOneInBounds {
						antiNodes.Add(antiNodeOne)
					}
					if nodeTwoInBounds {
						antiNodes.Add(antiNodeTwo)
					}
					if !nodeOneInBounds && !nodeTwoInBounds {
						break
					}
				}
			}
		}
	}

	return len(antiNodes)
}
