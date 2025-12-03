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

// IntPow raises base to the power of exponent.
func IntPow(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	result := base
	for i := 1; i < exponent; i++ {
		result *= base
	}
	return result
}

// CountDigits counts the number of digits n consists of.
func CountDigits(n int) int {
	x, count := 10, 1
	for x <= n {
		x *= 10
		count++
	}
	return count
}
