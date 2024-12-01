package util

import (
	"fmt"
	"strconv"
	"strings"
)

// Atoi is equivalent to strconv.Atoi but panics on error.
func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err == nil {
		return v
	}
	panic(fmt.Errorf("cannot convert string \"%s\" to int", s))
}

// CutToInts cuts s at sep using strings.Cut and converts the parts to ints.
// Panics if the parts cannot be converted.
func CutToInts(s, sep string) (l, r int) {
	a, b, _ := strings.Cut(s, sep)
	return Atoi(a), Atoi(b)
}
