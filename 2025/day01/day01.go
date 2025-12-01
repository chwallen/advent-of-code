package day01

import (
	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	dialStart     = 50
	dialPositions = 100
)

func PartOne(lines []string, extras ...any) any {
	onZero, _ := turnDial(lines)
	return onZero
}

func PartTwo(lines []string, extras ...any) any {
	_, zeroes := turnDial(lines)
	return zeroes
}

func turnDial(lines []string) (onZero, zeroes int) {
	dial := dialStart
	for _, line := range lines {
		dir := line[0]
		steps := util.Atoi(line[1:])

		if steps >= dialPositions {
			zeroes += steps / dialPositions
			steps %= dialPositions
		}

		modifier := 1
		if dir == 'L' {
			modifier = -1
		}
		for range steps {
			dial += modifier
			if dial == 0 || dial == dialPositions {
				zeroes++
			}
		}
		if dial == 0 || dial == dialPositions {
			onZero++
			dial = 0
		} else if dial < 0 || dial > dialPositions {
			dial -= dialPositions * modifier
		}
	}

	return onZero, zeroes
}
