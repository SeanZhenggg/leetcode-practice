package dynamicprogramming

import (
	"log"
	"math"
	"sort"
)

// backtracking, TLE error, TC: (c ^ n)
func coinChange(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return i > j
	})

	var minVal = -1

	var dfs func(count, total int)
	dfs = func(count, total int) {
		if total > amount {
			return
		}
		if total == amount {
			if minVal == -1 {
				minVal = count
			} else {
				minVal = min(minVal, count)
			}
			return
		}

		for i := 0; i < len(coins); i++ {
			if total+coins[i] <= amount {
				dfs(count, total)
			}
		}
	}

	dfs(0, 0)
	if minVal == math.MaxInt {
		return -1
	} else {
		return minVal
	}
}

// dp top-down, TC: (c * n), SC: (n)
func coinChange1(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		memo[i] = -1
	}

	var dfs func(total int) int
	dfs = func(total int) int {
		if total == 0 {
			return 0
		}

		if memo[total] != -1 {
			return memo[total]
		}

		res := math.MaxInt32 // count of coins
		for i := 0; i < len(coins); i++ {
			if total-coins[i] >= 0 {
				res = min(res, 1+dfs(total-coins[i]))
			}
		}
		memo[total] = res
		return res
	}

	res := dfs(amount)
	if res == math.MaxInt32 {
		return -1
	}

	return res
}

// dp bottom-up, TC: (c * n), SC: (n)
func coinChange2(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return i > j
	})

	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for j := 1; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if coins[i] <= j {
				dp[j] = min(1+dp[j-coins[i]], dp[j])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// dp bottom-up (simplified), TC: (c * n), SC: (n)
func coinChange21(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(1+dp[j-coins[i]], dp[j])
		}
	}

	if dp[amount] >= math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// dp[1] = min(1+dp[1-coins[0]], dp[1]) <-> dp[1] = min(1+dp[1-coins[0]], dp[1])
// dp[2] = min(1+dp[2-coins[0]], dp[2]) <-> dp[1] = min(1+dp[1-coins[1]], dp[1])
// dp[3] = min(1+dp[3-coins[0]], dp[3]) <-> dp[1] = min(1+dp[1-coins[2]], dp[1])
// ...
// dp[11] = min(1+dp[11-coins[0]], dp[11])

// dp[1] = min(1+dp[1-coins[1]], dp[1]) <-> dp[2] = min(1+dp[2-coins[0]], dp[2])
// dp[2] = min(1+dp[2-coins[1]], dp[2]) <-> dp[2] = min(1+dp[2-coins[1]], dp[2])
// dp[3] = min(1+dp[3-coins[1]], dp[3]) <-> dp[2] = min(1+dp[2-coins[2]], dp[2])
// ...
// dp[11] = min(1+dp[11-coins[1]], dp[11])

// dp[1] = min(1+dp[1-coins[2]], dp[1])
// dp[2] = min(1+dp[2-coins[2]], dp[2])
// dp[3] = min(1+dp[3-coins[2]], dp[3])
// ...
// dp[11] = min(1+dp[11-coins[2]], dp[11])

func Test_coinChange() {
	coins1 := []int{1, 2, 5}
	amount1 := 11
	ans1 := coinChange(coins1, amount1)
	log.Printf("ans1: %v", ans1)
}

func Test_coinChange1() {
	coins1 := []int{1, 2, 5}
	amount1 := 50
	ans1 := coinChange1(coins1, amount1)
	log.Printf("ans1: %v", ans1)

	coins2 := []int{2}
	amount2 := 3
	ans2 := coinChange2(coins2, amount2)
	log.Printf("ans2: %v", ans2)

	coins3 := []int{1}
	amount3 := 0
	ans3 := coinChange2(coins3, amount3)
	log.Printf("ans3: %v", ans3)
}

func Test_coinChange2() {
	coins1 := []int{1, 2, 5}
	amount1 := 50
	ans1 := coinChange2(coins1, amount1)
	log.Printf("ans1: %v", ans1)

	coins2 := []int{2}
	amount2 := 3
	ans2 := coinChange2(coins2, amount2)
	log.Printf("ans2: %v", ans2)

	coins3 := []int{1}
	amount3 := 0
	ans3 := coinChange2(coins3, amount3)
	log.Printf("ans3: %v", ans3)

}

func Test_coinChange21() {
	coins1 := []int{1, 2, 5}
	amount1 := 11
	ans1 := coinChange21(coins1, amount1)
	log.Printf("ans1: %v", ans1)

	coins2 := []int{2}
	amount2 := 3
	ans2 := coinChange21(coins2, amount2)
	log.Printf("ans2: %v", ans2)

	coins3 := []int{1}
	amount3 := 0
	ans3 := coinChange21(coins3, amount3)
	log.Printf("ans3: %v", ans3)

}
