package day05

import (
	"slices"

	"github.com/chwallen/advent-of-code/internal/util"
)

type idRange struct {
	start int
	end   int
}

func PartOne(lines []string, extras ...any) any {
	idRanges, idsStart := readRanges(lines)

	freshItems := 0
	for _, line := range lines[idsStart:] {
		id := util.Atoi(line)
		for _, r := range idRanges {
			if r.start <= id && id <= r.end {
				freshItems++
				break
			}
		}
	}
	return freshItems
}

func PartTwo(lines []string, extras ...any) any {
	idRanges, _ := readRanges(lines)

	deduplicatedRanges := make([]idRange, 0, len(idRanges))
	deduplicatedRanges = append(deduplicatedRanges, idRanges[0])

	for _, next := range idRanges[1:] {
		previous := &deduplicatedRanges[len(deduplicatedRanges)-1]
		if next.start > previous.end {
			deduplicatedRanges = append(deduplicatedRanges, next)
		} else {
			previous.end = max(previous.end, next.end)
		}
	}

	freshIds := 0
	for _, r := range deduplicatedRanges {
		freshIds += r.end - r.start + 1
	}
	return freshIds
}

func readRanges(lines []string) (idRanges []idRange, idsStart int) {
	var i int
	for i = 0; lines[i] != ""; i++ {
		start, end := util.CutToInts(lines[i], "-")
		idRanges = append(idRanges, idRange{start, end})
	}
	slices.SortFunc(idRanges, func(a, b idRange) int {
		return a.start - b.start
	})
	return idRanges, i + 1
}
