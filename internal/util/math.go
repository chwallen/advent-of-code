package util

// Abs gets the absolute value of v.
func Abs(v int) int {
	return max(-v, v)
}

// DivRem divides a with b and returns both the quotient and the remainder.
func DivRem(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a - b*quotient
	return quotient, remainder
}
