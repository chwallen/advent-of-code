package day19

import (
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
)

type node struct {
	white   *node
	blue    *node
	black   *node
	red     *node
	green   *node
	pattern string
}

func PartOne(lines []string, extras ...any) any {
	rootNode := createTree(lines[0])

	knownDesigns := ds.NewSet[string]()
	for pattern := range strings.SplitSeq(lines[0], ", ") {
		knownDesigns.Add(pattern)
	}

	possibleDesigns := 0
	for _, line := range lines[2:] {
		if isDesignPossible(rootNode, line, knownDesigns) {
			possibleDesigns += 1
		}
	}
	return possibleDesigns
}

func isDesignPossible(root *node, design string, knownDesigns ds.Set[string]) bool {
	if knownDesigns.Contains(design) {
		return true
	}

	currentNode := root
	for _, char := range design {
		currentNode = currentNode.getChild(char)
		if currentNode == nil {
			return false
		}

		patternLength := len(currentNode.getPattern())
		if patternLength > 0 && isDesignPossible(root, design[patternLength:], knownDesigns) {
			knownDesigns.Add(design)
			return true
		}
	}
	return false
}

func PartTwo(lines []string, extras ...any) any {
	rootNode := createTree(lines[0])

	knownDesigns := make(map[string]int, 20_000)
	for pattern := range strings.SplitSeq(lines[0], ", ") {
		countPossibleDesigns(rootNode, pattern, knownDesigns)
	}

	possibleDesignCombinations := 0
	for _, line := range lines[2:] {
		possibleDesignCombinations += countPossibleDesigns(rootNode, line, knownDesigns)
	}
	return possibleDesignCombinations
}

func countPossibleDesigns(root *node, design string, knownDesigns map[string]int) int {
	count, isKnownDesign := knownDesigns[design]
	if isKnownDesign {
		return count
	}
	currentNode := root
	for _, char := range design {
		currentNode = currentNode.getChild(char)
		if currentNode == nil {
			knownDesigns[design] = count
			return count
		} else if design == currentNode.getPattern() {
			count += 1
			knownDesigns[design] = count
			return count
		}
		patternLength := len(currentNode.getPattern())
		if patternLength > 0 {
			possibleDesigns := countPossibleDesigns(root, design[patternLength:], knownDesigns)
			knownDesigns[design] = possibleDesigns
			count += possibleDesigns
		}
	}
	knownDesigns[design] = count
	return count
}

func createTree(input string) *node {
	root := &node{}
	for pattern := range strings.SplitSeq(input, ", ") {
		currentNode := root
		for i, char := range pattern {
			currentNode = currentNode.getOrCreateChild(pattern, char, i)
		}
	}
	return root
}

func (n *node) getChild(char rune) *node {
	return *n.getChildField(char)
}

func (n *node) getPattern() string {
	return n.pattern
}

func (n *node) getChildField(char rune) **node {
	switch char {
	case 'w':
		return &n.white
	case 'u':
		return &n.blue
	case 'b':
		return &n.black
	case 'r':
		return &n.red
	case 'g':
		return &n.green
	default:
		panic("invalid char " + string(char))
	}
}

func (n *node) getOrCreateChild(pattern string, char rune, index int) *node {
	childField := n.getChildField(char)
	child := *childField
	if child == nil {
		child = &node{}
		*childField = child
	}
	if len(pattern)-1 == index {
		child.pattern = pattern
	}
	return child
}
