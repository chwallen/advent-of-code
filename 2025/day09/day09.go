package day09

import (
	"github.com/chwallen/advent-of-code/internal/geom"
	"github.com/chwallen/advent-of-code/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	tiles := readTiles(lines)
	var maxArea int64
	for i, cornerA := range tiles[:len(tiles)-1] {
		for _, cornerB := range tiles[i+1:] {
			area := calculateArea(cornerA, cornerB)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func PartTwo(lines []string, extras ...any) any {
	tiles := readTiles(lines)

	var maxArea int64
	for i, cornerA := range tiles[:len(lines)-1] {
	Inner:
		for _, cornerB := range tiles[i+1:] {
			for k, tileX := range tiles {
				var tileY geom.Point
				if k < len(tiles)-1 {
					tileY = tiles[k+1]
				} else {
					tileY = tiles[0]
				}
				if hasHole(cornerA, cornerB, tileX, tileY) {
					continue Inner
				}
			}

			area := calculateArea(cornerA, cornerB)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func readTiles(lines []string) (tiles []geom.Point) {
	tiles = make([]geom.Point, len(lines))
	for i, line := range lines {
		x, y := util.CutToInts(line, ",")
		tiles[i] = geom.Point{X: x, Y: y}
	}
	return tiles
}

func calculateArea(a, b geom.Point) int64 {
	dx := util.Abs(a.X-b.X) + 1
	dy := util.Abs(a.Y-b.Y) + 1
	return int64(dx) * int64(dy)
}

func hasHole(cornerA, cornerB, tileX, tileY geom.Point) bool {
	return max(tileX.X, tileY.X) > min(cornerB.X, cornerA.X) &&
		max(cornerB.X, cornerA.X) > min(tileX.X, tileY.X) &&
		max(tileX.Y, tileY.Y) > min(cornerB.Y, cornerA.Y) &&
		max(cornerB.Y, cornerA.Y) > min(tileX.Y, tileY.Y)
}
