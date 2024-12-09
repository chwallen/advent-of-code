package day09_test

import (
	_ "embed"
	"strings"
	"testing"

	day "github.com/chwallen/advent-of-code/2024/day09"
	"github.com/chwallen/advent-of-code/internal/test"
)

//go:embed input.txt
var input string

func TestDay(t *testing.T) {
	exampleInput := []string{"2333133121414131402"}
	realInput := strings.Split(input[0:len(input)-1], "\n")

	tests := []test.Test{
		{
			Name:     "part 1 example",
			DayFunc:  day.PartOne,
			Input:    exampleInput,
			Expected: 1928,
		},
		{
			Name:     "part 1 real",
			DayFunc:  day.PartOne,
			Input:    realInput,
			Expected: 6367087064415,
		},
		{
			Name:     "part 2 example",
			DayFunc:  day.PartTwo,
			Input:    exampleInput,
			Expected: 2858,
		},
		{
			Name:     "part 2 real",
			DayFunc:  day.PartTwo,
			Input:    realInput,
			Expected: 6390781891880,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := test.DayFunc(test.Input, test.Extras...)
			if actual != test.Expected {
				t.Errorf("Expected %d, actual %d", test.Expected, actual)
			}
		})
	}
}
