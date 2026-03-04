package slidingWindow

import (
	"testing"
)

type minWindowCase struct {
	s        string
	t        string
	expected string
}

var minWindowCaseCases = []minWindowCase{
	{s: "ADOBECODEBANC", t: "ABC", expected: "BANC"},
	{s: "a", t: "a", expected: "a"},
	{s: "a", t: "aa", expected: ""},
	{s: "a", t: "", expected: ""},
	{s: "", t: "a", expected: ""},
	{s: "cabwefgewcwaefgcf", t: "cae", expected: "cwae"},
}

func TestMinWindow(t *testing.T) {
	for _, c := range minWindowCaseCases {
		ans := minWindow(c.s, c.t)
		if ans != c.expected {
			t.Errorf("answer is %s, want %s", ans, c.expected)
		}
	}
}

func TestMinWindow2(t *testing.T) {
	for _, c := range minWindowCaseCases {
		ans := minWindow2(c.s, c.t)
		if ans != c.expected {
			t.Errorf("answer is %s, want %s", ans, c.expected)
		}
	}
}

func TestMinWindowReview(t *testing.T) {
	for _, c := range minWindowCaseCases {
		ans := minWindowReview(c.s, c.t)
		if ans != c.expected {
			t.Errorf("answer is %s, want %s", ans, c.expected)
		}
	}
}

func TestMinWindowBruteForce(t *testing.T) {
	for _, c := range minWindowCaseCases {
		ans := minWindowBruteForce(c.s, c.t)
		if ans != c.expected {
			t.Errorf("answer is %s, want %s", ans, c.expected)
		}
	}
}

func TestMinWindowSlidingWindow(t *testing.T) {
	for _, c := range minWindowCaseCases {
		ans := minWindowSlidingWindow(c.s, c.t)
		if ans != c.expected {
			t.Errorf("answer is %s, want %s", ans, c.expected)
		}
	}
}

func TestMinWindowSlidingWindowOptimal(t *testing.T) {
	for _, c := range minWindowCaseCases {
		ans := minWindowSlidingWindowOptimal(c.s, c.t)
		if ans != c.expected {
			t.Errorf("answer is %s, want %s", ans, c.expected)
		}
	}
}
