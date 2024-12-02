package medium

import "log"

func checkInclusion(s1 string, s2 string) bool {
	s1Count := [26]int{}
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Count := [26]int{}
		for j := i; j < i+len(s1); j++ {
			s2Count[s2[j]-'a']++
		}

		if s1Count == s2Count {
			return true
		}
	}
	return false
}

func Test_CheckInclusion() {
	case1S1, case1S2 := "ab", "eidbaooo"
	ans1 := checkInclusion(case1S1, case1S2)
	log.Printf("ans1: %t", ans1)
	case2S1, case2S2 := "ab", "eidboaoo"
	ans2 := checkInclusion(case2S1, case2S2)
	log.Printf("ans2: %t", ans2)
}
