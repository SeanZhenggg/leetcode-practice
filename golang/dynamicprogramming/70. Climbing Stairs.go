package dynamicprogramming

import "log"

// recursion solution - TLE error
func climbStairs(n int) int {
	var dfs func(step int) int
	dfs = func(n int) int {
		// 第 0 步的可能為 1
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}

		return dfs(n-1) + dfs(n-2) // 走 1 步到 n 的所有組合 + 走 2 步到 n 的所有組合
	}

	return dfs(n)
}

// recursion solution w/ memorization - TC: O(n), SC: O(n)
func climbStairs2(n int) int {
	visitedN := make([]int, n+1)
	for i := 0; i < len(visitedN); i++ {
		visitedN[i] = -1
	}

	var dfs func(step int) int
	dfs = func(n int) (sum int) {
		// 第 0 步的可能為 1
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}

		if visitedN[n] != -1 {
			return visitedN[n]
		}

		defer func() {
			visitedN[n] = sum
		}()

		return dfs(n-1) + dfs(n-2)
	}

	return dfs(n)
}

// dp solution - TC: O(n), SC: O(n)
func climbStairs3(n int) int {
	dp := make([]int, n+1)

	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// dp solution w/ space optimization - TC: O(n), SC: O(1)
func climbStairs4(n int) int {
	f0, f1 := 1, 1
	newF := 0
	for i := 2; i <= n; i++ {
		newF = f1 + f0
		f0 = f1
		f1 = newF
	}

	return f1
}

func Test_climbStairs() {
	n1 := 2
	ans1 := climbStairs(n1)
	log.Printf("ans1: %v", ans1)

	n2 := 3
	ans2 := climbStairs(n2)
	log.Printf("ans2: %v", ans2)

	n3 := 5
	ans3 := climbStairs(n3)
	log.Printf("ans3: %v", ans3)
}

func Test_climbStairs2() {
	n1 := 2
	ans1 := climbStairs2(n1)
	log.Printf("ans1: %v", ans1)

	n2 := 3
	ans2 := climbStairs2(n2)
	log.Printf("ans2: %v", ans2)

	n3 := 5
	ans3 := climbStairs2(n3)
	log.Printf("ans3: %v", ans3)
}

func Test_climbStairs3() {
	n1 := 2
	ans1 := climbStairs3(n1)
	log.Printf("ans1: %v", ans1)

	n2 := 3
	ans2 := climbStairs3(n2)
	log.Printf("ans2: %v", ans2)

	n3 := 5
	ans3 := climbStairs3(n3)
	log.Printf("ans3: %v", ans3)
}

func Test_climbStairs4() {
	n1 := 2
	ans1 := climbStairs4(n1)
	log.Printf("ans1: %v", ans1)

	n2 := 3
	ans2 := climbStairs4(n2)
	log.Printf("ans2: %v", ans2)

	n3 := 5
	ans3 := climbStairs4(n3)
	log.Printf("ans3: %v", ans3)
}
