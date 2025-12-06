package day06

import (
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

type problem struct {
	answer int
	op     string
}

func PartOne(lines []string, extras ...any) any {
	ops := strings.Fields(lines[len(lines)-1])
	problems := make([]problem, len(ops))
	for i, op := range ops {
		problems[i].op = op
		if op == "*" {
			problems[i].answer = 1
		} else {
			problems[i].answer = 0
		}
	}

	for _, line := range lines[:len(lines)-1] {
		items := strings.Fields(line)
		for i, item := range items {
			v := util.Atoi(item)
			if problems[i].op == "*" {
				problems[i].answer *= v
			} else {
				problems[i].answer += v
			}
		}
	}

	sum := 0
	for _, p := range problems {
		sum += p.answer
	}
	return sum
}

func PartTwo(lines []string, extras ...any) any {
	var (
		valueBuilder strings.Builder
		values       []int
		answers      []int
	)
	opsLine := lines[len(lines)-1]
	for x := len(lines[0]) - 1; x >= 0; x-- {
		for y := range len(lines) - 1 {
			char := lines[y][x]
			if char != ' ' {
				valueBuilder.WriteByte(char)
			}
		}
		if valueBuilder.Len() > 0 {
			values = append(values, util.Atoi(valueBuilder.String()))
			valueBuilder.Reset()
		}
		if op := opsLine[x]; op != ' ' {
			var answer int
			if op == '*' {
				answer = 1
			}

			for _, v := range values {
				if op == '*' {
					answer *= v
				} else {
					answer += v
				}
			}
			values = values[:0]
			answers = append(answers, answer)

			// Skip next blank column
			x--
		}
	}

	sum := 0
	for _, p := range answers {
		sum += p
	}
	return sum
}
