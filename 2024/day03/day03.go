package day03

import (
	"regexp"
	"strings"

	"github.com/chwallen/advent-of-code/internal/util"
)

var multiplicationRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func PartOne(lines []string, extras ...any) any {
	sum := 0
	for _, line := range lines {
		matches := multiplicationRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			sum += util.Atoi(match[1]) * util.Atoi(match[2])
		}
	}
	return sum
}

func PartTwo(lines []string, extras ...any) any {
	line := strings.Join(lines, "")
	sum := 0
	for _, s := range strings.Split(line, "do()") {
		disabledRegionStart := strings.Index(s, "don't()")
		if disabledRegionStart == -1 {
			disabledRegionStart = len(s)
		}
		enabledRegion := s[0:disabledRegionStart]
		for _, match := range multiplicationRegex.FindAllStringSubmatch(enabledRegion, -1) {
			sum += util.Atoi(match[1]) * util.Atoi(match[2])
		}
	}
	return sum
}
