package dynamicprogramming

import "log"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = -1
	}

	var dfs func(idx int) int
	dfs = func(idx int) (num int) {
		if idx < 0 {
			return 0
		}

		if dp[idx] != -1 {
			return dp[idx]
		}

		count := 1
		for i := idx - 1; i >= 0; i-- {
			if nums[idx] > nums[i] {
				count = max(count, 1+dfs(i))
			}
		}

		defer func() {
			dp[idx] = num
		}()
		return count
	}
	c := 1

	maxVal := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= maxVal {
			c = max(c, dfs(i))
			maxVal = nums[i]
		}
	}
	return c
}

func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	var maxVal int = dp[0]
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], 1+dp[j])

				maxVal = max(maxVal, dp[i])
			}
		}
	}

	return maxVal
}

func Test_lengthOfLIS() {
	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	ans1 := lengthOfLIS(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0, 1, 0, 3, 2, 3}
	ans2 := lengthOfLIS(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{7, 7, 7, 7, 7, 7, 7}
	ans3 := lengthOfLIS(nums3)
	log.Printf("ans3: %v", ans3)

	nums4 := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	ans4 := lengthOfLIS(nums4)
	log.Printf("ans4: %v", ans4)

}

func Test_lengthOfLIS2() {
	nums1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	ans1 := lengthOfLIS2(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0, 1, 0, 3, 2, 3}
	ans2 := lengthOfLIS2(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{7, 7, 7, 7, 7, 7, 7}
	ans3 := lengthOfLIS2(nums3)
	log.Printf("ans3: %v", ans3)

	nums4 := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	ans4 := lengthOfLIS2(nums4)
	log.Printf("ans4: %v", ans4)

}
