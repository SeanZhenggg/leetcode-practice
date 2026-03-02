package slidingWindow

// failed solution
func characterReplacementFailed(s string, k int) int {
	var maxLength int
	var l int
	var m = make(map[byte]int)
	var curK = k
	for r := 0; r < len(s); r++ {
		m[s[r]] += 1
		currLen := r - l + 1
		if s[l] != s[r] {
			if curK > 0 {
				curK -= 1
			} else {
				cur := s[l]
				for l < r && s[l] == cur {
					m[s[l]] -= 1
					l++
				}
				// reset k
				currLen = r - l + 1
				// available k = k - (current length - current element count)
				curK = k - (currLen - m[s[l]])
			}
		}

		if currLen > maxLength {
			maxLength = currLen
		}
	}

	m = make(map[byte]int)
	l = len(s) - 1
	curK = k
	for r := len(s) - 1; r >= 0; r-- {
		m[s[r]] += 1
		currLen := l - r + 1
		if s[l] != s[r] {
			if curK > 0 {
				curK -= 1
			} else {
				cur := s[l]
				for l > r && s[l] == cur {
					m[s[l]] -= 1
					l--
				}
				// reset k
				currLen = l - r + 1
				// available k = k - (current length - current element count)
				curK = k - (currLen - m[s[l]])
			}
		}

		if currLen > maxLength {
			maxLength = currLen
		}
	}

	return maxLength
}

// brute force
func characterReplacementBruteForce(s string, k int) int {
	maxFreq, m := 0, make(map[byte]int)
	res := 0
	for i := 0; i < len(s); i++ {
		maxFreq, m = 0, make(map[byte]int)
		for j := i; j < len(s); j++ {
			m[s[j]] += 1
			maxFreq = max(maxFreq, m[s[j]])
			if (j-i+1)-maxFreq <= k {
				res = max(res, j-i+1)
			}
		}
	}

	return res
}

// sliding window
func characterReplacementSlidingWindow(s string, k int) int {
	// 當前長度 - l 位置的字母出現次數 <= k 繼續
	// 否則移動 l 的位置, 然後接著下一次判斷

	m := make(map[byte]int)
	l, r := 0, 0
	maxLength := 0
	for r < len(s) {
		m[s[r]]++
		maxFreq := getMaxFreq(m)

		for l < r && (r-l+1)-maxFreq > k {
			m[s[l]]--
			l++
		}

		maxLength = max(maxLength, r-l+1)
		r++
	}

	return maxLength
}

func getMaxFreq(m map[byte]int) int {
	maxFreq := 0
	for _, v := range m {
		if v > maxFreq {
			maxFreq = v
		}
	}
	return maxFreq
}

// sliding window optimal:
// no need to update maxC at each iteration
// because of the equation: length - maxC <= k
// k is constant, which tells that we could get greater length only if maxC is greater than before, not less than before
// so there is no need for decrement maxC
func characterReplacementSlidingWindowOptimal(s string, k int) int {
	var maxLength int
	var l int
	var m = make(map[byte]int)
	var maxC int
	for r := 0; r < len(s); r++ {
		m[s[r]] += 1

		maxC = max(maxC, m[s[r]])

		for (r-l+1)-maxC > k {
			m[s[l]] -= 1
			l++
		}

		if (r - l + 1) > maxLength {
			maxLength = r - l + 1
		}
	}

	return maxLength
}
