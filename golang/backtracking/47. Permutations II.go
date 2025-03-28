package backtracking

import (
	"log"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)

	sort.Ints(nums)

	var backtrack func(current []int, left []int)
	backtrack = func(current []int, left []int) {
		if len(current) == len(nums) {
			p := make([]int, len(current))
			copy(p, current)
			ret = append(ret, p)
			return
		}

		for i := 0; i < len(left); i++ {
			if i > 0 && left[i] == left[i-1] {
				continue
			}
			current = append(current, left[i])
			l := append(append([]int{}, left[:i]...), left[i+1:]...)
			backtrack(current, l)
			current = current[:len(current)-1]
		}
	}

	backtrack([]int{}, nums)
	return ret
}

func Test_permuteUnique() {
	nums1 := []int{1, 1, 2}
	ans1 := permuteUnique(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{1, 2, 3}
	ans2 := permuteUnique(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{1, 1, 1}
	ans3 := permuteUnique(nums3)
	log.Printf("ans3: %v", ans3)

	nums4 := []int{3, 3, 0, 3}
	ans4 := permuteUnique(nums4)
	log.Printf("ans4: %v", ans4)
}
