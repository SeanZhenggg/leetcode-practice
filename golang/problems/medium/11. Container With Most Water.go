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

func maxAreaReview(height []int) int {
	l, r := 0, len(height)-1
	maxVolume := 0
	// we don't need to concern about l == r, because it will cause the width to be 0
	for l < r {
		minHeight := height[l]
		if minHeight > height[r] {
			minHeight = height[r]
		}

		volume := (r - l) * minHeight

		if maxVolume < volume {
			maxVolume = volume
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxVolume
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

func Test_MaxAreaReview() {
	case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans1 := maxAreaReview(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{1, 1}
	ans2 := maxAreaReview(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{3, 11, 4, 6, 8, 5, 8, 1, 9}
	ans3 := maxAreaReview(case3)
	log.Printf("ans3: %v", ans3)
}
