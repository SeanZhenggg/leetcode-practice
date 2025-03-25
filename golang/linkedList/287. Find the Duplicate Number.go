package linkedList

import (
	"log"
	"sort"
)

// 2 loop solution - TC: O(n^2), SC: O(1)
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

// sort solution - TC: O(n*logn), SC: O(1)
func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// hashmap/hashset solution - TC: O(n), SC: O(n)
func findDuplicate3(nums []int) int {
	m := make(map[int]struct{}, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			return nums[i]
		}

		m[nums[i]] = struct{}{}
	}
	return 0 //unreachable
}

// binary search solution - TC: O(n * log(n)), SC: O(1)
func findDuplicate4(nums []int) int {
	l, r := 1, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if isValid(nums, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isValid(nums []int, mid int) bool {
	count := 0
	for _, num := range nums {
		if num <= mid {
			count++
		}
	}

	return count > mid
}

// Floyd's Tortoise and Hare (Cycle Detection) solution - TC: O(n), SC: O(1)
func findDuplicate5(nums []int) int {
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

// array as hashmap(recursion) - TC: O(n), SC: O(n)
func findDuplicate6(nums []int) int {
	var dup func(nums []int, cur int) int
	dup = func(nums []int, cur int) int {
		if nums[cur] == cur {
			return cur
		}
		next := nums[cur]
		nums[cur] = cur
		return dup(nums, next)
	}
	return dup(nums, 0)
}

// array as hashmap(iteration) - TC: O(n), SC: O(1)
func findDuplicate7(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}
	return nums[0]
}

// mark negative solution - TC: O(n), SC: O(1)
func findDuplicate9(nums []int) int {
	duplicate := 0
	for _, num := range nums {
		cur := abs(num)
		if nums[cur] < 0 {
			duplicate = cur
			break
		}

		nums[cur] = -nums[cur]
	}
	return duplicate
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Test_findDuplicate() {
	nums1 := []int{1, 3, 4, 2, 2}
	ans1 := findDuplicate6(nums1)
	log.Printf("ans1: %v", ans1)
	nums2 := []int{3, 1, 4, 3, 2}
	ans2 := findDuplicate6(nums2)
	log.Printf("ans2: %v", ans2)
	nums3 := []int{3, 3, 3, 3, 3}
	ans3 := findDuplicate6(nums3)
	log.Printf("ans3: %v", ans3)
}
