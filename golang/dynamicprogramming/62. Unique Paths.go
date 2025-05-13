package dynamicprogramming

import "log"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[1][0] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	left := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[j] = dp[j] + left
			left = dp[j]
		}

		left = 0
	}
	return dp[n]
}

func uniquePaths3(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

func Test_uniquePaths() {
	m1, n1 := 3, 7
	ans1 := uniquePaths(m1, n1)
	log.Printf("ans1: %v", ans1)

	m2, n2 := 3, 2
	ans2 := uniquePaths(m2, n2)
	log.Printf("ans2: %v", ans2)

}

func Test_uniquePaths1() {
	m1, n1 := 3, 7
	ans1 := uniquePaths1(m1, n1)
	log.Printf("ans1: %v", ans1)

	m2, n2 := 3, 2
	ans2 := uniquePaths1(m2, n2)
	log.Printf("ans2: %v", ans2)

}

func Test_uniquePaths2() {
	m1, n1 := 3, 7
	ans1 := uniquePaths2(m1, n1)
	log.Printf("ans1: %v", ans1)

	m2, n2 := 3, 2
	ans2 := uniquePaths2(m2, n2)
	log.Printf("ans2: %v", ans2)

}

func Test_uniquePaths3() {
	m1, n1 := 3, 7
	ans1 := uniquePaths3(m1, n1)
	log.Printf("ans1: %v", ans1)

	m2, n2 := 3, 2
	ans2 := uniquePaths3(m2, n2)
	log.Printf("ans2: %v", ans2)

}
