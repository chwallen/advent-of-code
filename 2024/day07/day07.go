package day07

import (
	"strconv"
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

type operation int

const (
	add operation = iota
	mul
	concat
)

func PartOne(lines []string, extras ...any) any {
	return getSumOfValidTestValues(lines, add, mul)
}

func PartTwo(lines []string, extras ...any) any {
	return getSumOfValidTestValues(lines, add, mul, concat)
}

func getSumOfValidTestValues(lines []string, ops ...operation) int {
	sum := 0
	values := make([]int, 0, 50)
	for _, line := range lines {
		left, right, _ := strings.Cut(line, ": ")
		expected := util.Atoi(left)
		values = util.SplitToInts(right, " ", values)

		if hasSolution(values, len(values)-1, expected, ops) {
			sum += expected
		}

		values = values[:0]
	}
	return sum
}

// Tries to find the solution backwards as it's much faster to prune invalid
// branches that way.
func hasSolution(values []int, i, current int, ops []operation) bool {
	value := values[i]
	if i == 0 {
		return current == value
	}

	for _, op := range ops {
		switch op {
		case add:
			v := current - value
			if v >= values[i-1] && hasSolution(values, i-1, v, ops) {
				return true
			}
		case mul:
			v, remainder := util.DivRem(current, value)
			if remainder == 0 && hasSolution(values, i-1, v, ops) {
				return true
			}
		case concat:
			lhs := strconv.Itoa(current)
			rhs := strconv.Itoa(value)
			if strings.HasSuffix(lhs, rhs) && len(lhs) > len(rhs) {
				v := util.Atoi(lhs[:len(lhs)-len(rhs)])
				if hasSolution(values, i-1, v, ops) {
					return true
				}
			}
		}
	}
	return false
}
