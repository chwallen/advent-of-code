package day10

import (
	"fmt"
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/util"
	"github.com/draffensperger/golp"
)

type lightsState struct {
	current int
	presses int
}

func PartOne(lines []string, extras ...any) any {
	minimumRequiredPresses := 0
Machines:
	for _, line := range lines {
		parts := strings.Fields(line)

		sought := parseLightIndicators(parts[0])
		buttons := parseLightsButtons(parts)

		seen := make(map[int]int)
		var queue ds.Queue[lightsState]
		queue = queue.Push(lightsState{
			current: 0,
			presses: 0,
		})
		for {
			var s lightsState
			s, queue = queue.Pop()
			s.presses++
			for _, button := range buttons {
				next := s.current ^ button
				if next == sought {
					minimumRequiredPresses += s.presses
					continue Machines
				}

				if v, ok := seen[next]; ok && v <= s.presses {
					continue
				}
				seen[next] = s.presses

				queue = queue.Push(lightsState{
					current: next,
					presses: s.presses,
				})
			}
		}
	}
	return minimumRequiredPresses
}

func PartTwo(lines []string, extras ...any) any {
	minimumRequiredPresses := 0
	for _, line := range lines {
		parts := strings.Fields(line)

		sought := parseJoltage(parts[len(parts)-1])
		buttons := parseJoltageButtons(parts, len(sought))
		n := len(buttons)

		lp := golp.NewLP(0, n)

		obj := make([]float64, n)
		for i := range n {
			obj[i] = 1.0
			lp.SetInt(i, true)
		}
		lp.SetObjFn(obj)

		for i, s := range sought {
			row := make([]float64, n)
			for j := range n {
				row[j] = float64(buttons[j][i])
			}
			_ = lp.AddConstraint(row, golp.EQ, float64(s))
		}

		if lp.Solve() != golp.OPTIMAL {
			panic(fmt.Errorf("failed to find optimal solution for problem %s", line))
		}

		for _, v := range lp.Variables() {
			minimumRequiredPresses += int(v)
		}
	}
	return minimumRequiredPresses
}

func parseLightIndicators(indicators string) int {
	var sought int
	for i, light := range indicators[1 : len(indicators)-1] {
		active := 0
		if light == '#' {
			active = 1
		}
		sought |= active << i
	}
	return sought
}

func parseLightsButtons(parts []string) []int {
	parts = parts[1 : len(parts)-1]
	buttons := make([]int, len(parts))
	for i, group := range parts {
		for s := range strings.SplitSeq(group[1:len(group)-1], ",") {
			buttons[i] |= 1 << util.Atoi(s)
		}
	}
	return buttons
}

func parseJoltageButtons(parts []string, n int) [][]int {
	parts = parts[1 : len(parts)-1]
	buttons := ds.Allocate2DSlice[int](n, len(parts))
	for i, group := range parts {
		for s := range strings.SplitSeq(group[1:len(group)-1], ",") {
			buttons[i][util.Atoi(s)] = 1
		}
	}
	return buttons
}

func parseJoltage(part string) []int {
	var joltageRequirements []int
	for v := range strings.SplitSeq(part[1:len(part)-1], ",") {
		joltageRequirements = append(joltageRequirements, util.Atoi(v))
	}
	return joltageRequirements
}
