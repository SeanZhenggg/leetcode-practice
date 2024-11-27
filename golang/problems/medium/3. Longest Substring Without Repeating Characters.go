package medium

import "log"

// solution using dynamic sliding window + hash map
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	var l int
	var maxLength int
	m := make(map[byte]bool)
	for r := 0; r < len(s); r++ {
		for l < r && m[s[r]] {
			m[s[l]] = false
			l++
		}

		m[s[r]] = true

		if r-l+1 > maxLength {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func Test_LengthOfLongestSubstring() {
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
