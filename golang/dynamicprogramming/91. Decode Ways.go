package dynamicprogramming

import (
	"log"
	"strconv"
)

func numDecodings(s string) int {
	var dp = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = -1
	}
	var dfs func(idx int) int
	dfs = func(idx int) (num int) {
		if idx < 0 {
			return 1
		}

		if dp[idx] != -1 {
			return dp[idx]
		}
		var countA int
		var countB int
		if idx >= 0 && isValid(s[idx:idx+1]) {
			countA = dfs(idx - 1)
		}

		if idx > 0 && isValid(s[idx-1:idx+1]) {
			countB = dfs(idx - 2)
		}

		defer func() {
			dp[idx] = num
		}()
		return countA + countB
	}

	return dfs(len(s) - 1)
}

func numDecodings2(s string) int {
	var dp = make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		dp[i] = -1
	}
	var dfs func(idx int) int
	dfs = func(idx int) (num int) {
		if idx >= len(s) {
			return 1
		}

		if dp[idx] != -1 {
			return dp[idx]
		}
		var countA int
		var countB int
		if idx >= 0 && isValid2(s[idx:idx+1]) {
			countA = dfs(idx + 1)
		}

		if idx < len(s)-1 && isValid2(s[idx:idx+2]) {
			countB = dfs(idx + 2)
		}

		defer func() {
			dp[idx] = num
		}()
		return countA + countB
	}

	return dfs(0)
}

func numDecodings3(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		twoDigits, _ := strconv.Atoi(s[i-2 : i])
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func isValid(s string) bool {
	if len(s) == 2 && s[0] == '0' {
		return false
	}
	i, _ := strconv.Atoi(s)
	return i > 0 && i <= 26
}

func isValid2(s string) bool {
	if len(s) == 1 && s[0] == '0' {
		return false
	}

	if len(s) == 1 {
		return true
	}

	i, _ := strconv.Atoi(s)
	return i >= 10 && i <= 26
}

func Test_numDecodings() {
	s1 := "12"
	ans1 := numDecodings(s1)
	log.Printf("ans1: %v", ans1)

	s2 := "225"
	ans2 := numDecodings(s2)
	log.Printf("ans2: %v", ans2)
	//
	s3 := "111111111111111111111111111111111111111111111"
	ans3 := numDecodings(s3)
	log.Printf("ans3: %v", ans3)

	s4 := "11106"
	ans4 := numDecodings(s4)
	log.Printf("ans4: %v", ans4)

	s5 := "1106"
	ans5 := numDecodings(s5)
	log.Printf("ans5: %v", ans5)
}

func Test_numDecodings2() {
	s1 := "12"
	ans1 := numDecodings2(s1)
	log.Printf("ans1: %v", ans1)

	s2 := "2326"
	ans2 := numDecodings2(s2)
	log.Printf("ans2: %v", ans2)

	s3 := "111111111111111111111111111111111111111111111"
	ans3 := numDecodings2(s3)
	log.Printf("ans3: %v", ans3)

	s4 := "11106"
	ans4 := numDecodings2(s4)
	log.Printf("ans4: %v", ans4)

	s5 := "1106"
	ans5 := numDecodings2(s5)
	log.Printf("ans5: %v", ans5)
}

func Test_numDecodings3() {
	s1 := "12"
	ans1 := numDecodings3(s1)
	log.Printf("ans1: %v", ans1)

	s2 := "2326"
	ans2 := numDecodings3(s2)
	log.Printf("ans2: %v", ans2)

	s3 := "111111111111111111111111111111111111111111111"
	ans3 := numDecodings3(s3)
	log.Printf("ans3: %v", ans3)

	s4 := "11106"
	ans4 := numDecodings3(s4)
	log.Printf("ans4: %v", ans4)

	s5 := "1106"
	ans5 := numDecodings3(s5)
	log.Printf("ans5: %v", ans5)
}
