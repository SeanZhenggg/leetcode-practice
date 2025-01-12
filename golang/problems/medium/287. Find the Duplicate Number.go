package medium

import "log"

// 2 loop solution - O(n^2), O(1)
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				return nums[j]
			}
		}
	}
	return 0 //unreachable
}

// map solution - O(n), O(n)
func findDuplicate2(nums []int) int {
	m := make(map[int]struct{}, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			return nums[i]
		}

		m[nums[i]] = struct{}{}
	}
	return 0 //unreachable
}

// binary search solution - O(n * log(n)), O(1)
func findDuplicate3(nums []int) int {
	l, r := 1, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		lteCount := 0
		for _, num := range nums {
			if num <= mid {
				lteCount += 1
			}
		}

		if lteCount <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findDuplicate4(nums []int) int {
	fast := 0
	slow := 0
	fast = nums[nums[fast]]
	slow = nums[slow]

	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return slow
}

func Test_findDuplicate() {
	nums1 := []int{1, 3, 4, 2, 2}
	ans1 := findDuplicate(nums1)
	log.Printf("ans1: %v", ans1)
	nums2 := []int{3, 1, 3, 4, 2}
	ans2 := findDuplicate(nums2)
	log.Printf("ans2: %v", ans2)
	nums3 := []int{3, 3, 3, 3, 3}
	ans3 := findDuplicate(nums3)
	log.Printf("ans3: %v", ans3)
}

func Test_findDuplicate2() {
	nums1 := []int{1, 3, 4, 2, 2}
	ans1 := findDuplicate2(nums1)
	log.Printf("ans1: %v", ans1)
	nums2 := []int{3, 1, 3, 4, 2}
	ans2 := findDuplicate2(nums2)
	log.Printf("ans2: %v", ans2)
	nums3 := []int{3, 3, 3, 3, 3}
	ans3 := findDuplicate2(nums3)
	log.Printf("ans3: %v", ans3)
}

func Test_findDuplicate3() {
	nums1 := []int{1, 3, 4, 2, 2}
	ans1 := findDuplicate3(nums1)
	log.Printf("ans1: %v", ans1)
	nums2 := []int{3, 1, 3, 4, 2}
	ans2 := findDuplicate3(nums2)
	log.Printf("ans2: %v", ans2)
	nums3 := []int{3, 3, 3, 3, 3}
	ans3 := findDuplicate3(nums3)
	log.Printf("ans3: %v", ans3)
}

func Test_findDuplicate4() {
	nums1 := []int{1, 3, 4, 2, 2}
	ans1 := findDuplicate4(nums1)
	log.Printf("ans1: %v", ans1)
	nums2 := []int{3, 1, 3, 4, 2}
	ans2 := findDuplicate4(nums2)
	log.Printf("ans2: %v", ans2)
	nums3 := []int{3, 3, 3, 3, 3}
	ans3 := findDuplicate4(nums3)
	log.Printf("ans3: %v", ans3)
}
