package day15

import (
	"fmt"
	"slices"
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
)

func PartOne(lines []string, extras ...any) any {
	return getBoxCoordinateSum(lines, false)
}

func PartTwo(lines []string, extras ...any) any {
	return getBoxCoordinateSum(lines, true)
}

func getBoxCoordinateSum(lines []string, wide bool) int {
	warehouse, robot := createWarehouse(lines, wide)

	for _, line := range lines[len(warehouse)+1:] {
		for _, op := range line {
			dir := getDirectionFromOperation(op)
			next := robot.Add(dir)
			nextSymbol := warehouse[next.Y][next.X]
			switch nextSymbol {
			case '.':
				robot = next
			case 'O', '[', ']':
				if moveBox(warehouse, next, dir, wide) {
					robot = next
				}
			}
		}
	}

	var boxSymbol uint8 = 'O'
	if wide {
		boxSymbol = '['
	}
	coordinateSum := 0
	for y, line := range warehouse {
		for x, symbol := range line {
			if symbol == boxSymbol {
				coordinateSum += y*100 + x
			}
		}
	}
	return coordinateSum
}

func moveBox(
	warehouse [][]uint8,
	start geom.Point,
	dir geom.Direction,
	wide bool,
) bool {
	var boxPoints []geom.Point
	if wide && (dir == geom.Up || dir == geom.Down) {
		visited := ds.NewSet[geom.Point]()
		queue := ds.Queue[geom.Point]{}

		var point geom.Point
		queue = queue.Push(start)
		for !queue.IsEmpty() {
			point, queue = queue.Pop()
			if !visited.Add(point) {
				continue
			}
			boxPoints = append(boxPoints, point)

			if warehouse[point.Y][point.X] == '[' {
				queue = queue.Push(point.Add(geom.Right))
			} else {
				queue = queue.Push(point.Add(geom.Left))
			}

			next := point.Add(dir)
			switch warehouse[next.Y][next.X] {
			case '[', ']':
				queue = queue.Push(next)
			case '#':
				return false
			}
		}
	} else {
		symbol := warehouse[start.Y][start.X]
		current := start
		for symbol != '#' && symbol != '.' {
			boxPoints = append(boxPoints, current)
			current = current.Add(dir)
			symbol = warehouse[current.Y][current.X]
		}
		if symbol == '#' {
			return false
		}
	}

	for _, point := range slices.Backward(boxPoints) {
		next := point.Add(dir)
		warehouse[next.Y][next.X] = warehouse[point.Y][point.X]
		warehouse[point.Y][point.X] = '.'
	}
	return true
}

func createWarehouse(lines []string, wide bool) (warehouse [][]uint8, start geom.Point) {
	maxX := len(lines[0])
	maxY := 1
	for lines[maxY] != "" {
		maxY += 1
	}
	if wide {
		maxX *= 2
	}

	warehouse = ds.Allocate2DSlice[uint8](maxX, maxY)

	if wide {
		for y, line := range lines[:maxY] {
			row := warehouse[y]
			for x, symbol := range line {
				x *= 2
				switch symbol {
				case '.':
					row[x] = '.'
					row[x+1] = '.'
				case 'O':
					row[x] = '['
					row[x+1] = ']'
				case '#':
					row[x] = '#'
					row[x+1] = '#'
				case '@':
					// Ignore the robot symbol and just track its position.
					row[x] = '.'
					row[x+1] = '.'
					start = geom.Point{X: x, Y: y}
				}
			}
		}
	} else {
		startFound := false
		for y, line := range lines[:maxY] {
			copy(warehouse[y], line)
			if !startFound {
				startX := strings.IndexRune(line, '@')
				if startX != -1 {
					startFound = true
					warehouse[y][startX] = '.'
					start = geom.Point{X: startX, Y: y}
				}
			}
		}
	}

	return warehouse, start
}

func getDirectionFromOperation(op rune) geom.Direction {
	switch op {
	case '^':
		return geom.Up
	case '>':
		return geom.Right
	case 'v':
		return geom.Down
	case '<':
		return geom.Left
	default:
		panic(fmt.Errorf("unknown op %c", op))
	}
}
