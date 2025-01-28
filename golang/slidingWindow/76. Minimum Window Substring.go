package slidingWindow

import (
	"log"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tMap := [52]int{}
	subStrMap := [52]int{}

	for i := 0; i < len(t); i++ {
		if t[i] >= 65 && t[i] <= 90 {
			tMap[t[i]-'A']++
		} else {
			tMap[t[i]+26-'a']++
		}
	}

	l := 0
	minSubStr := ""
	set := false
	for r := 0; r < len(s); r++ {
		if s[r] >= 65 && s[r] <= 90 {
			subStrMap[s[r]-'A']++
		} else {
			subStrMap[s[r]+26-'a']++
		}
		for isValidSubstr(subStrMap, tMap) {
			if !set || r-l+1 < len(minSubStr) {
				set = true
				minSubStr = s[l : r+1]
			}

			if s[l] >= 65 && s[l] <= 90 {
				subStrMap[s[l]-'A']--
			} else {
				subStrMap[s[l]+26-'a']--
			}
			l++

		}
	}

	return minSubStr
}

func isValidSubstr(substr [52]int, t [52]int) bool {
	for i := 0; i < len(substr); i++ {
		if t[i] > 0 && substr[i] < t[i] {
			return false
		}
	}
	return true
}

func minWindow2(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	subStrMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	need := len(tMap)
	l := 0
	subStrLen := math.MaxInt16
	minSubStr := ""
	have := 0

	for r := 0; r < len(s); r++ {
		subStrMap[s[r]]++

		if subStrMap[s[r]] == tMap[s[r]] {
			have += 1
		}

		for have == need {
			if r-l+1 < subStrLen {
				subStrLen = r - l + 1
				minSubStr = s[l : r+1]
			}

			subStrMap[s[l]]--
			if subStrMap[s[l]] < tMap[s[l]] {
				have--
			}
			l++
		}
	}

	return minSubStr
}

func minWindowReview(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tMap := make(map[byte]int)
	hMap := make(map[byte]int)
	for r := 0; r < len(t); r++ {
		tMap[t[r]]++
	}

	have, need := 0, len(tMap)
	l := 0
	minLength := math.MaxInt16
	substr := ""
	for r := 0; r < len(s); r++ {
		hMap[s[r]]++

		if hMap[s[r]] == tMap[s[r]] {
			have++
		}

		if r+1 >= len(t) {
			for have == need {
				hMap[s[l]]--

				if minLength > r-l+1 {
					substr = s[l : r+1]
					minLength = r - l + 1
				}

				if hMap[s[l]] < tMap[s[l]] {
					have -= 1
				}
				l++
			}
		}
	}
	return substr
}

func Test_MinWindow() {
	case1 := "ADOBECODEBANC"
	t1 := "ABC"
	ans1 := minWindow(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := "a"
	t2 := "a"
	ans2 := minWindow(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := "a"
	t3 := "aa"
	ans3 := minWindow(case3, t3)
	log.Printf("ans3: %v", ans3)
	case4 := "cabwefgewcwaefgcf"
	t4 := "cae"
	ans4 := minWindowReview(case4, t4)
	log.Printf("ans4: %v", ans4)
}

func Test_MinWindow2() {
	case1 := "ADOBECODEBANC"
	t1 := "ABC"
	ans1 := minWindow2(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := "a"
	t2 := "a"
	ans2 := minWindow2(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := "a"
	t3 := "aa"
	ans3 := minWindow2(case3, t3)
	log.Printf("ans3: %v", ans3)
}

func Test_MinWindowReview() {
	case1 := "ADOBECODEBANC"
	t1 := "ABC"
	ans1 := minWindowReview(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := "a"
	t2 := "a"
	ans2 := minWindowReview(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := "a"
	t3 := "aa"
	ans3 := minWindowReview(case3, t3)
	log.Printf("ans3: %v", ans3)
	case4 := "cabwefgewcwaefgcf"
	t4 := "cae"
	ans4 := minWindowReview(case4, t4)
	log.Printf("ans4: %v", ans4)
}
