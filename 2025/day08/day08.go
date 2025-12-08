package day08

import (
	"slices"
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
	"github.com/chwallen/advent-of-code/internal/util"
)

type edge struct {
	from     geom.Point3D
	to       geom.Point3D
	distance int64
}

func PartOne(lines []string, extras ...any) any {
	connections := len(lines)
	if len(extras) > 0 {
		connections = extras[0].(int)
	}

	coordinates := getBoxCoordinates(lines)
	edges := getBoxEdges(coordinates)

	dsu := ds.NewDisjointSetUnion[geom.Point3D](len(coordinates))
	for i := range connections {
		dsu.Union(edges[i].from, edges[i].to)
	}

	var circuitSizes []int
	for _, coordinate := range coordinates {
		if dsu.IsRoot(coordinate) {
			circuitSizes = append(circuitSizes, dsu.GetSize(coordinate))
		}
	}

	slices.SortFunc(circuitSizes, func(a, b int) int {
		return b - a
	})

	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func PartTwo(lines []string, extras ...any) any {
	coordinates := getBoxCoordinates(lines)
	edges := getBoxEdges(coordinates)
	dsu := ds.NewDisjointSetUnion[geom.Point3D](len(coordinates))

	for _, edge := range edges {
		dsu.Union(edge.from, edge.to)
		if dsu.Count() == 1 {
			return edge.from.X * edge.to.X
		}
	}

	return 0
}

func getBoxCoordinates(lines []string) []geom.Point3D {
	coordinates := make([]geom.Point3D, len(lines))
	for i, line := range lines {
		firstComma := strings.Index(line, ",")
		x := util.Atoi(line[0:firstComma])
		y, z := util.CutToInts(line[firstComma+1:], ",")
		coordinates[i] = geom.Point3D{X: x, Y: y, Z: z}
	}
	return coordinates
}

func getBoxEdges(boxes []geom.Point3D) []edge {
	var edges []edge
	for i, outer := range boxes[:len(boxes)-1] {
		for _, inner := range boxes[i+1:] {
			dx := int64(outer.X - inner.X)
			dy := int64(outer.Y - inner.Y)
			dz := int64(outer.Z - inner.Z)
			distance := dx*dx + dy*dy + dz*dz
			edges = append(edges, edge{from: outer, to: inner, distance: distance})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		if a.distance-b.distance < 0 {
			return -1
		}
		return 1
	})

	return edges
}
