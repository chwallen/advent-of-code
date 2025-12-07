package day07

import (
	"strings"
)

func PartOne(lines []string, extras ...any) any {
	splits, _ := traceTachyon(lines)
	return splits
}

func PartTwo(lines []string, extras ...any) any {
	_, timelines := traceTachyon(lines)
	return timelines
}

func traceTachyon(lines []string) (splits, timelines int) {
	l := len(lines[0])
	prevCount := make([]int, l)
	currCount := make([]int, l)

	start := strings.Index(lines[0], "S")
	prevCount[start] = 1
	for _, line := range lines[1:] {
		for x, prev := range prevCount {
			if prev == 0 {
				continue
			}

			if line[x] == '^' {
				splits++
				if x > 0 {
					currCount[x-1] += prev
				}
				if x < l-1 {
					currCount[x+1] += prev
				}
			} else {
				currCount[x] += prev
			}
		}

		prevCount, currCount = currCount, prevCount
		clear(currCount)
	}

	for _, v := range prevCount {
		timelines += v
	}
	return splits, timelines
}
