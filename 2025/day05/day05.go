package day05

import (
	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/util"
)

type idRange struct {
	start int
	end   int
}

func PartOne(lines []string, extras ...any) any {
	rangeSet, idsStart := readRanges(lines)
	idRanges := rangeSet.Items()

	freshItems := 0
	for _, line := range lines[idsStart:] {
		id := util.Atoi(line)
		for _, r := range idRanges {
			if r.start <= id && id <= r.end {
				freshItems += 1
				break
			}
		}
	}
	return freshItems
}

func PartTwo(lines []string, extras ...any) any {
	currentRanges, _ := readRanges(lines)
	nextRanges := ds.NewSet[idRange]()
	idRanges := make([]idRange, 0, len(currentRanges))

	for len(idRanges) != len(currentRanges) {
		idRanges = idRanges[:0]
		for r := range currentRanges.All() {
			idRanges = append(idRanges, r)
			nextRanges.Add(r)
		}

		for i, outer := range idRanges[:len(idRanges)-1] {
			for _, inner := range idRanges[i+1:] {
				if (outer.start <= inner.start && outer.end >= inner.start) ||
					(inner.start <= outer.start && inner.end >= outer.start) {
					delete(nextRanges, outer)
					delete(nextRanges, inner)
					nextRanges.Add(idRange{min(outer.start, inner.start), max(outer.end, inner.end)})
				}
			}
		}

		clear(currentRanges)
		currentRanges, nextRanges = nextRanges, currentRanges
	}

	freshIds := 0
	for _, r := range idRanges {
		freshIds += r.end - r.start + 1
	}
	return freshIds
}

func readRanges(lines []string) (idRanges ds.Set[idRange], idsStart int) {
	idRanges = ds.NewSet[idRange]()
	var i int
	for i = 0; lines[i] != ""; i++ {
		start, end := util.CutToInts(lines[i], "-")
		idRanges.Add(idRange{start, end})
	}
	return idRanges, i + 1
}
