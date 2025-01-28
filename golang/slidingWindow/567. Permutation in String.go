package slidingWindow

import (
	"log"
)

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

func checkInclusionReview(s1 string, s2 string) bool {
	s1Map := [26]int{}
	s2Map := [26]int{}
	matches := 0
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if s1Map[i] == s2Map[i] {
			matches++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		index := s2[r] - 'a'
		s2Map[index]++
		if s2Map[index] == s1Map[index] {
			matches++
		} else if s2Map[index] == s1Map[index]+1 {
			matches--
		}

		index = s2[l] - 'a'
		s2Map[index]--
		if s2Map[index] == s1Map[index] {
			matches++
		} else if s2Map[index] == s1Map[index]-1 {
			matches--
		}
		l++
	}

	return matches == 26
}

func checkInclusionReview2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	var l = 0
	for r := len(s1); r < len(s2); r++ {
		if isEqual(s1Map, s2Map) {
			return true
		}

		s2Map[s2[r]-'a']++
		s2Map[s2[l]-'a']--

		l++
	}

	return isEqual(s1Map, s2Map)
}

func isEqual(s1, s2 map[byte]int) bool {
	for r := 0; r < 26; r++ {
		if s1[byte(r)] != s2[byte(r)] {
			return false
		}
	}

	return true
}

func Test_CheckInclusion() {
	case1S1, case1S2 := "ab", "eidbaooo"
	ans1 := checkInclusion(case1S1, case1S2)
	log.Printf("ans1: %t", ans1)
	case2S1, case2S2 := "ab", "eidboaoo"
	ans2 := checkInclusion(case2S1, case2S2)
	log.Printf("ans2: %t", ans2)
}

func Test_CheckInclusionReview() {
	//case1S1, case1S2 := "ab", "eidbaooo"
	//ans1 := checkInclusionReview2(case1S1, case1S2)
	//log.Printf("ans1: %t", ans1)
	//case2S1, case2S2 := "ab", "eidboaoo"
	//ans2 := checkInclusionReview2(case2S1, case2S2)
	//log.Printf("ans2: %t", ans2)
	case3S1, case3S2 := "adc", "dcda"
	ans3 := checkInclusionReview2(case3S1, case3S2)
	log.Printf("ans3: %t", ans3)
}
