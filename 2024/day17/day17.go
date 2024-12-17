package day17

import (
	"slices"
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

type instruction struct {
	input1 *int
	input2 *int
	output *int
	op     func(*int, *int, *int)
}

type program struct {
	instructions []instruction
	a            int
	b            int
	c            int
	pc           int
	outputIndex  int
	output       []int
}

func PartOne(lines []string, extras ...any) any {
	output := parseIntoProgram(lines, false).execute()

	var sb strings.Builder
	sb.WriteByte(byte(output[0] + '0'))
	for _, v := range output[1:] {
		sb.WriteByte(',')
		sb.WriteByte(byte(v + '0'))
	}
	return sb.String()
}

func PartTwo(lines []string, extras ...any) any {
	p := parseIntoProgram(lines, true)

	tokens := strings.TrimPrefix(lines[4], "Program: ")
	expectedOutput := util.SplitToInts(tokens, ",", make([]int, 0, 20))

	// A must be at least 8^(len(expectedOutput)-1) to produce len(expectedOutput) symbols
	initialA := util.IntPow(8, len(expectedOutput)-1)
	for {
		p.a = initialA
		actualOutput := p.execute()
		if slices.Equal(expectedOutput, actualOutput) {
			return initialA
		}

		for i := range slices.Backward(expectedOutput) {
			if expectedOutput[i] != actualOutput[i] {
				// The token at position i changes every 8^i increment of A
				initialA += util.IntPow(8, i)
				p.outputIndex = 0
				p.b = 0
				p.c = 0
				break
			}
		}
	}
}

func parseIntoProgram(lines []string, recyleOutput bool) *program {
	p := &program{}

	p.a = util.Atoi(strings.TrimPrefix(lines[0], "Register A: "))
	p.b = util.Atoi(strings.TrimPrefix(lines[1], "Register B: "))
	p.c = util.Atoi(strings.TrimPrefix(lines[2], "Register C: "))

	tokens := strings.Split(strings.TrimPrefix(lines[4], "Program: "), ",")
	p.instructions = make([]instruction, 0, len(tokens)/2)
	p.output = make([]int, 0, len(tokens))
	p.outputIndex = 0

	for i := 0; i < len(tokens); i += 2 {
		opCode := tokens[i]
		operand := util.Atoi(tokens[i+1])

		var op func(*int, *int, *int)
		var input1, input2, output *int

		switch opCode {
		case "0":
			op = divideAndStore
			input1 = &p.a
			input2 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			output = &p.a
		case "1":
			op = xorAndStore
			input1 = &p.b
			input2 = &operand
			output = &p.b
		case "2":
			op = mod8AndStore
			input1 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			input2 = nil
			output = &p.b
		case "3":
			op = func(a, jumpTarget, pc *int) {
				if *a != 0 {
					// Compensate for pc being half of expected and loop adding 1
					*pc = (*jumpTarget / 2) - 1
				}
			}
			input1 = &p.a
			input2 = &operand
			output = &p.pc
		case "4":
			op = xorAndStore
			input1 = &p.b
			input2 = &p.c
			output = &p.b
		case "5":
			op = func(input, outIndex, _ *int) {
				value := *input % 8
				if len(p.output) == len(tokens) && recyleOutput {
					p.output[*outIndex] = value
					*outIndex = (*outIndex + 1) % len(tokens)
				} else {
					p.output = append(p.output, value)
				}
			}
			input1 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			input2 = &p.outputIndex
			output = nil
		case "6":
			op = divideAndStore
			input1 = &p.a
			input2 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			output = &p.b
		case "7":
			op = divideAndStore
			input1 = &p.a
			input2 = getComboOperandValue(&operand, &p.a, &p.b, &p.c)
			output = &p.c
		default:
			panic("Invalid operation code " + opCode)
		}
		p.instructions = append(p.instructions, instruction{input1, input2, output, op})
	}

	return p
}

func (p *program) execute() []int {
	for p.pc = 0; p.pc < len(p.instructions); p.pc += 1 {
		inst := p.instructions[p.pc]
		inst.op(inst.input1, inst.input2, inst.output)
	}
	return p.output
}

func divideAndStore(numerator, factor, output *int) {
	*output = *numerator / util.IntPow(2, *factor)
}

func mod8AndStore(numerator, _, output *int) {
	*output = *numerator % 8
}

func xorAndStore(left, right, output *int) {
	*output = *left ^ *right
}

func getComboOperandValue(operand, a, b, c *int) *int {
	switch *operand {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		return operand
	}
}
