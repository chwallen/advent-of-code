package day11

import (
	"strings"
)

type node struct {
	name    string
	outputs []*node
}

type visitedKey struct {
	nodeName string
	seenDac  bool
	seenFft  bool
}

func PartOne(lines []string, extras ...any) any {
	start := readNodesTree(lines, "you")
	visited := make(map[visitedKey]int)
	return findPathsToOut(start, visited, true, true)
}

func PartTwo(lines []string, extras ...any) any {
	start := readNodesTree(lines, "svr")
	visited := make(map[visitedKey]int)
	return findPathsToOut(start, visited, false, false)
}

func findPathsToOut(
	n node,
	visited map[visitedKey]int,
	hasSeenDac bool,
	hasSeenFft bool,
) int {
	paths := 0
	for _, o := range n.outputs {
		if o.name == "out" {
			if hasSeenDac && hasSeenFft {
				paths += 1
			}
			continue
		}

		seenDac := hasSeenDac || o.name == "dac"
		seenFft := hasSeenFft || o.name == "fft"
		key := visitedKey{nodeName: o.name, seenDac: seenDac, seenFft: seenFft}
		if v, found := visited[key]; found {
			paths += v
		} else {
			v = findPathsToOut(*o, visited, seenDac, seenFft)
			visited[key] = v
			paths += v
		}
	}

	return paths
}

func readNodesTree(lines []string, start string) node {
	nodes := make(map[string]*node)
	for _, line := range lines {
		s := strings.Index(line, ":")
		n := node{name: line[:s]}
		nodes[n.name] = &n
	}
	nodes["out"] = &node{name: "out"}

	for _, line := range lines {
		s := strings.Index(line, ":")
		n := nodes[line[:s]]

		for o := range strings.FieldsSeq(line[s+2:]) {
			n.outputs = append(n.outputs, nodes[o])
		}
	}

	return *nodes[start]
}
