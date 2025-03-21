package backtracking

import (
	"log"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	current := make([]int, 0)
	ret := make([][]int, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i])
			appendData(&ret, current)
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return ret
}

func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)
	current := make([]int, 0)
	ret := make([][]int, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx >= len(nums) {
			appendData(&ret, current)
			return
		}

		current = append(current, nums[idx])
		backtrack(idx + 1)
		current = current[:len(current)-1]

		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx += 1
		}
		backtrack(idx + 1)
	}

	backtrack(0)
	return ret
}

func appendData(ret *[][]int, src []int) {
	p := make([]int, len(src))
	copy(p, src)
	*ret = append(*ret, p)
}

func Test_subsetsWithDup() {
	nums1 := []int{1, 2, 2}
	ans1 := subsetsWithDup(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0}
	ans2 := subsetsWithDup(nums2)
	log.Printf("ans2: %v", ans2)
}

func Test_subsetsWithDup2() {
	nums1 := []int{1, 2, 2}
	ans1 := subsetsWithDup2(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0}
	ans2 := subsetsWithDup2(nums2)
	log.Printf("ans2: %v", ans2)

}
