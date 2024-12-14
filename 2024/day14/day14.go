package day14

import (
	"iter"
	"strings"

	"github.com/chwallen/advent-of-code/internal/geom"
	"github.com/chwallen/advent-of-code/internal/util"
)

var (
	gridHeight = 103
	gridWidth  = 101
)

type robot struct {
	position geom.Point
	vector   geom.Point
}

func PartOne(lines []string, extras ...any) any {
	height, width := getParameters(extras)
	var finalState []robot
	for _, finalState = range simulateRobotMovements(lines, height, width, 100) {
	}

	topLeftQuadrant := 0
	topRightQuadrant := 0
	bottomRightQuadrant := 0
	bottomLeftQuadrant := 0
	for _, r := range finalState {
		x, y := r.position.XY()
		if x > width/2 {
			if y > height/2 {
				bottomRightQuadrant += 1
			} else if y < height/2 {
				topRightQuadrant += 1
			}
		} else if x < width/2 {
			if y > height/2 {
				bottomLeftQuadrant += 1
			} else if y < height/2 {
				topLeftQuadrant += 1
			}
		}
	}
	return topLeftQuadrant * topRightQuadrant * bottomRightQuadrant * bottomLeftQuadrant
}

func PartTwo(lines []string, extras ...any) any {
	height, width := getParameters(extras)
	positions := make([]bool, height*width)
	for second, robots := range simulateRobotMovements(lines, height, width, height*width) {
		if hasNoRobotOverlaps(robots, height, width, positions) {
			return second
		}
		clear(positions)
	}
	panic("simulation did not find a christmas tree")
}

func getParameters(extras []any) (height, width int) {
	if len(extras) == 2 {
		return extras[0].(int), extras[1].(int)
	}
	return gridHeight, gridWidth
}

func hasNoRobotOverlaps(
	robots []robot,
	_ int,
	width int,
	positions []bool,
) bool {
	for _, robot := range robots {
		x, y := robot.position.XY()
		i := y*width + x
		if positions[i] {
			return false
		}
		positions[i] = true
	}
	return true
}

func simulateRobotMovements(
	lines []string,
	height int,
	width int,
	seconds int,
) iter.Seq2[int, []robot] {
	robots := make([]robot, 0, 500)

	for _, line := range lines {
		position, velocity, _ := strings.Cut(line, " ")
		pX, pY := util.CutToInts(strings.TrimPrefix(position, "p="), ",")
		vX, vY := util.CutToInts(strings.TrimPrefix(velocity, "v="), ",")

		robots = append(robots, robot{
			position: geom.Point{X: pX, Y: pY},
			vector:   geom.Point{X: vX, Y: vY},
		})
	}

	return func(yield func(int, []robot) bool) {
		for second := 1; second <= seconds; second++ {
			for i, r := range robots {
				pX, pY := r.position.XY()
				vX, vY := r.vector.XY()
				nextX := (pX + vX + width) % width
				nextY := (pY + vY + height) % height
				robots[i].position = geom.Point{X: nextX, Y: nextY}
			}
			if !yield(second, robots) {
				return
			}
		}
	}
}
