package binarySearch

import "log"

func searchInsert(nums []int, target int) int {
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

	return l
}

func Test_searchInsert() {
	case1 := []int{1, 3, 5, 6}
	t1 := 5
	ans1 := searchInsert(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{1, 3, 5, 6}
	t2 := 2
	ans2 := searchInsert(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{1, 3, 5, 6}
	t3 := 7
	ans3 := searchInsert(case3, t3)
	log.Printf("ans3: %v", ans3)
}
