package arraysAndHashing

import (
	"sort"
)

// 1,3,-2,-5,4,8 target=6
// 1,3,4 = 8
// 1,-2,8 = 7
// 3,4,-2 = 5
// 3,-5,8 = 6
func threeSumClosest(nums []int, target int) int {
	curClosest := 100000

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				l++
				if abs(target, sum) < abs(target, curClosest) {
					curClosest = sum
				}
			} else if sum > target {
				r--
				if abs(target, sum) < abs(target, curClosest) {
					curClosest = sum
				}
			} else {
				curClosest = sum
				return curClosest
			}
		}
	}
	return curClosest
}

func abs(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}
