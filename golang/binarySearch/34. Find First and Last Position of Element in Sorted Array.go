package binarySearch

import "log"

func searchRange(nums []int, target int) []int {
	notFound := []int{-1, -1}
	if len(nums) == 0 {
		return notFound
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l > len(nums)-1 || l < 0 {
		return notFound
	}
	if nums[l] != target {
		return notFound
	}

	ret := make([]int, 2)
	ret[0] = l

	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	ret[1] = r

	return ret
}

func searchRangeReview(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] >= target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}

	if l < 0 || l > len(nums)-1 || nums[l] != target {
		return []int{-1, -1}
	} else {
		j := l
		for i := l; i < len(nums); i++ {
			if nums[i] == target {
				j = i
			}
		}
		return []int{l, j}
	}
}

func Test_searchRange() {
	case1 := []int{5, 7, 7, 8, 8, 10}
	t1 := 8
	ans1 := searchRange(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{5, 7, 7, 8, 8, 10}
	t2 := 6
	ans2 := searchRange(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{}
	t3 := 0
	ans3 := searchRange(case3, t3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{2, 2}
	t4 := 3
	ans4 := searchRange(case4, t4)
	log.Printf("ans4: %v", ans4)
	case5 := []int{3, 3}
	t5 := 2
	ans5 := searchRange(case5, t5)
	log.Printf("ans5: %v", ans5)
}

func Test_searchRangeReview() {
	case1 := []int{5, 7, 7, 8, 8, 10}
	t1 := 8
	ans1 := searchRangeReview(case1, t1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{5, 7, 7, 8, 8, 10}
	t2 := 6
	ans2 := searchRangeReview(case2, t2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{}
	t3 := 0
	ans3 := searchRangeReview(case3, t3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{2, 2}
	t4 := 3
	ans4 := searchRangeReview(case4, t4)
	log.Printf("ans4: %v", ans4)
	case5 := []int{3, 3}
	t5 := 2
	ans5 := searchRangeReview(case5, t5)
	log.Printf("ans5: %v", ans5)
}
