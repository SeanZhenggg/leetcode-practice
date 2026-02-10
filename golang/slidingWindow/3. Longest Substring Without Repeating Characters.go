package slidingWindow

// 2 for loop
func lengthOfLongestSubstringDoubleForLoop(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		m := map[byte]bool{}
		for j := i; j < len(s); j++ {
			if m[s[j]] {
				longest = max(longest, j-i)
				break
			}
			m[s[j]] = true
		}
	}

	return longest
}

func lengthOfLongestSubstringSlidingWindow(s string) int {
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

		maxLength = max(maxLength, r-l+1)
	}

	return maxLength
}
