package slidingWindow

import "log"

// solution using prefixSum

// [7,1,5,3,6,4]
// prefixMin => 7 => 1 => 1 => 1 => 1 => 1
// maxProfit => 0 => 0 => 4 => 4 => 5 => 5

// [7, 6, 4, 3, 1]
// prefixMin => 7 => 6 => 4 => 3 => 1
// maxProfit => 0 => 0 => 0 => 0 => 0

func maxProfit(prices []int) int {
	maxProfits := 0
	prefixMin := prices[0]
	for i := 0; i < len(prices); i++ {
		if prefixMin > prices[i] {
			prefixMin = prices[i]
		}

		profit := prices[i] - prefixMin
		if profit > maxProfits {
			maxProfits = profit
		}
	}

	return maxProfits
}

//solution using dynamic sliding window

// 7    1    5    3    6    4
// l
// r
// r -> r
// l -> l
//      r -> r
//   		 r -> r
// 				  r -> r
// 					   r -> r

// 7    6    4    3    1
// l
// r
// r -> r
// l -> l
// 		r -> r
// 		l -> l
// 			 r -> r
// 			 l -> l
// 				  r -> r
// 				  l -> l

func maxProfit2(prices []int) int {
	l := 0
	maxProfits := 0
	for r := 0; r < len(prices); r++ {
		for prices[l] > prices[r] {
			l++
		}

		profit := prices[r] - prices[l]
		if profit > maxProfits {
			maxProfits = profit
		}
	}

	return maxProfits
}

func maxProfitReview(prices []int) int {
	l := 0
	maxP := 0
	for r := 0; r < len(prices); r++ {
		if prices[l] > prices[r] {
			l = r
		}

		maxP = max(maxP, prices[r]-prices[l])
	}

	return maxP
}

func maxProfitReview2(prices []int) int {
	dp := make([][2]int, 0, len(prices))

	dp = append(dp, [2]int{-prices[0], 0})
	for i := 1; i < len(prices); i++ {
		dp = append(dp, [2]int{max(dp[i-1][0], -prices[i]), max(prices[i]+dp[i-1][0], dp[i-1][1])})
	}

	return dp[len(prices)-1][1]
}

func Test_MaxProfit() {
	case1 := []int{7, 1, 5, 3, 6, 4}
	ans1 := maxProfit(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{7, 6, 4, 3, 1}
	ans2 := maxProfit(case2)
	log.Printf("ans2: %v", ans2)
}

func Test_MaxProfitReview() {
	case1 := []int{7, 1, 5, 3, 6, 4}
	ans1 := maxProfitReview2(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{7, 6, 4, 3, 1}
	ans2 := maxProfitReview2(case2)
	log.Printf("ans2: %v", ans2)
}
