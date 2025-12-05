package day05_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/chwallen/advent-of-code/2025/day05"
	"github.com/chwallen/advent-of-code/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	exampleInput := strings.Split(
		`3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
		"\n",
	)

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayPart:  day.PartOne,
			Input:    exampleInput,
			Expected: 3,
		},
		{
			Name:     "part 1 real",
			DayPart:  day.PartOne,
			Input:    realInput,
			Expected: 739,
		},
		{
			Name:     "part 2 example",
			DayPart:  day.PartTwo,
			Input:    exampleInput,
			Expected: 14,
		},
		{
			Name:     "part 2 real",
			DayPart:  day.PartTwo,
			Input:    realInput,
			Expected: 344486348901788,
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
