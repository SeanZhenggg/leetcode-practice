package dynamicprogramming

import (
	"log"
)

// two pointer solution - TC:O(n^2), SC: O(1)
func longestPalindrome(s string) string {
	maxLength := 0
	maxSubStr := ""
	for i, _ := range s {
		l, r := i, i
		// shift at the beginning for duplicated adjacent string, include substring of both odd/even length cases
		// could separate even/odd length cases as well, see below longestPalindrome1
		for r < len(s)-1 && s[r] == s[r+1] {
			r++
		}
		for l > 0 && s[l] == s[l-1] {
			l--
		}

		if r-l+1 > maxLength {
			maxLength = r - l + 1
			maxSubStr = s[l : r+1]
			maxLength = max(maxLength, r-l+1)
		}

		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			if r-l+1 > maxLength {
				maxLength = r - l + 1
				maxSubStr = s[l : r+1]
				maxLength = max(maxLength, r-l+1)
			}

			l--
			r++
		}

	}

	return maxSubStr
}

func longestPalindrome1(s string) string {
	maxSubStr := ""
	for i, _ := range s {
		// odd case
		l, r := i, i
		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			if r-l+1 > len(maxSubStr) {
				maxSubStr = s[l : r+1]
			}

			l--
			r++
		}

		// even case
		l, r = i, i+1
		for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
			if r-l+1 > len(maxSubStr) {
				maxSubStr = s[l : r+1]
			}

			l--
			r++
		}
	}

	return maxSubStr
}

// two pointer "Check All Substrings" solution(editorial) - TC: O(n^3), SC: O(1)
func longestPalindrome2(s string) string {
	var check func(i, j int) bool
	check = func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	for length := len(s); length > 0; length-- {
		for left := 0; left < len(s)-length+1; left++ {
			if check(left, left+length-1) {
				return s[left : left+length]
			}
		}
	}
	return ""
}

// dp solution - TC: O(n^2), SC: O(n^2)
func longestPalindrome3(s string) string {
	if len(s) == 1 {
		return s
	}

	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	ans := ""

	for length := 1; length <= len(s); length++ {
		for left := 0; left <= len(s)-length; left++ {
			right := left + length - 1

			if s[left] == s[right] && (right-left < 2 || dp[left+1][right-1]) {
				dp[left][right] = true
				if (right - left + 1) > len(ans) {
					ans = s[left : right+1]
				}
			}
		}
	}
	return s[ans[0] : ans[1]+1]
}

func Test_longestPalindrome() {
	s1 := "babad"
	ans1 := longestPalindrome(s1)
	log.Printf("ans1: %v", ans1)
	s2 := "cbbd"
	ans2 := longestPalindrome(s2)
	log.Printf("ans2: %v", ans2)
	s3 := "baabad"
	ans3 := longestPalindrome(s3)
	log.Printf("ans3: %v", ans3)
	s4 := "baaabbad"
	ans4 := longestPalindrome(s4)
	log.Printf("ans4: %v", ans4)
	s5 := "aaaa"
	ans5 := longestPalindrome(s5)
	log.Printf("ans5: %v", ans5)
	s6 := "a"
	ans6 := longestPalindrome(s6)
	log.Printf("ans6: %v", ans6)
	s7 := "ac"
	ans7 := longestPalindrome(s7)
	log.Printf("ans7: %v", ans7)
}

func Test_longestPalindrome1() {
	s1 := "babad"
	ans1 := longestPalindrome1(s1)
	log.Printf("ans1: %v", ans1)
	s2 := "cbbd"
	ans2 := longestPalindrome1(s2)
	log.Printf("ans2: %v", ans2)
	s3 := "baabad"
	ans3 := longestPalindrome1(s3)
	log.Printf("ans3: %v", ans3)
	s4 := "baaabbad"
	ans4 := longestPalindrome1(s4)
	log.Printf("ans4: %v", ans4)
	s5 := "aaaa"
	ans5 := longestPalindrome1(s5)
	log.Printf("ans5: %v", ans5)
	s6 := "a"
	ans6 := longestPalindrome1(s6)
	log.Printf("ans6: %v", ans6)
	s7 := "ac"
	ans7 := longestPalindrome1(s7)
	log.Printf("ans7: %v", ans7)
}

func Test_longestPalindrome2() {
	s1 := "babad"
	ans1 := longestPalindrome2(s1)
	log.Printf("ans1: %v", ans1)
	s2 := "cbbd"
	ans2 := longestPalindrome2(s2)
	log.Printf("ans2: %v", ans2)
	s3 := "baabad"
	ans3 := longestPalindrome2(s3)
	log.Printf("ans3: %v", ans3)
	s4 := "baaabbad"
	ans4 := longestPalindrome2(s4)
	log.Printf("ans4: %v", ans4)
	s5 := "aaaa"
	ans5 := longestPalindrome2(s5)
	log.Printf("ans5: %v", ans5)
	s6 := "a"
	ans6 := longestPalindrome2(s6)
	log.Printf("ans6: %v", ans6)
	s7 := "ac"
	ans7 := longestPalindrome2(s7)
	log.Printf("ans7: %v", ans7)
}

func Test_longestPalindrome3() {
	s1 := "babad"
	ans1 := longestPalindrome3(s1)
	log.Printf("ans1: %v", ans1)
	s2 := "cbbd"
	ans2 := longestPalindrome3(s2)
	log.Printf("ans2: %v", ans2)
	s3 := "baabad"
	ans3 := longestPalindrome3(s3)
	log.Printf("ans3: %v", ans3)
	s4 := "baaabbad"
	ans4 := longestPalindrome3(s4)
	log.Printf("ans4: %v", ans4)
	s5 := "aaaa"
	ans5 := longestPalindrome3(s5)
	log.Printf("ans5: %v", ans5)
	s6 := "a"
	ans6 := longestPalindrome3(s6)
	log.Printf("ans6: %v", ans6)
	s7 := "ac"
	ans7 := longestPalindrome3(s7)
	log.Printf("ans7: %v", ans7)
}
