package medium

import (
	"log"
	"math"
)

// brute force solution - O(m*n), where m is the maximum pile in piles and n is the number of piles
func minEatingSpeed(piles []int, h int) int {
	maxPile := math.MinInt16
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	for i := 1; i <= maxPile; i++ {
		totalHours := 0
		for j := 0; j < len(piles); j++ {
			totalHours += int(math.Ceil(float64(piles[j]) / float64(i)))
		}
		if totalHours <= h {
			return i
		}
	}
	return 0
}

// binary search solution - O(n * log(m)), where m is the maximum pile in piles and n is the number of piles
func minEatingSpeed2(piles []int, h int) int {
	maxPile := math.MinInt16
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	res := 0
	l, r := 1, maxPile
	for l <= r {
		mid := l + (r-l)/2

		totalHours := 0
		for j := 0; j < len(piles); j++ {
			totalHours += int(math.Ceil(float64(piles[j]) / float64(mid)))
		}

		if totalHours <= h {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func Test_minEatingSpeed() {
	piles1 := []int{3, 6, 7, 11}
	h1 := 8
	ans1 := minEatingSpeed(piles1, h1)
	log.Printf("ans1: %v", ans1)
	piles2 := []int{30, 11, 23, 4, 20}
	h2 := 5
	ans2 := minEatingSpeed(piles2, h2)
	log.Printf("ans2: %v", ans2)
	piles3 := []int{30, 11, 23, 4, 20}
	h3 := 6
	ans3 := minEatingSpeed(piles3, h3)
	log.Printf("ans3: %v", ans3)

}

func Test_minEatingSpeed2() {
	piles1 := []int{3, 6, 7, 11}
	h1 := 8
	ans1 := minEatingSpeed2(piles1, h1)
	log.Printf("ans1: %v", ans1)
	piles2 := []int{30, 11, 23, 4, 20}
	h2 := 5
	ans2 := minEatingSpeed2(piles2, h2)
	log.Printf("ans2: %v", ans2)
	piles3 := []int{30, 11, 23, 4, 20}
	h3 := 6
	ans3 := minEatingSpeed2(piles3, h3)
	log.Printf("ans3: %v", ans3)

}
