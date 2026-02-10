package slidingWindow

// solution using prefix min

// [7,1,5,3,6,4]
// prefixMin => 7 => 1 => 1 => 1 => 1 => 1
// maxProfit => 0 => 0 => 4 => 4 => 5 => 5

// [7, 6, 4, 3, 1]
// prefixMin => 7 => 6 => 4 => 3 => 1
// maxProfit => 0 => 0 => 0 => 0 => 0

func maxProfitPrefixMin(prices []int) int {
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

func maxProfitDynamicSlidingWindow(prices []int) int {
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

//solution using dynamic programming

// f(n) = total 的最大獲利
// dp[i][0] 表示第 i 天持有股票所得的最多現金
// dp[i][1] 表示第 i 天不持有股票所得的最多現金
// 持有 != 買入, 持有可能是前幾天就持有, 今天保持持有的狀態
// dp[i][0] = max(dp[i-1][0], -prices[i]) (前一天就持有 vs. 當天買入持有)
// dp[i][1] = max(dp[i-1][1], prices[i] + dp[i-1][0]) (前一天就不持有 vs. 當天賣出股票不持有)
// (Q: 為什麼是 prices[i] + dp[i-1][0] ? A: 因為持有的時候是直接 -prices[i], 賣出抵銷的時候要+回來)

func maxProfitReviewDP(prices []int) int {
	dp := make([][2]int, len(prices))

	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	}

	return dp[len(prices)-1][1]
}

func maxProfitKadane(prices []int) int {
	maxProfit := 0
	minBuy := prices[0]

	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-minBuy)
		minBuy = min(minBuy, prices[i])
	}

	return maxProfit
}
