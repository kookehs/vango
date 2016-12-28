package algorithm

import (
	"math/rand"
	"testing"
)

func TestFuzzyStringSearch(t *testing.T) {
	matchingTests := []struct {
		s1, s2   string
		expected float64
	}{
		{
			s1:       "test",
			s2:       "",
			expected: 0,
		},
		{
			s1:       "hello",
			s2:       "pass",
			expected: 0,
		},
		{
			s1:       "hello",
			s2:       "hello",
			expected: 1,
		},
	}

	for _, v := range matchingTests {
		e := v.expected
		f := FuzzyStringSearch(v.s1, v.s2)

		if f != e {
			t.Errorf("Found %v expect %v\n", f, e)
		}
	}

	relativeTests := []struct {
		b        string
		s1, s2   string
		expected byte
	}{
		{
			b:        "Hello World!",
			s1:       "hello",
			s2:       "Hello",
			expected: '<',
		},
		{
			b:        "Hello World!",
			s1:       "H",
			s2:       "He",
			expected: '<',
		},
		{
			b:        "Hello World!",
			s1:       "HW",
			s2:       "HD",
			expected: '>',
		},
		{
			b:        "Hello World!",
			s1:       "Hel",
			s2:       "Hrld",
			expected: '>',
		},
		{
			b:        "Hello World!",
			s1:       "h",
			s2:       "w",
			expected: '>',
		},
	}

	for _, v := range relativeTests {
		e := v.expected

		f1 := FuzzyStringSearch(v.b, v.s1)
		f2 := FuzzyStringSearch(v.b, v.s2)

		if e == '=' {
			if f1 != f2 {
				t.Errorf("Found %v expected %v", f1, f2)
			}
		} else if e == '<' {
			if f1 > f2 {
				t.Errorf("Found %v expected %v", f1, f2)
			}
		} else if e == '>' {
			if f1 < f2 {
				t.Errorf("Found %v expected %v", f1, f2)
			}
		} else {
			t.Errorf("Unexpected value: %v", e)
		}
	}
}

// Repeatability is tested.
// Statistical distribution needs to be tested.
func TestStringRandom(t *testing.T) {
	tests := []struct {
		expected string
		seed     int64
		set      string
		size     int
	}{
		{
			expected: "bcd",
			seed:     42,
			set:      "abcdef",
			size:     3,
		},
		{
			expected: "deddcdbbdacdabcd",
			seed:     42,
			set:      "abcdef",
			size:     16,
		},
	}

	for _, v := range tests {
		rand.Seed(v.seed)
		a := StringRandom(v.size, v.set)
		e := v.expected

		if a != e {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}

// Repeatability is tested.
// Statistical distribution needs to be tested.
func TestStringShuffle(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
		seed     int64
	}{
		{
			actual:   "shuffle",
			expected: "sfhufel",
			seed:     42,
		},
	}

	for _, v := range tests {
		rand.Seed(v.seed)
		a := StringShuffle(v.actual)
		e := v.expected

		if a != e {
			t.Errorf("Found %v expected %v\n", a, e)
		}
	}
}
