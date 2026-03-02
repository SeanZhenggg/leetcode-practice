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

func TestCharacterReplacement(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacement(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacement1(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacement1(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacement2(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacement2(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacement3(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacement3(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementReview(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementReview(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementReview2(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementReview2(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementReview3(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementReview3(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestCharacterReplacementReview5(t *testing.T) {
	for _, c := range cases {
		ans := characterReplacementReview5(c.s, c.k)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
