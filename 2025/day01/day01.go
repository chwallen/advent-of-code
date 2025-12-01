package day01

import (
	"fmt"

	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	dialStart     = 50
	dialPositions = 100
)

func PartOne(lines []string, extras ...any) any {
	dial := dialStart
	zeros := 0
	for _, line := range lines {
		dir := line[0]
		steps := util.Atoi(line[1:])

		steps %= dialPositions
		switch dir {
		case 'R':
			dial += steps
			dial %= dialPositions
		case 'L':
			dial -= steps
			if dial < 0 {
				dial += dialPositions
			}
		default:
			panic(fmt.Errorf("unexpected direction %c", dir))
		}

		if dial == 0 {
			zeros += 1
		}
	}

	return zeros
}

func PartTwo(lines []string, extras ...any) any {
	dial := dialStart
	zeros := 0
	for _, line := range lines {
		dir := line[0]
		steps := util.Atoi(line[1:])

		if steps >= dialPositions {
			zeros += steps / dialPositions
			steps %= dialPositions
		}

		if steps == 0 {
			continue
		}

		switch dir {
		case 'R':
			dial += steps
			if dial >= dialPositions {
				if dial-steps != 0 {
					zeros += 1
				}
				dial -= dialPositions
			}
		case 'L':
			dial -= steps
			if dial <= 0 {
				if dial+steps != 0 {
					zeros += 1
				}
				if dial < 0 {
					dial += dialPositions
				}
			}
		default:
			panic(fmt.Errorf("unexpected direction %c", dir))
		}
	}

	return zeros
}
