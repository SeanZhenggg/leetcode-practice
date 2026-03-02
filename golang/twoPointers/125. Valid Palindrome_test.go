package twoPointers

import "testing"

type isPalindromeCase struct {
	s        string
	expected bool
}

var isPalindromeCases = []isPalindromeCase{
	{s: "A man, a plan, a canal: Panama", expected: true},
	{s: "race a car", expected: false},
	{s: " ", expected: true},
}

func TestIsPalindrome(t *testing.T) {
	for _, c := range isPalindromeCases {
		ans := isPalindrome(c.s)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
