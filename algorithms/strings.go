package algorithm

import (
	"math/rand"
	"strconv"
	"strings"
)

// Fuzz is the penalty for a mismatched character.
// The value should be between 0 and 1.
var Fuzz = 0.5

func FuzzyStringSearch(s, t string) float64 {
	// Empty objective
	if t == "" {
		return 0
	}

	// Perfect match to objective
	if s == t {
		return 1
	}

	// Case-insensitive matching
	ls := strings.ToLower(s)
	lt := strings.ToLower(t)

	// Overall mismatch of strings
	fuzziness := 1.0

	// Cached lengths
	sl := len(s)
	tl := len(t)
	// Denotes the start of search space
	x := 0

	// Amount of confidence in the two characters matching
	score := 0.0
	// Amount of confidence during the process of matching
	total := 0.0

	for i := 0; i < tl; i++ {
		// Substring that denotes the remaining content to search
		sub := ls[x:]
		// Case-insensitive character matching
		y := strings.IndexByte(sub, lt[i])

		if y == -1 {
			fuzziness += Fuzz
		} else {
			if x == x+y {
				// First index character match and consecutive characters
				score = 0.7
			} else {
				score = 0.1

				if s[x+y-1] == ' ' {
					// Bonus for acronyms as a result of two character matches
					score += 0.8
				}
			}

			if s[x+y] == t[i] {
				// Bonus for same case
				score += 0.1
			}

			// Update total with the score of the character
			total += score
			// Update x to reflect start of search substring
			x += y + 1
		}
	}

	// Amount of confidence in the two strings matching
	// Reduce penalty of long strings
	confidence := 0.5 * (total/float64(sl) + total/float64(tl)) / fuzziness

	if (lt[0] == ls[0]) && (confidence < 0.85) {
		confidence += 0.15
	}

	return confidence
}

func StringRandom(n int, s string) string {
	l := len(s)

	// bits is equavlent to the number of bits needed to represent s
	bits := uint(len(strconv.FormatInt(int64(l), 2)))
	// We subtract 1 from bits to clamp it between 0 and l - 1
	mask := int64(1<<bits - 1)
	// The number of characters than can be generated from a single int64
	availble := 63 / bits
	// Buffer to temporarily store sequence
	buffer := make([]byte, n)

	// Reduce calls to random by generating an int64 to produce multiple characters
	for i, cache, remain := n-1, rand.Int63(), availble; i >= 0; {
		if remain == 0 {
			cache = rand.Int63()
			remain = availble
		}

		if j := int(cache & mask); j < l {
			buffer[i] = s[j]
			i--
		}

		// Shift bits to access remaining bits of the int64
		cache >>= bits
		remain--
	}

	return string(buffer)
}

func StringShuffle(s string) string {
	b := []byte(s)

	for i := len(b) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
