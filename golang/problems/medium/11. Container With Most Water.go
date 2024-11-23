package medium

import (
	"log"
	"math"
)

func maxArea(height []int) int {
	var l, r = 0, len(height) - 1
	var maxAr float64
	var minH int
	for l <= r {
		if height[l] < height[r] {
			minH = height[l]
		} else {
			minH = height[r]
		}
		maxAr = math.Max(maxAr, float64(minH*(r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return int(maxAr)
}

func Test_MaxArea() {
	case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans1 := maxArea(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{1, 1}
	ans2 := maxArea(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{3, 11, 4, 6, 8, 5, 8, 1, 9}
	ans3 := maxArea(case3)
	log.Printf("ans3: %v", ans3)
}
