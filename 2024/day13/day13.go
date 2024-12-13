package day13

import (
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	prizePointCoordinateModifier = 10_000_000_000_000
)

func PartOne(lines []string, extras ...any) any {
	return getFewestTokensToWinAllPrizes(lines, 0)
}

func PartTwo(lines []string, extras ...any) any {
	return getFewestTokensToWinAllPrizes(lines, prizePointCoordinateModifier)
}

func parseLine(line string) (int, int) {
	_, xy, _ := strings.Cut(line, ": ")
	x, y, _ := strings.Cut(xy, ", ")
	return util.Atoi(x[2:]), util.Atoi(y[2:])
}

func getFewestTokensToWinAllPrizes(lines []string, modifier int) int {
	tokens := 0
	for i := 0; i < len(lines); i += 4 {
		ax, ay := parseLine(lines[i])
		bx, by := parseLine(lines[i+1])
		px, py := parseLine(lines[i+2])
		px += modifier
		py += modifier

		// Let ax and ay be button A's x and y values, respectively, and let bx and
		// by be the equivalent for button B. Let px and py be the equivalent for
		// the prize coordinates. You get the system of equations where x and y are
		// the unknowns:
		//
		//  ax * x + bx * y = px (1)
		//  ay * x + by * y = py (2)
		//
		// We solve for x by:
		//  1. multiplying (1) with by and (2) with bx
		//  2. subtract (2) from (1)
		//  3. divide (1) by its left-hand side
		//
		// The expression for x will be x = (px*by*y - py*bx*y) / (ax*by*y - ay*bx*y)
		x := (px*by - py*bx) / (ax*by - ay*bx)
		// Solve for y in (1) by:
		//  1. subtract ax*x from (1)
		//  2. divide (1) by bx
		//  3. substitute x from above
		y := (px - ax*x) / bx
		// Check that x and y are valid solutions to the system.
		if x*ax+y*bx == px && x*ay+y*by == py {
			tokens += 3*x + y
		}
	}
	return tokens
}
