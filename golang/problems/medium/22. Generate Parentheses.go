package medium

import (
	"log"
	"strings"
)

func generateParenthesis(n int) []string {
	var ret []string
	m := make(map[string]int)
	var dfs func(str string, countMap map[string]int)
	dfs = func(str string, countMap map[string]int) {
		if countMap["("] > n || countMap[")"] > n {
			return
		}
		if countMap["("] == n && countMap[")"] == n {
			ret = append(ret, str)
			return
		}
		if countMap[")"] > countMap["("] {
			return
		}

		countMap["("]++
		dfs(str+"(", countMap)
		countMap["("]--

		countMap[")"]++
		dfs(str+")", countMap)
		countMap[")"]--
	}

	dfs("", m)

	return ret
}

func generateParenthesis2(n int) []string {
	st := make([]string, 0, 2*n)
	ret := make([]string, 0)
	var dfs func(openN int, closeN int)

	dfs = func(openN int, closeN int) {
		if openN == n && closeN == n {
			ret = append(ret, strings.Join(st, ""))
			return
		}

		if openN < n {
			st = append(st, "(")
			dfs(openN+1, closeN)
			st = st[:len(st)-1]
		}

		if closeN < openN {
			st = append(st, ")")
			dfs(openN, closeN+1)
			st = st[:len(st)-1]
		}
	}

	dfs(0, 0)

	return ret
}

func Test_GenerateParenthesis() {
	ans1 := generateParenthesis(3)
	log.Printf("ans1: %v", ans1)
	ans2 := generateParenthesis(1)
	log.Printf("ans2: %v", ans2)
}

func Test_GenerateParenthesis2() {
	ans1 := generateParenthesis2(3)
	log.Printf("ans1: %v", ans1)
	ans2 := generateParenthesis2(1)
	log.Printf("ans2: %v", ans2)
}
