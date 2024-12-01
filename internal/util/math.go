package util

// Abs gets the absolute value of v.
func Abs(v int) int {
	return max(-v, v)
}
