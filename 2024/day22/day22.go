package day22

import (
	"slices"

	"github.com/chwallen/advent-of-code/internal/util"
)

const (
	// Keep the 24 least significant bits
	bitsToKeep = 0xffffff
	steps      = 2000
)

func PartOne(lines []string, extras ...any) any {
	result := 0
	for _, line := range lines {
		secret := util.Atoi(line)
		for range steps {
			secret = calculateNextSecret(secret)
		}
		result += secret
	}

	return result
}

func PartTwo(lines []string, extras ...any) any {
	// There are 19^4 = 130,321 possible sequences as each price can only be
	// in the range [-9, 9] due to mod 10. We can therefore pre-allocate a
	// slot for each, making access much faster than hashing.
	const maxSequences = 19 * 19 * 19 * 19
	// Use uint16 for reduced memory footprint and improved cache efficiency.
	bananas := make([]uint16, maxSequences)
	seenSequences := make([]bool, maxSequences)

	for _, line := range lines {
		secret := util.Atoi(line)

		previousPrice := uint16(secret % 10)
		priceSequence := 0

		for i := range steps {
			secret = calculateNextSecret(secret)

			price := uint16(secret % 10)
			// Convert [-9, 9] to [0, 18]
			priceChange := int(price - previousPrice + 9)
			previousPrice = price

			// Shift by 19, add this loop's change, and keep just the four latest changes.
			priceSequence = (priceSequence*19 + priceChange) % maxSequences

			if i >= 3 && !seenSequences[priceSequence] {
				seenSequences[priceSequence] = true
				bananas[priceSequence] += price
			}
		}

		clear(seenSequences)
	}

	return int(slices.Max(bananas))
}

func calculateNextSecret(secret int) int {
	// multiply by 64
	secret ^= (secret << 6) & bitsToKeep
	// divide by 32
	secret ^= (secret >> 5) & bitsToKeep
	// multiply by 2048
	secret ^= (secret << 11) & bitsToKeep
	return secret
}
