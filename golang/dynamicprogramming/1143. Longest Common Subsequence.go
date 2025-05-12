package dynamicprogramming

import "log"

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m-1; i++ {
		for j := 0; j <= n-1; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func Test_longestCommonSubsequence() {
	text11, text12 := "abcde", "ace"
	ans1 := longestCommonSubsequence(text11, text12)
	log.Printf("ans1: %v", ans1)

	text21, text22 := "abc", "abc"
	ans2 := longestCommonSubsequence(text21, text22)
	log.Printf("ans2: %v", ans2)

	text31, text32 := "abc", "def"
	ans3 := longestCommonSubsequence(text31, text32)
	log.Printf("ans3: %v", ans3)
}
