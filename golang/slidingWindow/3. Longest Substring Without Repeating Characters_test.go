package slidingWindow

import "testing"

type longestSubstringCase struct {
	s        string
	expected int
}

var longestSubstringCases = []longestSubstringCase{
	{s: "abcabcbb", expected: 3},
	{s: "bbbbb", expected: 1},
	{s: "pwwkew", expected: 3},
	{s: "pwagewegklew", expected: 5},
}

func TestLengthOfLongestSubstringForLoop(t *testing.T) {
	for _, c := range longestSubstringCases {
		ans := lengthOfLongestSubstringDoubleForLoop(c.s)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestLengthOfLongestSubstringSlidingWindow(t *testing.T) {
	for _, c := range longestSubstringCases {
		ans := lengthOfLongestSubstringSlidingWindow(c.s)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
