package day12

import (
	"iter"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/geom"
	"github.com/chwallen/advent-of-code/internal/util"
)

type groupType int

const (
	unvisitedSameCrop groupType = iota
	visitedSameCrop
	otherCrop
)

type cropRegion struct {
	area      int
	perimeter int
	sides     int
}

type gardenTile struct {
	point   geom.Point
	crop    rune
	visited bool
}

var directions = []geom.Direction{geom.Up, geom.Right, geom.Down, geom.Left, geom.Up}

func PartOne(lines []string, extras ...any) any {
	price := 0
	for region := range getCropRegions(lines) {
		price += region.area * region.perimeter
	}
	return price
}

func PartTwo(lines []string, extras ...any) any {
	price := 0
	for region := range getCropRegions(lines) {
		price += region.area * region.sides
	}
	return price
}

func getCropRegions(lines []string) iter.Seq[cropRegion] {
	width, height := len(lines[0]), len(lines)
	grid := ds.NewGrid[gardenTile](width, height)

	for y := range height {
		for x := range width {
			tile := grid.Get(x, y)
			tile.point = geom.Point{X: x, Y: y}
			tile.crop = rune(lines[y][x])
		}
	}

	return func(yield func(region cropRegion) bool) {
		for y := range height {
			for x := range width {
				tile := grid.Get(x, y)
				if tile.visited {
					continue
				}

				region := getCropRegion(grid, tile)
				if !yield(region) {
					return
				}
			}
		}
	}
}

func getCropRegion(
	grid *ds.Grid[gardenTile],
	origin *gardenTile,
) cropRegion {
	queue := ds.Queue[*gardenTile]{}
	region := cropRegion{}
	crop := origin.crop

	groupNeighbors := func(neighbor geom.Point) groupType {
		t := getTileOrDefault(grid, neighbor)
		switch {
		case t.crop != crop:
			return otherCrop
		case t.visited:
			return visitedSameCrop
		default:
			return unvisitedSameCrop
		}
	}

	var tile *gardenTile
	queue = queue.Push(origin)
	for !queue.IsEmpty() {
		tile, queue = queue.Pop()
		// Unvisited tiles may become visited while in the queue
		if tile.visited {
			continue
		}
		tile.visited = true

		groups := util.GroupBy(tile.point.Neighbors(), groupNeighbors)
		region.area += 1
		region.perimeter += len(groups[otherCrop])
		region.sides += countCorners(grid, tile.point, crop)

		for _, p := range groups[unvisitedSameCrop] {
			queue = queue.Push(grid.Get(p.XY()))
		}
	}

	return region
}

func countCorners(
	grid *ds.Grid[gardenTile],
	origin geom.Point,
	crop rune,
) int {
	corners := 0
	for i := 0; i < len(directions)-1; i++ {
		dir1 := directions[i]
		dir2 := directions[i+1]

		side1Crop := getTileOrDefault(grid, origin.Add(dir1)).crop
		side2Crop := getTileOrDefault(grid, origin.Add(dir2)).crop
		diagonalCrop := getTileOrDefault(grid, origin.Add(dir1).Add(dir2)).crop

		if (crop != side1Crop && crop != side2Crop) ||
			(crop == side1Crop && crop == side2Crop && crop != diagonalCrop) {
			corners += 1
		}
	}
	return corners
}

func getTileOrDefault(grid *ds.Grid[gardenTile], point geom.Point) gardenTile {
	if grid.IsWithinBounds(point.XY()) {
		return *grid.Get(point.XY())
	}
	return gardenTile{}
}
