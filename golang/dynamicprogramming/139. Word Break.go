package dynamicprogramming

import (
	"log"
)

func wordBreak(s string, wordDict []string) bool {
	var dfs func(st, end int) bool

	dfs = func(st, end int) bool {
		if st == end {
			return true
		}

		for i := st; i < end; i++ {
			for j := 0; j < len(wordDict); j++ {
				if wordDict[j] == s[st:i+1] && dfs(i+1, end) {
					return true
				}
			}
		}

		return false
	}
	return dfs(0, len(s))
}

func wordBreak1(s string, wordDict []string) bool {
	var dfs func(st int) bool

	dfs = func(st int) (ret bool) {
		if st == len(s) {
			return true
		}

		for _, w := range wordDict {
			if len(s[st:]) >= len(w) && w == s[st:st+len(w)] {
				if dfs(st + len(w)) {
					return true
				}
			}
		}

		return false
	}

	return dfs(0)
}

func wordBreak2(s string, wordDict []string) bool {
	var dfs func(st int) bool
	memo := make([]int, len(s)+1)
	dfs = func(st int) (ret bool) {
		if memo[st] != 0 {
			return memo[st] == 1
		}

		if st == len(s) {
			return true
		}

		defer func() {
			if ret {
				memo[st] = 1
			} else {
				memo[st] = 2
			}
		}()

		for _, w := range wordDict {
			if len(s[st:]) >= len(w) && w == s[st:st+len(w)] {
				if dfs(st + len(w)) {
					return true
				}
			}
		}

		return false
	}

	return dfs(0)
}

func wordBreak3(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if len(s[i:]) >= len(w) && w == s[i:i+len(w)] {
				if !dp[i] {
					dp[i] = dp[i+len(w)]
				}
			}
		}
	}
	return dp[0]
}

func wordBreak4(s string, wordDict []string) bool {
	words := make(map[string]bool)
	queue := make([]int, 0)
	seen := make([]bool, len(s)+1) // 重复的子字串会导致 queue 被塞爆 golang/dynamicprogramming/139. Word Break.go:236
	for _, word := range wordDict {
		words[word] = true
	}
	queue = append(queue, 0)
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if start == len(s) {
			return true
		}
		for end := start + 1; end <= len(s); end++ {
			if seen[end] {
				continue
			}
			if _, ok := words[s[start:end]]; ok {
				queue = append(queue, end)
				seen[end] = true
			}
		}
	}
	return false
}

func Test_wordBreak() {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	ans1 := wordBreak(s1, wordDict1)
	log.Printf("ans1: %v", ans1)

	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	ans2 := wordBreak(s2, wordDict2)
	log.Printf("ans2: %v", ans2)

	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "og", "and", "cat"}
	ans3 := wordBreak(s3, wordDict3)
	log.Printf("ans3: %v", ans3)

	s4 := "aaaaaaa"
	wordDict4 := []string{"aaaa", "aaa"}
	ans4 := wordBreak(s4, wordDict4)
	log.Printf("ans4: %v", ans4)
}

func Test_wordBreak1() {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	ans1 := wordBreak1(s1, wordDict1)
	log.Printf("ans1: %v", ans1)

	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	ans2 := wordBreak1(s2, wordDict2)
	log.Printf("ans2: %v", ans2)

	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "og", "and", "cat"}
	ans3 := wordBreak1(s3, wordDict3)
	log.Printf("ans3: %v", ans3)

	s4 := "aaaaaaa"
	wordDict4 := []string{"aaaa", "aaa"}
	ans4 := wordBreak1(s4, wordDict4)
	log.Printf("ans4: %v", ans4)
}

func Test_wordBreak2() {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	ans1 := wordBreak2(s1, wordDict1)
	log.Printf("ans1: %v", ans1)

	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	ans2 := wordBreak2(s2, wordDict2)
	log.Printf("ans2: %v", ans2)

	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "og", "and", "cat"}
	ans3 := wordBreak2(s3, wordDict3)
	log.Printf("ans3: %v", ans3)

	s4 := "aaaaaaa"
	wordDict4 := []string{"aaaa", "aaa"}
	ans4 := wordBreak2(s4, wordDict4)
	log.Printf("ans4: %v", ans4)
}

func Test_wordBreak3() {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	ans1 := wordBreak3(s1, wordDict1)
	log.Printf("ans1: %v", ans1)

	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	ans2 := wordBreak3(s2, wordDict2)
	log.Printf("ans2: %v", ans2)

	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "og", "and", "cat"}
	ans3 := wordBreak3(s3, wordDict3)
	log.Printf("ans3: %v", ans3)

	s4 := "aaaaaaa"
	wordDict4 := []string{"aaaa", "aaa"}
	ans4 := wordBreak3(s4, wordDict4)
	log.Printf("ans4: %v", ans4)
}

func Test_wordBreak4() {
	//s1 := "leetcode"
	//wordDict1 := []string{"leet", "code"}
	//ans1 := wordBreak4(s1, wordDict1)
	//log.Printf("ans1: %v", ans1)
	//
	//s2 := "applepenapple"
	//wordDict2 := []string{"apple", "pen"}
	//ans2 := wordBreak4(s2, wordDict2)
	//log.Printf("ans2: %v", ans2)
	//
	//s3 := "catsandog"
	//wordDict3 := []string{"cats", "dog", "sand", "og", "and", "cat"}
	//ans3 := wordBreak4(s3, wordDict3)
	//log.Printf("ans3: %v", ans3)
	//
	//s4 := "aaaaaaa"
	//wordDict4 := []string{"aaaa", "aaa"}
	//ans4 := wordBreak4(s4, wordDict4)
	//log.Printf("ans4: %v", ans4)

	s5 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict5 := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	ans5 := wordBreak4(s5, wordDict5)
	log.Printf("ans5: %v", ans5)
}
