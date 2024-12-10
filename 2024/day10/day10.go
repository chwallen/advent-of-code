package day10

import (
	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
)

type trail struct {
	x0, y0, x, y int
}

func PartOne(lines []string, extras ...any) any {
	uniqueTrails, _ := hikeTrails(lines)
	return len(uniqueTrails)
}

func PartTwo(lines []string, extras ...any) any {
	sum := 0
	_, trailRatings := hikeTrails(lines)
	for _, rating := range trailRatings {
		sum += rating
	}
	return sum
}

func hikeTrails(
	lines []string,
) (uniqueTrails ds.Set[trail], trailRatings map[geom.Point]int) {
	maxY := len(lines)
	maxX := len(lines[0])

	uniqueTrails = ds.NewSet[trail]()
	trailRatings = make(map[geom.Point]int, 1000)

	for y, line := range lines {
		for x, char := range line {
			if char == '0' {
				trailStart := geom.Point{X: x, Y: y}
				travelTrail(lines, trailStart, trailStart, maxX, maxY, uniqueTrails, trailRatings, '1')
			}
		}
	}

	return uniqueTrails, trailRatings
}

func travelTrail(
	lines []string,
	trailStart geom.Point,
	previous geom.Point,
	maxX int,
	maxY int,
	uniqueTrails ds.Set[trail],
	trailRatings map[geom.Point]int,
	soughtSymbol uint8,
) {
	for _, neighbor := range previous.Neighbors() {
		x, y := neighbor.XY()
		if neighbor.IsWithinBounds(0, 0, maxX, maxY) && lines[y][x] == soughtSymbol {
			if soughtSymbol == '9' {
				uniqueTrails.Add(trail{x0: trailStart.X, y0: trailStart.Y, x: x, y: y})
				trailRatings[neighbor] += 1
			} else {
				travelTrail(lines, trailStart, neighbor, maxX, maxY, uniqueTrails, trailRatings, soughtSymbol+1)
			}
		}
	}
}
