package day24

import (
	"fmt"
	"slices"
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/util"
)

type gate struct {
	input1 string
	input2 string
	op     string
	output string
}

func PartOne(lines []string, extras ...any) any {
	_, wires := parseGatesAndWires(lines)

	zWires := make([]string, 0, 50)
	for name := range wires {
		if name[0] == 'z' {
			zWires = append(zWires, name)
		}
	}

	slices.Sort(zWires)
	slices.Reverse(zWires)

	value := 0
	for _, zWire := range zWires {
		bit := wires[zWire]
		value = (value << 1) | bit
	}
	return value
}

func PartTwo(lines []string, extras ...any) any {
	connections, wires := parseGatesAndWires(lines)

	lastZ := "z00"
	for s := range wires {
		if s[0] == 'z' {
			n := util.Atoi(s[1:])
			if n > util.Atoi(lastZ[1:]) {
				lastZ = s
			}
		}
	}

	invalidWires := ds.NewSet[string]()
	for _, gates := range connections {
		for _, g := range gates {
			if g.isZOutput() && g.op != "XOR" && g.output != lastZ {
				invalidWires.Add(g.output)
			} else if g.op == "XOR" && !g.isZOutput() && !g.isXAndYInputs() {
				invalidWires.Add(g.output)
			} else if g.op == "AND" && g.input1 != "x00" && g.input2 != "x00" {
				for _, subConnection := range connections[g.output] {
					if subConnection.op != "OR" {
						invalidWires.Add(g.output)
					}
				}
			} else if g.op == "XOR" {
				for _, subConnection := range connections[g.output] {
					if subConnection.op == "OR" {
						invalidWires.Add(g.output)
					}
				}
			}
		}
	}

	return strings.Join(slices.Sorted(invalidWires.All()), ",")
}

func (g gate) isXAndYInputs() bool {
	return (g.input1[0] == 'x' && g.input2[0] == 'y') ||
		(g.input1[0] == 'y' && g.input2[0] == 'x')
}

func (g gate) isZOutput() bool {
	return g.output[0] == 'z'
}

func parseGatesAndWires(lines []string) (connections map[string][]gate, wires map[string]int) {
	connections = make(map[string][]gate)
	wires = make(map[string]int)

	unprocessedGates := make([]gate, 0, 1000)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line[3] == ':' {
			name, value, _ := strings.Cut(line, ": ")
			wires[name] = util.Atoi(value)
		} else {
			parts := strings.Fields(line)
			g := gate{input1: parts[0], input2: parts[2], op: parts[1], output: parts[4]}
			connections[g.input1] = append(connections[g.input1], g)
			connections[g.input2] = append(connections[g.input2], g)

			if !processGate(g, wires) {
				unprocessedGates = append(unprocessedGates, g)
			}
		}
	}

	var g gate
	for len(unprocessedGates) > 0 {
		g, unprocessedGates = unprocessedGates[0], unprocessedGates[1:]
		if !processGate(g, wires) {
			unprocessedGates = append(unprocessedGates, g)
		}
	}

	return connections, wires
}

func processGate(g gate, wires map[string]int) bool {
	left, isLeftOk := wires[g.input1]
	right, isRightOk := wires[g.input2]
	if isLeftOk && isRightOk {
		switch g.op {
		case "AND":
			wires[g.output] = left & right
		case "OR":
			wires[g.output] = left | right
		case "XOR":
			wires[g.output] = left ^ right
		default:
			panic(fmt.Errorf("unknown op %s", g.op))
		}
		return true
	} else {
		return false
	}
}
