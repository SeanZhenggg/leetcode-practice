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
