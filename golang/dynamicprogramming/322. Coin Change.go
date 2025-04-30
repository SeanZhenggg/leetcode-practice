package dynamicprogramming

import (
	"log"
	"math"
	"sort"
)

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
				total += coins[i]
				count++
				dfs(count, total)
				total -= coins[i]
				count--
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

func coinChange2(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return i > j
	})

	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt64
	}

	dp[0] = 0

	for j := 1; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if coins[i] == j {
				dp[j] = 1
				continue
			} else if coins[i] < j {
				if dp[j-coins[i]] != math.MaxInt64 {
					dp[j] = min(1+dp[j-coins[i]], dp[j])
				}
			}
		}
	}

	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func coinChange21(coins []int, amount int) int {
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

	if dp[amount] >= math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func Test_coinChange() {
	coins1 := []int{1, 2, 5}
	amount1 := 11
	ans1 := coinChange(coins1, amount1)
	log.Printf("ans1: %v", ans1)
}

func Test_coinChange2() {
	//coins1 := []int{1, 2, 5}
	//amount1 := 11
	//ans1 := coinChange2(coins1, amount1)
	//log.Printf("ans1: %v", ans1)

	coins2 := []int{2}
	amount2 := 3
	ans2 := coinChange2(coins2, amount2)
	log.Printf("ans2: %v", ans2)

	coins3 := []int{1}
	amount3 := 0
	ans3 := coinChange2(coins3, amount3)
	log.Printf("ans3: %v", ans3)

}
