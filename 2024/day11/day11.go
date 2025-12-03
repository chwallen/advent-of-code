package day11

import (
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	partOneGenerations = 25
	partTwoGenerations = 75
)

var powersOf10 = [19]int{
	1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000,
	1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000,
	100000000000000, 1000000000000000, 10000000000000000, 100000000000000000,
	1000000000000000000,
}

func PartOne(lines []string, extras ...any) any {
	return getNumberOfStones(lines, partOneGenerations)
}

func PartTwo(lines []string, extras ...any) any {
	return getNumberOfStones(lines, partTwoGenerations)
}

func getNumberOfStones(lines []string, generations int) int {
	line := lines[0]
	currentGeneration := make(map[int]int, 1000)
	nextGeneration := make(map[int]int, 1000)

	for inputStone := range strings.FieldsSeq(line) {
		stone := util.Atoi(inputStone)
		currentGeneration[stone] += 1
	}

	for range generations {
		for stoneNumber, count := range currentGeneration {
			if stoneNumber == 0 {
				nextGeneration[1] += count
				continue
			}

			digits := util.CountDigits(stoneNumber)
			half, remainder := util.DivRem(digits, 2)
			if remainder == 0 {
				divisor := powersOf10[half]
				firstHalf, secondHalf := util.DivRem(stoneNumber, divisor)

				nextGeneration[firstHalf] += count
				nextGeneration[secondHalf] += count
			} else {
				nextGeneration[stoneNumber*2024] += count
			}
		}
		currentGeneration, nextGeneration = nextGeneration, currentGeneration
		clear(nextGeneration)
	}

	stones := 0
	for _, count := range currentGeneration {
		stones += count
	}
	return stones
}
