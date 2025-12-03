package day02

import (
	"iter"
	"math"
	"strings"

	"github.com/chwallen/advent-of-code/internal/ds"
	"github.com/chwallen/advent-of-code/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	return sumInvalidIds(lines, true)
}

func PartTwo(lines []string, extras ...any) any {
	return sumInvalidIds(lines, false)
}

func sumInvalidIds(
	lines []string,
	evenOnly bool,
) int {
	sum := 0
	seen := ds.NewSet[int]()

	var getFactors func(int) iter.Seq[int]
	if evenOnly {
		getFactors = func(n int) iter.Seq[int] {
			return func(yield func(int) bool) {
				yield(n / 2)
			}
		}
	} else {
		getFactors = func(n int) iter.Seq[int] {
			return func(yield func(int) bool) {
				for i := 1; i*2 <= n; i++ {
					if n%i == 0 && !yield(i) {
						return
					}
				}
			}
		}
	}

	for r := range strings.SplitSeq(lines[0], ",") {
		start, end := util.CutToInts(r, "-")
		startDigits := util.CountDigits(start)
		endDigits := util.CountDigits(end)

		if evenOnly && startDigits%2 == 1 {
			divisor := util.IntPow(10, startDigits)
			start = int(math.Ceil(float64(start)/float64(divisor))) * divisor
			startDigits = util.CountDigits(start)
		}

		if startDigits == endDigits {
			for factor := range getFactors(startDigits) {
				sum += sumInvalidIDsInRange(
					start,
					end,
					startDigits,
					factor,
					seen,
				)
			}
		} else {
			for factor := range getFactors(startDigits) {
				sum += sumInvalidIDsInRange(
					start,
					min(util.IntPow(10, startDigits)-1, end),
					startDigits,
					factor,
					seen,
				)
			}

			for digits := startDigits + 1; digits <= endDigits; digits++ {
				if evenOnly && digits%2 == 1 {
					continue
				}
				lowerBound := util.IntPow(10, digits-1)
				upperBound := min(util.IntPow(10, digits)-1, end)
				for factor := range getFactors(digits) {
					sum += sumInvalidIDsInRange(
						lowerBound,
						upperBound,
						digits,
						factor,
						seen,
					)
				}
			}
		}

		clear(seen)
	}

	return sum
}

func sumInvalidIDsInRange(
	start int,
	end int,
	digits int,
	factor int,
	seen ds.Set[int],
) int {
	chunks := digits / factor
	multiplier := 1
	power := 1
	for range chunks - 1 {
		power *= util.IntPow(10, factor)
		multiplier += power
	}

	blockMin := (start + multiplier - 1) / multiplier
	blockMax := end / multiplier

	sum := 0
	for block := blockMin; block <= blockMax; block++ {
		id := block * multiplier
		if seen.Add(id) {
			sum += id
		}
	}
	return sum
}
