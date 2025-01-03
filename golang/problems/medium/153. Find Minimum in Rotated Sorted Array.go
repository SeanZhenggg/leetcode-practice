package medium

import (
	"log"
	"math"
)

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return nums[l]
	}

	search := math.MinInt16
	minVal := math.MaxInt16
	for l <= r {
		mid := l + (r-l)/2

		isRotated := false
		if (nums[l] < nums[mid] && nums[mid] < nums[r] && nums[l] < nums[r]) || (nums[mid] < nums[l] && nums[mid] < nums[r] && nums[l] > nums[r]) {
			isRotated = false
		} else {
			isRotated = true
		}

		if minVal > nums[mid] {
			minVal = nums[mid]
		}

		if nums[mid] <= search {
			if isRotated {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if isRotated {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return minVal
}

func Test_findMin() {
	case1 := []int{3, 4, 5, 1, 2}
	ans1 := findMin(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 5, 6, 7, 0, 1, 2}
	ans2 := findMin(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{11, 13, 15, 17}
	ans3 := findMin(case3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{3, 1, 2}
	ans4 := findMin(case4)
	log.Printf("ans4: %v", ans4)
	case5 := []int{5, 1, 2, 3, 4}
	ans5 := findMin(case5)
	log.Printf("ans5: %v", ans5)

}
