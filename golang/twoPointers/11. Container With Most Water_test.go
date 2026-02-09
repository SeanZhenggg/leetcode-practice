package twoPointers

import (
	"log"
	"testing"
)

func TestMaxArea(t *testing.T) {
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

func TestMaxAreaReview(t *testing.T) {
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
