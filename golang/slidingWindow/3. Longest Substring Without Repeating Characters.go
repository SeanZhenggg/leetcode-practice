package slidingWindow

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

func lengthOfLongestSubstringReview(s string) int {
	// a map to record character's existence
	m := make(map[byte]bool)
	l := 0
	maxLength := 0

	// loop each character
	//     loop to check if there was same character in the map
	//        1. remove character existence from map
	//        then shift l to the right
	//
	//     mark each character as existed

	for r := 0; r < len(s); r++ {
		for l < r && m[s[r]] {
			m[s[l]] = false
			l++
		}
		curLength := r - l + 1
		if curLength > maxLength {
			maxLength = curLength
		}

		m[s[r]] = true
	}

	return maxLength
}
