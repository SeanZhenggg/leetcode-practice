package dynamicprogramming

import (
	"log"
	"math"
	"slices"
)

// brute-force, TLE error, TC: O(n^3), SC: O(1)
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxVal := math.MinInt32 - 1
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= len(nums)-i; j++ {
			val := 1
			for k := j; k < j+i; k++ {
				val *= nums[k]
			}
			maxVal = max(maxVal, val)
		}
	}

	return maxVal
}

// maxf(n) = max(maxf(n-1) * nums[n], minf(n-1) * nums[n], nums[n])
func maxProduct2(nums []int) int {
	maxF := make([]int, len(nums))
	minF := make([]int, len(nums))
	maxF[0], minF[0] = nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], minF[i-1]*nums[i], nums[i])
		minF[i] = min(maxF[i-1]*nums[i], minF[i-1]*nums[i], nums[i])

	}

	return slices.Max(maxF)
}

func maxProduct3(nums []int) int {
	maxF, minF := nums[0], nums[0]
	res := maxF

	for i := 1; i < len(nums); i++ {
		newMaxF := max(maxF*nums[i], minF*nums[i], nums[i])
		newMinF := min(maxF*nums[i], minF*nums[i], nums[i])
		maxF = newMaxF
		minF = newMinF
		res = max(res, maxF)
	}

	return res
}

func Test_maxProduct() {
	case1 := []int{2, 3, -2, 4}
	ans1 := maxProduct(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{-2, 0, 1}
	ans2 := maxProduct(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{-2}
	ans3 := maxProduct(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-4, -3}
	ans4 := maxProduct(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_maxProduct2() {
	case1 := []int{2, 3, -2, 4}
	ans1 := maxProduct2(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{-2, 0, 1}
	ans2 := maxProduct2(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{-2}
	ans3 := maxProduct2(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-4, -3}
	ans4 := maxProduct2(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_maxProduct3() {
	case1 := []int{2, 3, -2, 4}
	ans1 := maxProduct3(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{-2, 0, 1}
	ans2 := maxProduct3(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{-2}
	ans3 := maxProduct3(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-4, -3}
	ans4 := maxProduct3(case4)
	log.Printf("ans4: %v", ans4)
}
