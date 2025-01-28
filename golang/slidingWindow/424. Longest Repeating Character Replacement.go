package slidingWindow

import (
	"log"
)

// failed solution
func characterReplacement(s string, k int) int {
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

func characterReplacement2(s string, k int) int {
	var maxLength int
	var l int
	var m = make(map[byte]int)
	for r := 0; r < len(s); r++ {
		m[s[r]] += 1

		for {
			var maxC int
			for i, count := range m {
				if m[i] > maxC {
					maxC = count
				}
			}

			if (r-l+1)-maxC <= k {
				break
			}
			m[s[l]] -= 1
			l++
		}

		if (r - l + 1) > maxLength {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func characterReplacement3(s string, k int) int {
	var maxLength int
	var l int
	var m = make(map[byte]int)
	var maxC int
	for r := 0; r < len(s); r++ {
		m[s[r]] += 1

		if m[s[r]] > maxC {
			maxC = m[s[r]]
		}
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

func characterReplacementReview(s string, k int) int {
	maxLength := 0
	m := make(map[byte]int)
	l := 0

	for r := 0; r < len(s); r++ {
		m[s[r]]++

		maxF := 0
		for _, c := range m {
			if maxF < c {
				maxF = c
			}
		}
		for (r-l+1)-maxF > k {
			m[s[l]]--
			l++

			for _, c := range m {
				if maxF < c {
					maxF = c
				}
			}
		}

		if maxLength < r-l+1 {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func characterReplacementReview2(s string, k int) int {
	maxLength := 0
	m := make(map[byte]int)
	l := 0
	maxF := 0

	for r := 0; r < len(s); r++ {
		m[s[r]]++

		if maxF < m[s[r]] {
			maxF = m[s[r]]
		}
		for (r-l+1)-maxF > k {
			m[s[l]]--
			l++
		}

		if maxLength < r-l+1 {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func characterReplacementReview3(s string, k int) int {
	maxF := 0
	m := make(map[byte]int)
	l := 0
	maxL := 0
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		for _, v := range m {
			maxF = max(maxF, v)
		}

		for r-l+1-maxF > k {
			m[s[l]]--
			l++
			for _, v := range m {
				maxF = max(maxF, v)
			}
		}
		maxL = max(maxL, r-l+1)
	}

	return maxL
}

func characterReplacementReview4(s string, k int) int {
	maxF := 0
	m := make(map[byte]int)
	l := 0
	maxL := 0
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		if maxF < m[s[r]] {
			maxF = m[s[r]]
		}

		for r-l+1-maxF > k {
			m[s[l]]--
			l++
		}
		maxL = max(maxL, r-l+1)
	}

	return maxL
}

func Test_CharacterReplacement() {
	case1 := "ABAB"
	k1 := 2
	ans1 := characterReplacement(case1, k1)
	log.Printf("ans1: %d", ans1)
	case2 := "AABABBA"
	k2 := 1
	ans2 := characterReplacement(case2, k2)
	log.Printf("ans2: %d", ans2)
	case3 := "ABCBDHBA"
	k3 := 2
	ans3 := characterReplacement(case3, k3)
	log.Printf("ans3: %d", ans3)
	case4 := "ABBCBHBA"
	k4 := 3
	ans4 := characterReplacement(case4, k4)
	log.Printf("ans4: %d", ans4)
	case5 := "ABBB"
	k5 := 2
	ans5 := characterReplacement(case5, k5)
	log.Printf("ans5: %d", ans5)
	case6 := "BAAAB"
	k6 := 2
	ans6 := characterReplacement(case6, k6)
	log.Printf("ans6: %d", ans6)
	case7 := "BAAABC"
	k7 := 2
	ans7 := characterReplacement2(case7, k7)
	log.Printf("ans7: %d", ans7)
}

func Test_CharacterReplacement2() {
	case1 := "ABAB"
	k1 := 2
	ans1 := characterReplacement2(case1, k1)
	log.Printf("ans1: %d", ans1)
	case2 := "AABABBA"
	k2 := 1
	ans2 := characterReplacement2(case2, k2)
	log.Printf("ans2: %d", ans2)
	case3 := "ABCBDHBA"
	k3 := 2
	ans3 := characterReplacement2(case3, k3)
	log.Printf("ans3: %d", ans3)
	case4 := "ABBCBHBA"
	k4 := 3
	ans4 := characterReplacement2(case4, k4)
	log.Printf("ans4: %d", ans4)
	case5 := "ABBB"
	k5 := 2
	ans5 := characterReplacement2(case5, k5)
	log.Printf("ans5: %d", ans5)
	case6 := "BAAAB"
	k6 := 2
	ans6 := characterReplacement2(case6, k6)
	log.Printf("ans6: %d", ans6)
	case7 := "BAAABC"
	k7 := 2
	ans7 := characterReplacement2(case7, k7)
	log.Printf("ans7: %d", ans7)
}

func Test_CharacterReplacement3() {
	case1 := "ABAB"
	k1 := 2
	ans1 := characterReplacement3(case1, k1)
	log.Printf("ans1: %d", ans1)
	case2 := "AABABBA"
	k2 := 1
	ans2 := characterReplacement3(case2, k2)
	log.Printf("ans2: %d", ans2)
	case3 := "ABCBDHBA"
	k3 := 2
	ans3 := characterReplacement3(case3, k3)
	log.Printf("ans3: %d", ans3)
	case4 := "ABBCBHBA"
	k4 := 3
	ans4 := characterReplacement3(case4, k4)
	log.Printf("ans4: %d", ans4)
	case5 := "ABBB"
	k5 := 2
	ans5 := characterReplacement3(case5, k5)
	log.Printf("ans5: %d", ans5)
	case6 := "BAAAB"
	k6 := 2
	ans6 := characterReplacement3(case6, k6)
	log.Printf("ans6: %d", ans6)
	case7 := "BAAABC"
	k7 := 2
	ans7 := characterReplacement3(case7, k7)
	log.Printf("ans7: %d", ans7)
}

func Test_CharacterReplacementReview() {
	case1 := "ABAB"
	k1 := 2
	ans1 := characterReplacementReview(case1, k1)
	log.Printf("ans1: %d", ans1)
	case2 := "AABABBA"
	k2 := 1
	ans2 := characterReplacementReview(case2, k2)
	log.Printf("ans2: %d", ans2)
	case3 := "ABCBDHBA"
	k3 := 2
	ans3 := characterReplacementReview(case3, k3)
	log.Printf("ans3: %d", ans3)
	case4 := "ABBCBHBA"
	k4 := 3
	ans4 := characterReplacementReview(case4, k4)
	log.Printf("ans4: %d", ans4)
	case5 := "ABBB"
	k5 := 2
	ans5 := characterReplacementReview(case5, k5)
	log.Printf("ans5: %d", ans5)
	case6 := "BAAAB"
	k6 := 2
	ans6 := characterReplacementReview(case6, k6)
	log.Printf("ans6: %d", ans6)
	case7 := "BAAABC"
	k7 := 2
	ans7 := characterReplacementReview(case7, k7)
	log.Printf("ans7: %d", ans7)
}

func Test_CharacterReplacementReview2() {
	case1 := "ABAB"
	k1 := 2
	ans1 := characterReplacementReview2(case1, k1)
	log.Printf("ans1: %d", ans1)
	case2 := "AABABBA"
	k2 := 1
	ans2 := characterReplacementReview2(case2, k2)
	log.Printf("ans2: %d", ans2)
	case3 := "ABCBDHBA"
	k3 := 2
	ans3 := characterReplacementReview2(case3, k3)
	log.Printf("ans3: %d", ans3)
	case4 := "ABBCBHBA"
	k4 := 3
	ans4 := characterReplacementReview2(case4, k4)
	log.Printf("ans4: %d", ans4)
	case5 := "ABBB"
	k5 := 2
	ans5 := characterReplacementReview2(case5, k5)
	log.Printf("ans5: %d", ans5)
	case6 := "BAAAB"
	k6 := 2
	ans6 := characterReplacementReview2(case6, k6)
	log.Printf("ans6: %d", ans6)
	case7 := "BAAABC"
	k7 := 2
	ans7 := characterReplacementReview2(case7, k7)
	log.Printf("ans7: %d", ans7)
}

func Test_CharacterReplacementReview3() {
	case1 := "ABAB"
	k1 := 2
	ans1 := characterReplacementReview3(case1, k1)
	log.Printf("ans1: %d", ans1)
	case2 := "AABABBA"
	k2 := 1
	ans2 := characterReplacementReview3(case2, k2)
	log.Printf("ans2: %d", ans2)
	case3 := "ABCBDHBA"
	k3 := 2
	ans3 := characterReplacementReview3(case3, k3)
	log.Printf("ans3: %d", ans3)
	case4 := "ABBCBHBA"
	k4 := 3
	ans4 := characterReplacementReview3(case4, k4)
	log.Printf("ans4: %d", ans4)
	case5 := "ABBB"
	k5 := 2
	ans5 := characterReplacementReview3(case5, k5)
	log.Printf("ans5: %d", ans5)
	case6 := "BAAAB"
	k6 := 2
	ans6 := characterReplacementReview3(case6, k6)
	log.Printf("ans6: %d", ans6)
	case7 := "BAAABC"
	k7 := 2
	ans7 := characterReplacementReview3(case7, k7)
	log.Printf("ans7: %d", ans7)
}
