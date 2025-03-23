package backtracking

import "log"

func permute(nums []int) [][]int {
	ret := make([][]int, 0)

	var backtrack func(current []int, left []int)
	backtrack = func(current []int, left []int) {
		if len(current) == len(nums) {
			p := make([]int, len(current))
			copy(p, current)
			ret = append(ret, p)
			return
		}

		for i := 0; i < len(left); i++ {
			current = append(current, left[i])
			l := append(append([]int{}, left[:i]...), left[i+1:]...)
			backtrack(current, l)
			current = current[:len(current)-1]
		}
	}

	backtrack([]int{}, nums)
	return ret
}

func permute2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	perms := permute2(nums[1:])
	var res [][]int
	for _, p := range perms {
		for i := 0; i <= len(p); i++ {
			pCopy := append([]int{}, p...)
			pCopy = append(pCopy[:i], append([]int{nums[0]}, pCopy[i:]...)...)
			res = append(res, pCopy)
		}
	}
	return res
}

func Test_permute() {
	nums1 := []int{1, 2, 3}
	ans1 := permute(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0, 1}
	ans2 := permute(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{1}
	ans3 := permute(nums3)
	log.Printf("ans3: %v", ans3)
}

func Test_permute2() {
	nums1 := []int{1, 2, 3}
	ans1 := permute2(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{0, 1}
	ans2 := permute2(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{1}
	ans3 := permute2(nums3)
	log.Printf("ans3: %v", ans3)
}
