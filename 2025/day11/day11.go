package day11

import (
	"strings"
)

type node struct {
	name    string
	outputs []*node
}

type dacFftKey struct {
	nodeName string
	seenDac  bool
	seenFft  bool
}

func PartOne(lines []string, extras ...any) any {
	start := readNodesTree(lines, "you")
	return findPathsToOut(start)
}

func PartTwo(lines []string, extras ...any) any {
	start := readNodesTree(lines, "svr")

	visited := make(map[dacFftKey]int)
	return findPathsToOutViaDacAndFft(start, visited, false, false)
}

func findPathsToOut(n node) int {
	paths := 0
	for _, o := range n.outputs {
		if o.name == "out" {
			paths += 1
		} else {
			paths += findPathsToOut(*o)
		}
	}

	return paths
}

func findPathsToOutViaDacAndFft(
	n node,
	visited map[dacFftKey]int,
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
		key := dacFftKey{nodeName: o.name, seenDac: seenDac, seenFft: seenFft}
		if v, found := visited[key]; found {
			paths += v
		} else {
			v = findPathsToOutViaDacAndFft(*o, visited, seenDac, seenFft)
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
		name := line[:s]
		n := nodes[name]

		outputs := strings.Split(line[len(n.name)+2:], " ")
		n.outputs = make([]*node, len(outputs))
		for j, o := range outputs {
			n.outputs[j] = nodes[o]
		}
	}

	return *nodes[start]
}
