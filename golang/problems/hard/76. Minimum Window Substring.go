package hard

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
		for l <= r && isValidSubstr(subStrMap, tMap) {
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
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	have := 0
	need := len(tMap)
	l := 0
	minLength := math.MaxInt16
	ret := ""
	for r := 0; r < len(s); r++ {
		sMap[s[r]]++
		if sMap[s[r]] == tMap[s[r]] {
			have++
		}

		for have == need {
			sMap[s[l]]--
			if sMap[s[l]] < tMap[s[l]] {
				have--
			}

			if r-l+1 < minLength {
				minLength = r - l + 1
				ret = s[l : r+1]
			}

			l++
		}
	}

	return ret
}

func minWindowReview2(s string, t string) string {
	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		s1Map[t[i]] += 1
	}
	need := len(s1Map)
	have := 0
	l := 0
	minLength := math.MaxInt16
	ret := ""
	for r := 0; r < len(s); r++ {
		s2Map[s[r]]++

		if s1Map[s[r]] == s2Map[s[r]] {
			have++
		}

		for have == need {
			if r-l+1 < minLength {
				minLength = r - l + 1
				ret = s[l : r+1]
			}

			s2Map[s[l]]--
			if s2Map[s[l]] < s1Map[s[l]] {
				have--
			}

			l++
		}
	}
	return ret
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
	ans1 := minWindowReview2(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := "a"
	t2 := "a"
	ans2 := minWindowReview2(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := "a"
	t3 := "aa"
	ans3 := minWindowReview2(case3, t3)
	log.Printf("ans3: %v", ans3)
	case4 := "cabwefgewcwaefgcf"
	t4 := "cae"
	ans4 := minWindowReview2(case4, t4)
	log.Printf("ans4: %v", ans4)
}
