package day04_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/chwallen/advent-of-code/2024/day04"
	"github.com/chwallen/advent-of-code/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	exampleInput := strings.Split(
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
		"\n",
	)

	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayPart:  day.PartOne,
			Input:    exampleInput,
			Expected: 18,
		},
		{
			Name:     "part 1 real",
			DayPart:  day.PartOne,
			Input:    realInput,
			Expected: 2390,
		},
		{
			Name:     "part 2 example",
			DayPart:  day.PartTwo,
			Input:    exampleInput,
			Expected: 9,
		},
		{
			Name:     "part 2 real",
			DayPart:  day.PartTwo,
			Input:    realInput,
			Expected: 1809,
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
