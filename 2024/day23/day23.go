package day23

import (
	"slices"
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
)

func PartOne(lines []string, extras ...any) any {
	graph := buildComputerGraph(lines)

	result := 0
	uniqueConnections := make(map[string]bool)

	for nodeA, neighborsSet := range graph {
		neighbours := neighborsSet.Items()
		for i, nodeB := range neighbours[:len(neighbours)-1] {
			for _, nodeC := range neighbours[i+1:] {
				if (nodeA[0] == 't' || nodeB[0] == 't' || nodeC[0] == 't') && graph[nodeB].Contains(nodeC) {
					nodes := []string{nodeA, nodeB, nodeC}
					slices.Sort(nodes)
					key := strings.Join(nodes, ",")
					if !uniqueConnections[key] {
						uniqueConnections[key] = true
						result += 1
					}
				}
			}
		}
	}

	return result
}

func PartTwo(lines []string, extras ...any) any {
	graph := buildComputerGraph(lines)

	R := ds.NewSet[string]()
	P := ds.NewSet[string]()
	X := ds.NewSet[string]()
	for node := range graph {
		P.Add(node)
	}

	var largestClique []string
	findLargestClique(graph, R, P, X, &largestClique)
	slices.Sort(largestClique)

	return strings.Join(largestClique, ",")
}

// Bron-Kerbosch algorithm with pivot
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func findLargestClique(
	graph map[string]ds.Set[string],
	R, P, X ds.Set[string],
	largestClique *[]string,
) {
	if len(P) == 0 && len(X) == 0 {
		if len(R) > len(*largestClique) {
			*largestClique = R.Items()
		}
		return
	}

	pivotNode := findNodeWithLargestNeighborhood(graph, P.Union(X))
	// Process nodes that aren't in the pivot's neighborhood
	for node := range P.Difference(graph[pivotNode]).All() {
		newR := R.Clone()
		newR.Add(node)
		findLargestClique(
			graph,
			newR,
			P.Intersection(graph[node]),
			X.Intersection(graph[node]),
			largestClique,
		)

		delete(P, node)
		X.Add(node)
	}
}

func findNodeWithLargestNeighborhood(
	graph map[string]ds.Set[string],
	nodes ds.Set[string],
) string {
	var node string
	for n := range nodes.All() {
		if len(graph[n]) > len(graph[node]) {
			node = n
		}
	}
	return node
}

func buildComputerGraph(lines []string) map[string]ds.Set[string] {
	computerGraph := make(map[string]ds.Set[string], 200)
	for _, line := range lines {
		a, b, _ := strings.Cut(line, "-")

		addConnection(computerGraph, a, b)
		addConnection(computerGraph, b, a)
	}

	return computerGraph
}

func addConnection(graph map[string]ds.Set[string], nodeA, nodeB string) {
	if _, setExists := graph[nodeA]; !setExists {
		graph[nodeA] = ds.NewSet[string]()
	}
	graph[nodeA].Add(nodeB)
}
