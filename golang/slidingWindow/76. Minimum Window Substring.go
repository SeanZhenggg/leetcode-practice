package slidingWindow

import (
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

// brute force, TLE error
func minWindowBruteForce(s string, t string) string {
	m := make(map[byte]int)
	for _, i2 := range t {
		m[byte(i2)]++
	}
	res := 100_000
	res2 := ""

	for i := 0; i < len(s); i++ {
		mForS := make(map[byte]int)
		for j := i; j < len(s); j++ {
			mForS[s[j]]++

			if mapIsEqual(m, mForS) {
				if res > j-i+1 {
					res = j - i + 1
					res2 = s[i : j+1]
				}
				break
			}
		}
	}

	return res2
}

// sliding window
func minWindowSlidingWindow(s string, t string) string {
	l := 0
	mForT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mForT[t[i]] += 1
	}
	mForS := make(map[byte]int)

	minLen := 100_000
	minS := ""
	for r := 0; r < len(s); r++ {
		mForS[s[r]]++

		for mapIsEqual(mForT, mForS) {
			if minLen > r-l+1 {
				minLen = r - l + 1
				minS = s[l : r+1]
			}
			mForS[s[l]] -= 1
			l++
		}
	}
	return minS
}

func minWindowSlidingWindowOptimal(s string, t string) string {
	// edge case
	if len(t) == 0 {
		return ""
	}

	mForT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mForT[t[i]] += 1
	}

	mForS := make(map[byte]int)
	minS := ""
	have, need := 0, len(mForT)
	l := 0
	for r := 0; r < len(s); r++ {
		mForS[s[r]] += 1

		if mForT[s[r]] > 0 && mForS[s[r]] == mForT[s[r]] {
			have += 1
		}

		for have == need {
			// use minS == "" to omit for another variable recording min length
			if minS == "" || len(minS) > r-l+1 {
				minS = s[l : r+1]
			}

			mForS[s[l]] -= 1
			if mForT[s[l]] > 0 && mForS[s[l]] < mForT[s[l]] {
				have -= 1
			}
			l++
		}
	}

	return minS
}

func mapIsEqual(m map[byte]int, m2 map[byte]int) bool {
	for k, v := range m {
		if m2[k] < v {
			return false
		}
	}
	return true
}
