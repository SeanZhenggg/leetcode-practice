package binarySearch

import "log"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if target > nums[mid] {
			l = mid + 1
		} else if target == nums[mid] {
			return mid
		} else {
			r = mid - 1
		}
	}

	return -1
}

func Test_search() {
	case1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	ans1 := search(case1, target1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	ans2 := search(case2, target2)
	log.Printf("ans2: %v", ans2)
}
