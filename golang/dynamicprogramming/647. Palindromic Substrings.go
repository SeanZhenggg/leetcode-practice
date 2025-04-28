package dynamicprogramming

import "log"

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	count := 0
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(dp); i++ {
		dp[i][i] = true
		count++
	}

	for length := 1; length < len(s); length++ {
		for i := 0; i < len(s)-length; i++ {
			j := i + length
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				count++
				dp[i][j] = true
			}
		}
	}

	return count
}

func Test_countSubstrings() {
	s1 := "abc"
	ans1 := countSubstrings(s1)
	log.Printf("ans1: %v", ans1)

	s2 := "aaa"
	ans2 := countSubstrings(s2)
	log.Printf("ans2: %v", ans2)
	s3 := "babad"
	ans3 := countSubstrings(s3)
	log.Printf("ans2: %v", ans3)

	s4 := "cbbd"
	ans4 := countSubstrings(s4)
	log.Printf("ans2: %v", ans4)

	s5 := "baabad"
	ans5 := countSubstrings(s5)
	log.Printf("ans2: %v", ans5)

	s6 := "baaabbad"
	ans6 := countSubstrings(s6)
	log.Printf("ans2: %v", ans6)

	s7 := "aaaa"
	ans7 := countSubstrings(s7)
	log.Printf("ans2: %v", ans7)

}
