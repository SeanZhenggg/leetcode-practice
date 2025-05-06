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

// dp solution, TC: O(n), SC: O(n)
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

// dp solution w/ space optimization, TC: O(n), SC: O(1)
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

// prefix/suffix solution, TC: O(n), SC: O(1)
func maxProduct4(nums []int) int {
	prefix, suffix := 1, 1
	res := nums[0]

	for i := 0; i < len(nums); i++ {
		if prefix == 0 { // IMPORTANT!!! e.g. nums = [-1, -2, -3, 0] or [-3, 0, 1, -2]
			prefix = 1
		}
		if suffix == 0 { // IMPORTANT!!! e.g. nums = [-1, -2, -3, 0] or [-3, 0, 1, -2]
			suffix = 1
		}

		prefix *= nums[i]
		suffix *= nums[len(nums)-1-i]
		res = max(res, prefix, suffix)
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

	case5 := []int{-3, 0, 1, -2}
	ans5 := maxProduct(case5)
	log.Printf("ans5: %v", ans5)

	case6 := []int{-1, -2, -3, 0}
	ans6 := maxProduct(case6)
	log.Printf("ans6: %v", ans6)
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

	case5 := []int{-3, 0, 1, -2}
	ans5 := maxProduct2(case5)
	log.Printf("ans5: %v", ans5)

	case6 := []int{-1, -2, -3, 0}
	ans6 := maxProduct2(case6)
	log.Printf("ans6: %v", ans6)
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

	case5 := []int{-3, 0, 1, -2}
	ans5 := maxProduct3(case5)
	log.Printf("ans5: %v", ans5)

	case6 := []int{-1, -2, -3, 0}
	ans6 := maxProduct3(case6)
	log.Printf("ans6: %v", ans6)
}

func Test_maxProduct4() {
	case1 := []int{2, 3, -2, 4}
	ans1 := maxProduct4(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{-2, 0, 1}
	ans2 := maxProduct4(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{-2}
	ans3 := maxProduct4(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{-4, -3}
	ans4 := maxProduct4(case4)
	log.Printf("ans4: %v", ans4)

	case5 := []int{-3, 0, 1, -2}
	ans5 := maxProduct4(case5)
	log.Printf("ans5: %v", ans5)

	case6 := []int{-1, -2, -3, 0}
	ans6 := maxProduct4(case6)
	log.Printf("ans6: %v", ans6)
}
