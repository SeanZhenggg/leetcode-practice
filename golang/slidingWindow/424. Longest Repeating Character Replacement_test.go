package slidingWindow

import (
	"testing"
)

type testcase struct {
	s        string
	k        int
	expected int
}

var (
	cases = []testcase{
		{s: "ABAB", k: 2, expected: 4},
		{s: "AABABBA", k: 1, expected: 4},
		{s: "ABCBDHBA", k: 2, expected: 4},
		{s: "ABBCBHBA", k: 3, expected: 7},
		{s: "ABBB", k: 2, expected: 4},
		{s: "BAAAB", k: 2, expected: 5},
		{s: "BAAABC", k: 2, expected: 5},
	}
)

func TestCharacterReplacementFailed(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementFailed(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementBruteForce(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementBruteForce(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementSlidingWindow(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementSlidingWindow(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementSlidingWindowOptimal(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementSlidingWindowOptimal(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
