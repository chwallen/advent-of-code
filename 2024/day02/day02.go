package day02

import (
	"slices"

	"github.com/chwallen/advent-of-code/internal/util"
)

func PartOne(lines []string, extras ...any) any {
	return getNumberOfSafeReports(lines, isReportSafe)
}

func PartTwo(lines []string, extras ...any) any {
	return getNumberOfSafeReports(lines, isReportSafeWhenModified)
}

func isReportSafeWhenModified(report []int) bool {
	for toSkip, item := range report {
		report = append(report[:toSkip], report[toSkip+1:]...)
		if isReportSafe(report) {
			return true
		}
		report = slices.Insert(report, toSkip, item)
	}
	return false
}

func getNumberOfSafeReports(lines []string, isSafe func(report []int) bool) int {
	safeReports := 0
	report := make([]int, 0, 10)
	for _, line := range lines {
		report = util.SplitToInts(line, " ", report)
		if isSafe(report) {
			safeReports += 1
		}
		report = report[:0]
	}
	return safeReports
}

func isReportSafe(report []int) bool {
	lastDiff := 0
	for i, item := range report[1:] {
		// i starts at 0 but range starts at 1
		prev := report[i]
		diff := item - prev

		// Sign change
		if diff*lastDiff < 0 {
			return false
		}
		absDiff := util.Abs(item - prev)
		if absDiff == 0 || absDiff > 3 {
			return false
		}
		lastDiff = diff
	}
	return true
}
