package binarySearch

import (
	"log"
	"math"
)

// first solution
func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return nums[l]
	}

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

		if isRotated {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return minVal
}

func findMinReview(nums []int) int {
	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return nums[l]
	}

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

func findMin2Review(nums []int) int {
	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return nums[l]
	}

	for l < r {
		if nums[l] < nums[r] {
			return nums[l]
		}

		mid := l + (r-l)/2
		if nums[mid] < nums[l] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
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
	case6 := []int{2, 3, 4, 5, 1}
	ans6 := findMin(case6)
	log.Printf("ans6: %v", ans6)
}

func Test_findMinReview() {
	case1 := []int{3, 4, 5, 1, 2}
	ans1 := findMinReview(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 5, 6, 7, 0, 1, 2}
	ans2 := findMinReview(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{11, 13, 15, 17}
	ans3 := findMinReview(case3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{3, 1, 2}
	ans4 := findMinReview(case4)
	log.Printf("ans4: %v", ans4)
	case5 := []int{5, 1, 2, 3, 4}
	ans5 := findMinReview(case5)
	log.Printf("ans5: %v", ans5)
	case6 := []int{2, 3, 4, 5, 1}
	ans6 := findMinReview(case6)
	log.Printf("ans6: %v", ans6)
}

func Test_findMin2Review() {
	case1 := []int{3, 4, 5, 1, 2}
	ans1 := findMin2Review(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 5, 6, 7, 0, 1, 2}
	ans2 := findMin2Review(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{11, 13, 15, 17}
	ans3 := findMin2Review(case3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{3, 1, 2}
	ans4 := findMin2Review(case4)
	log.Printf("ans4: %v", ans4)
	case5 := []int{5, 1, 2, 3, 4}
	ans5 := findMin2Review(case5)
	log.Printf("ans5: %v", ans5)
	case6 := []int{2, 3, 4, 5, 1}
	ans6 := findMin2Review(case6)
	log.Printf("ans6: %v", ans6)
}
