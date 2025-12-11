package day10_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/chwallen/advent-of-code/2025/day10"
	"github.com/chwallen/advent-of-code/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	exampleInput := strings.Split(
		`[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`,
		"\n",
	)

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayPart:  day.PartOne,
			Input:    exampleInput,
			Expected: 7,
		},
		{
			Name:     "part 1 real",
			DayPart:  day.PartOne,
			Input:    realInput,
			Expected: 514,
		},
		{
			Name:     "part 2 example",
			DayPart:  day.PartTwo,
			Input:    exampleInput,
			Expected: 33,
		},
		{
			Name:     "part 2 real",
			DayPart:  day.PartTwo,
			Input:    realInput,
			Expected: 21824,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := test.DayPart(test.Input, test.Extras...)
			if actual != test.Expected {
				t.Errorf("Expected %d, actual %d", test.Expected, actual)
			}
		})
	}
}
