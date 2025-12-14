package day10

import (
	"math"
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

type buttonCombination struct {
	changes []int
	presses int
}

type machine struct {
	sought             []int
	buttonCombinations []buttonCombination
}

func PartOne(lines []string, extras ...any) any {
	minimumRequiredPresses := 0
	for _, line := range lines {
		m := parseMachine(line, true)
		presses := math.MaxInt

	Combinations:
		for _, combination := range m.buttonCombinations {
			for i, v := range combination.changes {
				if v%2 != m.sought[i] {
					continue Combinations
				}
			}
			presses = min(presses, combination.presses)
		}
		minimumRequiredPresses += presses
	}
	return minimumRequiredPresses
}

func PartTwo(lines []string, extras ...any) any {
	pressesChannel := make(chan int, len(lines))
	solve := func(line string) {
		m := parseMachine(line, false)
		pressesChannel <- findFewestPresses(m.sought, m.buttonCombinations)
	}

	for _, line := range lines {
		go solve(line)
	}

	minimumRequiredPresses := 0
	for range len(lines) {
		minimumRequiredPresses += <-pressesChannel
	}
	return minimumRequiredPresses
}

func findFewestPresses(sought []int, combinations []buttonCombination) int {
	if isZeroSlice(sought) {
		return 0
	}

	next := make([]int, len(sought))
	presses := math.MaxInt
Combinations:
	for _, combination := range combinations {
		for i, v := range sought {
			c := combination.changes[i]
			if c > v || c%2 != v%2 {
				continue Combinations
			}
		}

		for i, v := range sought {
			next[i] = (v - combination.changes[i]) / 2
		}

		n := findFewestPresses(next, combinations)
		if n == 0 {
			presses = min(presses, combination.presses)
		} else if n < math.MaxInt {
			presses = min(presses, 2*n+combination.presses)
		}
	}

	return presses
}

func isZeroSlice(s []int) bool {
	for _, v := range s {
		if v > 0 {
			return false
		}
	}
	return true
}

func parseMachine(line string, parseLights bool) machine {
	parts := strings.Fields(line)
	n := len(parts[0]) - 2

	sought := make([]int, 0, n)
	if parseLights {
		part := parts[0]
		for _, light := range part[1 : n+1] {
			active := 0
			if light == '#' {
				active = 1
			}
			sought = append(sought, active)
		}
	} else {
		part := parts[len(parts)-1]
		for v := range strings.SplitSeq(part[1:len(part)-1], ",") {
			sought = append(sought, util.Atoi(v))
		}
	}

	buttonParts := parts[1 : len(parts)-1]
	buttonsCount := len(buttonParts)
	buttons := make([][]int, buttonsCount)
	for i, group := range buttonParts {
		buttons[i] = make([]int, 0, len(group))
		for s := range strings.SplitSeq(group[1:len(group)-1], ",") {
			buttons[i] = append(buttons[i], util.Atoi(s))
		}
	}

	combinationsCount := 1 << buttonsCount
	buttonCombinations := make([]buttonCombination, 0, combinationsCount)
	combinations := make([]int, combinationsCount*n)
	for i := range combinationsCount {
		start := i * n
		changes := combinations[start : start+n]
		presses := 0
		for j := range buttonsCount {
			if i&(1<<j) != 0 {
				presses++
				for _, index := range buttons[j] {
					changes[index]++
				}
			}
		}

		buttonCombinations = append(buttonCombinations, buttonCombination{
			changes: changes,
			presses: presses,
		})
	}

	return machine{
		sought:             sought,
		buttonCombinations: buttonCombinations,
	}
}
