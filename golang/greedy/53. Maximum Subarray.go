package greedy

import (
	"log"
	"math"
)

// TLE error
func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sum := 0
	maxSum := math.MinInt64
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= len(nums)-i; j++ {
			for k := j; k < j+i; k++ {
				sum += nums[k]
			}
			if sum > maxSum {
				maxSum = sum
			}
			sum = 0
		}
	}

	return maxSum
}

// kadane's algorithm, TC:O(n), SC:O(1)
func MaxSubArray2(nums []int) int {
	curSum, maxSum := 0, nums[0]

	for i := 0; i < len(nums); i++ {
		curSum = max(curSum+nums[i], nums[i])
		maxSum = max(maxSum, curSum)
	}

	return maxSum
}

func MaxSubArray3(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

func MaxSubArray4(nums []int) int {
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return math.MinInt64
		}

		mid := l + (r-l)/2
		curr := 0
		leftSum := 0
		for i := mid - 1; i >= l; i-- {
			curr += nums[i]
			if leftSum < curr {
				leftSum = curr
			}
		}

		curr = 0
		rightSum := 0
		for i := mid + 1; i <= r; i++ {
			curr += nums[i]
			if rightSum < curr {
				rightSum = curr
			}
		}

		combineSum := nums[mid] + leftSum + rightSum

		leftSum = dfs(l, mid-1)
		rightSum = dfs(mid+1, r)

		return max(combineSum, leftSum, rightSum)
	}

	return dfs(0, len(nums)-1)
}

func Test_MaxSubArray() {
	testcase1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans1 := MaxSubArray(testcase1)
	log.Println(ans1)

	testcase2 := []int{1}
	ans2 := MaxSubArray(testcase2)
	log.Println(ans2)

	testcase3 := []int{5, 4, -1, 7, 8}
	ans3 := MaxSubArray(testcase3)
	log.Println(ans3)
}

func Test_MaxSubArray2() {
	testcase1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans1 := MaxSubArray2(testcase1)
	log.Println(ans1)

	testcase2 := []int{1}
	ans2 := MaxSubArray2(testcase2)
	log.Println(ans2)

	testcase3 := []int{5, 4, -1, 7, 8}
	ans3 := MaxSubArray2(testcase3)
	log.Println(ans3)
}

func Test_MaxSubArray3() {
	testcase1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans1 := MaxSubArray3(testcase1)
	log.Println(ans1)

	testcase2 := []int{1}
	ans2 := MaxSubArray3(testcase2)
	log.Println(ans2)

	testcase3 := []int{5, 4, -1, 7, 8}
	ans3 := MaxSubArray3(testcase3)
	log.Println(ans3)
}

func Test_MaxSubArray4() {
	testcase1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans1 := MaxSubArray4(testcase1)
	log.Println(ans1)

	testcase2 := []int{1}
	ans2 := MaxSubArray4(testcase2)
	log.Println(ans2)

	testcase3 := []int{5, 4, -1, 7, 8}
	ans3 := MaxSubArray4(testcase3)
	log.Println(ans3)
}
