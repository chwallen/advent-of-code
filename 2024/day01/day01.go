package day01

import (
	"slices"

	"github.com/chwallen/advent-of-code/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	linesCount := len(lines)
	data := make([]int, 2*linesCount)
	leftColumn := data[:linesCount]
	rightColumn := data[linesCount:]

	for i, line := range lines {
		leftColumn[i], rightColumn[i] = util.CutToInts(line, "   ")
	}

	slices.Sort(leftColumn)
	slices.Sort(rightColumn)

	distance := 0
	for i := range linesCount {
		distance += util.Abs(leftColumn[i] - rightColumn[i])
	}

	return distance
}

func PartTwo(lines []string, extras ...any) any {
	linesLen := len(lines)
	leftColumn := make([]int, linesLen)
	rightColumn := make(map[int]int, linesLen)

	for i, line := range lines {
		leftItem, rightItem := util.CutToInts(line, "   ")
		leftColumn[i] = leftItem
		rightColumn[rightItem] += 1
	}

	similarity := 0
	for _, leftItem := range leftColumn {
		similarity += leftItem * rightColumn[leftItem]
	}

	return similarity
}
