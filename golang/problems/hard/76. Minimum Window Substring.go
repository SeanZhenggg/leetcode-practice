package hard

import (
	"log"
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
