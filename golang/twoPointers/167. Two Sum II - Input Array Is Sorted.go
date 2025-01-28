package twoPointers

import (
	"log"
)

// binary search solution, O(n * ㏒N)
func twoSum(numbers []int, target int) []int {
	// exactly one solution, this is for edge case
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	for i := 0; i < len(numbers); i++ {
		if idx := binarySearch(numbers, target-numbers[i]); idx != -1 && idx != i {
			if idx > i {
				return []int{i + 1, idx + 1}
			} else {
				return []int{idx + 1, i + 1}
			}
		}
	}
	return []int{} // unreachable
}

func binarySearch(nums []int, target int) int {
	var left, right int

	left = 0
	right = len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// two pointers solution, O(N)
func twoSum2(numbers []int, target int) []int {
	//exactly one solution, this is for edge case
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	var left, right int

	left = 0
	right = len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{} //unreachable
}

func twoSumReview(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]

		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{} // unreachable
}

func Test_TwoSum_II_Input_Array_Is_Sorted() {
	var case1 = []int{2, 7, 11, 15}
	var target1 = 9
	ans1 := twoSum(case1, target1)
	log.Printf("ans1: %v", ans1)

	var case2 = []int{2, 3, 4}
	var target2 = 6
	ans2 := twoSum(case2, target2)
	log.Printf("ans2: %v", ans2)

	var case3 = []int{-1, 0}
	var target3 = -1
	ans3 := twoSum(case3, target3)
	log.Printf("ans3: %v", ans3)

	var case4 = []int{1, 2, 3, 4, 4, 9, 56, 90}
	var target4 = 8
	ans4 := twoSum(case4, target4)
	log.Printf("ans4: %v", ans4)
}

func Test_TwoSum_2_II_Input_Array_Is_Sorted() {
	var case1 = []int{2, 7, 11, 15}
	var target1 = 9
	ans1 := twoSum2(case1, target1)
	log.Printf("ans1: %v", ans1)

	var case2 = []int{2, 3, 4}
	var target2 = 6
	ans2 := twoSum2(case2, target2)
	log.Printf("ans2: %v", ans2)

	var case3 = []int{-1, 0}
	var target3 = -1
	ans3 := twoSum2(case3, target3)
	log.Printf("ans3: %v", ans3)

	var case4 = []int{1, 2, 3, 4, 4, 9, 56, 90}
	var target4 = 8
	ans4 := twoSum2(case4, target4)
	log.Printf("ans4: %v", ans4)
}

func Test_TwoSumReview() {
	var case1 = []int{2, 7, 11, 15}
	var target1 = 9
	ans1 := twoSumReview(case1, target1)
	log.Printf("ans1: %v", ans1)

	var case2 = []int{2, 3, 4}
	var target2 = 6
	ans2 := twoSumReview(case2, target2)
	log.Printf("ans2: %v", ans2)

	var case3 = []int{-1, 0}
	var target3 = -1
	ans3 := twoSumReview(case3, target3)
	log.Printf("ans3: %v", ans3)

	var case4 = []int{1, 2, 3, 4, 4, 9, 56, 90}
	var target4 = 8
	ans4 := twoSumReview(case4, target4)
	log.Printf("ans4: %v", ans4)
}
