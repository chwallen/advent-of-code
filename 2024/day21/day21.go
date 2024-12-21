package day21

import (
	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	partOneRobots = 3
	partTwoRobots = 26
)

type pathKey struct {
	start, end rune
}

func PartOne(lines []string, extras ...any) any {
	return calculateComplexitySum(lines, partOneRobots)
}

func PartTwo(lines []string, extras ...any) any {
	return calculateComplexitySum(lines, partTwoRobots)
}

func calculateComplexitySum(lines []string, robots int) int {
	robotPathCaches := make([]map[string]int, robots+1)
	for i := range len(robotPathCaches) {
		robotPathCaches[i] = make(map[string]int)
	}
	sum := 0
	for _, line := range lines {
		numberPart := util.Atoi(line[0 : len(line)-1])
		length := getPathLength(line, robots, robotPathCaches)
		sum += length * numberPart
	}
	return sum
}

func getPathLength(
	path string,
	robot int,
	robotPathCaches []map[string]int,
) int {
	if cachedValue, exists := robotPathCaches[robot][path]; exists {
		return cachedValue
	}

	nextRobot := robot - 1
	length := 0
	current := 'A'
	for _, next := range path {
		// Pressing the same button repeatedly
		if current == next {
			length += 1
		} else {
			p := paths[pathKey{current, next}]
			if nextRobot == 0 {
				length += len(p)
			} else {
				length += getPathLength(p, nextRobot, robotPathCaches)
			}
		}
		current = next
	}

	robotPathCaches[robot][path] = length
	return length
}

// These are the most efficient paths from one symbol to the next, including
// pressing A.
var paths = map[pathKey]string{
	// Directional paths
	{'A', '^'}: "<A",
	{'A', '>'}: "vA",
	{'A', 'v'}: "<vA",
	{'A', '<'}: "v<<A",
	{'^', 'A'}: ">A",
	{'^', 'v'}: "vA",
	{'^', '>'}: "v>A",
	{'^', '<'}: "v<A",
	{'>', 'A'}: "^A",
	{'>', '^'}: "<^A",
	{'>', 'v'}: "<A",
	{'>', '<'}: "<<A",
	{'v', 'A'}: "^>A",
	{'v', '^'}: "^A",
	{'v', '>'}: ">A",
	{'v', '<'}: "<A",
	{'<', 'A'}: ">>^A",
	{'<', '^'}: ">^A",
	{'<', '>'}: ">>A",
	{'<', 'v'}: ">A",
	// Numerical paths
	// A -> x
	{'A', '0'}: "<A",
	{'A', '1'}: "^<<A",
	{'A', '2'}: "<^A",
	{'A', '3'}: "^A",
	{'A', '4'}: "^^<<A",
	{'A', '5'}: "<^^A",
	{'A', '6'}: "^^A",
	{'A', '7'}: "^^^<<A",
	{'A', '8'}: "<^^^A",
	{'A', '9'}: "^^^A",
	// 0 -> x
	{'0', 'A'}: ">A",
	{'0', '1'}: "^<A",
	{'0', '2'}: "^A",
	{'0', '3'}: "^>A",
	{'0', '4'}: "^^<A",
	{'0', '5'}: "^^A",
	{'0', '6'}: "^^>A",
	{'0', '7'}: "^^^<A",
	{'0', '8'}: "^^^A",
	{'0', '9'}: "^^^>A",
	// 1 -> x
	{'1', 'A'}: ">>vA",
	{'1', '0'}: ">vA",
	{'1', '2'}: ">A",
	{'1', '3'}: ">>A",
	{'1', '4'}: "^A",
	{'1', '5'}: "^>A",
	{'1', '6'}: "^>>A",
	{'1', '7'}: "^^A",
	{'1', '8'}: "^^>A",
	{'1', '9'}: "^^>>A",
	// 2 -> x
	{'2', 'A'}: "v>A",
	{'2', '0'}: "vA",
	{'2', '1'}: "<A",
	{'2', '3'}: ">A",
	{'2', '4'}: "<^A",
	{'2', '5'}: "^A",
	{'2', '6'}: "^>A",
	{'2', '7'}: "<^^A",
	{'2', '8'}: "^^A",
	{'2', '9'}: "^^>A",
	// 3 -> x
	{'3', 'A'}: "vA",
	{'3', '0'}: "<vA",
	{'3', '1'}: "<<A",
	{'3', '2'}: "<A",
	{'3', '4'}: "<<^A",
	{'3', '5'}: "<^A",
	{'3', '6'}: "^A",
	{'3', '7'}: "<<^^A",
	{'3', '8'}: "<^^A",
	{'3', '9'}: "^^A",
	// 4 -> x
	{'4', 'A'}: ">>vvA",
	{'4', '0'}: ">vvA",
	{'4', '1'}: "vA",
	{'4', '2'}: "v>A",
	{'4', '3'}: "v>>A",
	{'4', '5'}: ">A",
	{'4', '6'}: ">>A",
	{'4', '7'}: "^A",
	{'4', '8'}: "^>A",
	{'4', '9'}: "^>>A",
	// 5 -> x
	{'5', 'A'}: "vv>A",
	{'5', '0'}: "vvA",
	{'5', '1'}: "<vA",
	{'5', '2'}: "vA",
	{'5', '3'}: "v>A",
	{'5', '4'}: "<A",
	{'5', '6'}: ">A",
	{'5', '7'}: "<^A",
	{'5', '8'}: "^A",
	{'5', '9'}: "^>A",
	// 6 -> x
	{'6', 'A'}: "vvA",
	{'6', '0'}: "<vvA",
	{'6', '1'}: "<<vA",
	{'6', '2'}: "<vA",
	{'6', '3'}: "vA",
	{'6', '4'}: "<<A",
	{'6', '5'}: "<A",
	{'6', '7'}: "<<^A",
	{'6', '8'}: "<^A",
	{'6', '9'}: "^A",
	// 7 -> x
	{'7', 'A'}: ">>vvvA",
	{'7', '0'}: ">vvvA",
	{'7', '1'}: "vvA",
	{'7', '2'}: "vv>A",
	{'7', '3'}: "vv>>A",
	{'7', '4'}: "vA",
	{'7', '5'}: "v>A",
	{'7', '6'}: "v>>A",
	{'7', '8'}: ">A",
	{'7', '9'}: ">>A",
	// 8 -> x
	{'8', 'A'}: "vvv>A",
	{'8', '0'}: "vvvA",
	{'8', '1'}: "<vvA",
	{'8', '2'}: "vvA",
	{'8', '3'}: "vv>A",
	{'8', '4'}: "<vA",
	{'8', '5'}: "vA",
	{'8', '6'}: "v>A",
	{'8', '7'}: "<A",
	{'8', '9'}: ">A",
	// 9 -> x
	{'9', 'A'}: "vvvA",
	{'9', '0'}: "<vvvA",
	{'9', '1'}: "<<vvA",
	{'9', '2'}: "<vvA",
	{'9', '3'}: "vvA",
	{'9', '4'}: "<<vA",
	{'9', '5'}: "<vA",
	{'9', '6'}: "vA",
	{'9', '7'}: "<<A",
	{'9', '9'}: "<A",
}
