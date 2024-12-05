package day05

import (
	"slices"

	"github.com/chwallen/advent-of-code/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	return getSumOfGoodUpdates(lines, slices.Equal)
}

func PartTwo(lines []string, extras ...any) any {
	return getSumOfGoodUpdates(lines, notEqual)
}

func notEqual(update, fixedUpdate []int) bool {
	return !slices.Equal(update, fixedUpdate)
}

func getSumOfGoodUpdates(
	lines []string,
	isGoodUpdate func(update, fixedUpdate []int) bool,
) int {
	rules := make(map[int][]int)

	var i int
	for i = 0; lines[i] != ""; i++ {
		precedingPage, subsequentPage := util.CutToInts(lines[i], "|")
		rules[precedingPage] = append(rules[precedingPage], subsequentPage)
	}

	fixUpdate := func(pageA, pageB int) int {
		if slices.Contains(rules[pageA], pageB) {
			return -1
		}
		return 1
	}

	update := make([]int, 0, 30)
	updateCopy := make([]int, 30)
	sum := 0
	for _, line := range lines[i+1:] {
		update = util.SplitToInts(line, ",", update)
		copy(updateCopy, update)

		fixedUpdate := updateCopy[:len(update)]
		slices.SortFunc(fixedUpdate, fixUpdate)

		if isGoodUpdate(update, fixedUpdate) {
			sum += fixedUpdate[len(fixedUpdate)/2]
		}

		update = update[:0]
	}
	return sum
}
