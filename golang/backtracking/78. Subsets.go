package backtracking

import "log"

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	start := make([]int, 0)
	result = append(result, start)

	var backtrack func(nums []int, current []int, idx int)
	backtrack = func(nums []int, current []int, idx int) {
		if idx > len(nums)-1 {
			return
		}

		for idx < len(nums) {
			current = append(current, nums[idx])
			c := make([]int, len(current))
			copy(c, current)
			result = append(result, c)
			backtrack(nums, current, idx+1)
			current = current[:len(current)-1]
			idx++
		}
	}

	backtrack(nums, start, 0)

	return result
}

func Test_subsets() {
	nums1 := []int{1, 2, 3}
	ans1 := subsets(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0}
	ans2 := subsets(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{1, 2, 4, 5}
	ans3 := subsets(nums3)
	log.Printf("ans3: %v", ans3)

}
