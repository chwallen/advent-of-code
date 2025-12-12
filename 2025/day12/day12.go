package day12

import (
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

type treeRegion struct {
	area     int
	packages []int
}

func PartOne(lines []string, extras ...any) any {
	var regions []treeRegion
	shapeSizes := make([]int, 0, 6)

	shapeSize := 0
	for _, line := range lines {
		switch len(line) {
		case 0, 1, 2:
			if shapeSize > 0 {
				shapeSizes = append(shapeSizes, shapeSize)
				shapeSize = 0
			}
		case 3:
			shapeSize += strings.Count(line, "#")
		default:
			region := treeRegion{packages: make([]int, 0, 6)}

			size, count, _ := strings.Cut(line, ": ")
			width, height := util.CutToInts(size, "x")
			region.area = width * height

			for c := range strings.FieldsSeq(count) {
				region.packages = append(region.packages, util.Atoi(c))
			}
			regions = append(regions, region)
		}
	}

	sum := 0
	for _, region := range regions {
		size := 0

		for i, p := range region.packages {
			size += shapeSizes[i] * p
		}

		if region.area >= size {
			sum++
		}
	}
	return sum
}
