package medium

import "log"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if target > nums[mid] {
			if target > nums[r] && nums[mid] <= nums[r] {
				r = mid - 1
			} else {
				// t <= r or m > r
				l = mid + 1
			}
		} else if target == nums[mid] {
			return mid
		} else {
			if target < nums[l] && nums[mid] >= nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func Test_search() {
	nums1 := []int{2, 3, 4, 5, 1}
	t1 := 1
	ans1 := search(nums1, t1)
	log.Printf("ans1: %v", ans1)
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	t2 := 3
	ans2 := search(nums2, t2)
	log.Printf("ans2: %v", ans2)
	nums3 := []int{1}
	t3 := 0
	ans3 := search(nums3, t3)
	log.Printf("ans3: %v", ans3)
	nums4 := []int{5, 6, 7, 0, 1, 2, 4}
	t4 := 7
	ans4 := search(nums4, t4)
	log.Printf("ans4: %v", ans4)
	nums5 := []int{1, 2, 4, 5, 6, 7, 0}
	t5 := 6
	ans5 := search(nums5, t5)
	log.Printf("ans5: %v", ans5)
}
