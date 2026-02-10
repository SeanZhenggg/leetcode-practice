package slidingWindow

import (
	"log"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	case1 := []int{7, 1, 5, 3, 6, 4}
	ans1 := maxProfitPrefixMin(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{7, 6, 4, 3, 1}
	ans2 := maxProfitPrefixMin(case2)
	log.Printf("ans2: %v", ans2)
}

func TestMaxProfitDynamicSlidingWindow(t *testing.T) {
	case1 := []int{7, 1, 5, 3, 6, 4}
	ans1 := maxProfitDynamicSlidingWindow(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{7, 6, 4, 3, 1}
	ans2 := maxProfitDynamicSlidingWindow(case2)
	log.Printf("ans2: %v", ans2)
}

func TestMaxProfitReviewDP(t *testing.T) {
	case1 := []int{7, 1, 5, 3, 6, 4}
	ans1 := maxProfitReviewDP(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{7, 6, 4, 3, 1}
	ans2 := maxProfitReviewDP(case2)
	log.Printf("ans2: %v", ans2)
}

func TestMaxProfitKadane(t *testing.T) {
	case1 := []int{7, 1, 5, 3, 6, 4}
	ans1 := maxProfitKadane(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{7, 6, 4, 3, 1}
	ans2 := maxProfitKadane(case2)
	log.Printf("ans2: %v", ans2)
}
