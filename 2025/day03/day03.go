package day03

import (
	"github.com/chwallen/advent-of-code/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	return getJoltageForLines(lines, 2)
}

func PartTwo(lines []string, extras ...any) any {
	return getJoltageForLines(lines, 12)
}

func getJoltageForLines(lines []string, batteriesToActivate int) int {
	sum := 0
	for _, line := range lines {
		start := 0
		stop := len(line) - batteriesToActivate + 1
		remaining := batteriesToActivate
		for remaining > 0 {
			battery, offset := getNextBattery(line[start:stop])
			sum += battery * util.IntPow(10, remaining-1)
			start += offset + 1
			stop++
			remaining--
		}
	}

	return sum
}

func getNextBattery(line string) (battery, index int) {
	for i, char := range line {
		v := int(char - '0')
		if v > battery {
			battery = v
			index = i
		}
	}
	return battery, index
}
