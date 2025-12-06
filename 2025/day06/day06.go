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
	var problems []problem
	for op := range strings.FieldsSeq(lines[len(lines)-1]) {
		p := problem{answer: 0, op: op}
		if op == "*" {
			p.answer = 1
		}
		problems = append(problems, p)
	}

	for _, line := range lines[:len(lines)-1] {
		i := 0
		for item := range strings.FieldsSeq(line) {
			v := util.Atoi(item)
			if problems[i].op == "*" {
				problems[i].answer *= v
			} else {
				problems[i].answer += v
			}
			i++
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
	)
	opsLine := lines[len(lines)-1]
	sum := 0
	for x := len(lines[0]) - 1; x >= 0; x-- {
		for y := range len(lines) - 1 {
			char := lines[y][x]
			if char != ' ' {
				valueBuilder.WriteByte(char)
			}
		}

		values = append(values, util.Atoi(valueBuilder.String()))
		valueBuilder.Reset()

		switch op := opsLine[x]; op {
		case '*':
			answer := 1
			for _, v := range values {
				answer *= v
			}
			sum += answer
		case '+':
			for _, v := range values {
				sum += v
			}
		default:
			continue
		}

		values = values[:0]
		// Skip next blank column
		x--
	}

	return sum
}
