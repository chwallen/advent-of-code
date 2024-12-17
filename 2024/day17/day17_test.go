package day17_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/chwallen/advent-of-code/2024/day17"
	"github.com/chwallen/advent-of-code/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	partOneExampleInput := strings.Split(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
		"\n",
	)

	partTwoExampleInput := strings.Split(
		`Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`,
		"\n",
	)

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayPart:  day.PartOne,
			Input:    partOneExampleInput,
			Expected: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			Name:     "part 1 real",
			DayPart:  day.PartOne,
			Input:    realInput,
			Expected: "4,1,7,6,4,1,0,2,7",
		},
		{
			Name:     "part 2 example",
			DayPart:  day.PartTwo,
			Input:    partTwoExampleInput,
			Expected: 117440,
		},
		{
			Name:     "part 2 real",
			DayPart:  day.PartTwo,
			Input:    realInput,
			Expected: 164279024971453,
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
