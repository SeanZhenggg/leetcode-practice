package binarySearch

import "log"

// one pass solution, TC: O(logn), SC: O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if target > nums[mid] {
			if target > nums[r] && nums[mid] <= nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
			//if target <= nums[r] {
			//	l = mid + 1
			//} else {
			//	if nums[mid] <= nums[r] {
			//		r = mid - 1
			//	} else {
			//		l = mid + 1
			//	}
			//}
		} else if target < nums[mid] {
			if target < nums[l] && nums[mid] >= nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
			//if target >= nums[l] {
			//	r = mid - 1
			//} else {
			//	if nums[mid] >= nums[l] {
			//		l = mid + 1
			//	} else {
			//		r = mid - 1
			//	}
			//}
		} else {
			return mid
		}
	}
	return -1
}

// find pivot and do bs both sides solution, TC: O(3*logn), SC: O(1)
func search2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] <= nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	var bs func(l, r int, target int) int
	bs = func(l, r int, target int) int {
		for l <= r {
			m := l + (r-l)/2

			if nums[m] == target {
				return m
			} else if nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return -1
	}

	idx := bs(0, l-1, target)
	if idx != -1 {
		return idx
	}

	return bs(l, len(nums)-1, target)
}

// two pass solution, TC: O(2*logn), SC: O(1)
func search3(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] <= nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	p := l
	l, r = 0, len(nums)-1
	if target >= nums[p] && target <= nums[r] {
		l = p
	} else {
		r = p - 1
	}
	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
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
