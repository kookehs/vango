package algorithm

import "testing"

func TestFuzzyStringSearch(t *testing.T) {
	var tests = []struct {
		s1, s2   string
		expected float64
	}{
		{
			s1:       "hello",
			s2:       "hello",
			expected: 1,
		},
	}

	for _, v := range tests {
		e := v.expected
		f := FuzzyStringSearch(v.s1, v.s2)

		if f != e {
			t.Errorf("Found %v expect %v\n", f, e)
		}
	}
}
