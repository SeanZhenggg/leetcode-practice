package slidingWindow

import (
	"log"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	case1 := "abcabcbb"
	ans1 := lengthOfLongestSubstring(case1)
	log.Printf("ans1: %v", ans1)

	case2 := "bbbbb"
	ans2 := lengthOfLongestSubstring(case2)
	log.Printf("ans2: %v", ans2)

	case3 := "pwwkew"
	ans3 := lengthOfLongestSubstring(case3)
	log.Printf("ans3: %v", ans3)

	case4 := "pwagewegklew"
	ans4 := lengthOfLongestSubstring(case4)
	log.Printf("ans4: %v", ans4)
}

func TestLengthOfLongestSubstringForLoop(t *testing.T) {
	case1 := "abcabcbb"
	ans1 := lengthOfLongestSubstringDoubleForLoop(case1)
	log.Printf("ans1: %v", ans1)

	case2 := "bbbbb"
	ans2 := lengthOfLongestSubstringDoubleForLoop(case2)
	log.Printf("ans2: %v", ans2)

	case3 := "pwwkew"
	ans3 := lengthOfLongestSubstringDoubleForLoop(case3)
	log.Printf("ans3: %v", ans3)

	case4 := "pwagewegklew"
	ans4 := lengthOfLongestSubstringDoubleForLoop(case4)
	log.Printf("ans4: %v", ans4)
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	case1 := "abcabcbb"
	ans1 := lengthOfLongestSubstringSlidingWindow(case1)
	log.Printf("ans1: %v", ans1)

	case2 := "bbbbb"
	ans2 := lengthOfLongestSubstringSlidingWindow(case2)
	log.Printf("ans2: %v", ans2)

	case3 := "pwwkew"
	ans3 := lengthOfLongestSubstringSlidingWindow(case3)
	log.Printf("ans3: %v", ans3)

	case4 := "pwagewegklew"
	ans4 := lengthOfLongestSubstringSlidingWindow(case4)
	log.Printf("ans4: %v", ans4)
}
