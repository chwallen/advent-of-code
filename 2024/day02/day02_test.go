package day02_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/chwallen/advent-of-code/2024/day02"
	"github.com/chwallen/advent-of-code/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	exampleInput := strings.Split(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
		"\n",
	)

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayPart:  day.PartOne,
			Input:    exampleInput,
			Expected: 2,
		},
		{
			Name:     "part 1 real",
			DayPart:  day.PartOne,
			Input:    realInput,
			Expected: 534,
		},
		{
			Name:     "part 2 example",
			DayPart:  day.PartTwo,
			Input:    exampleInput,
			Expected: 4,
		},
		{
			Name:     "part 2 real",
			DayPart:  day.PartTwo,
			Input:    realInput,
			Expected: 577,
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
