package slidingWindow

import "testing"

type maxProfitCase struct {
	prices   []int
	expected int
}

var maxProfitCases = []maxProfitCase{
	{prices: []int{7, 1, 5, 3, 6, 4}, expected: 5},
	{prices: []int{7, 6, 4, 3, 1}, expected: 0},
}

func TestMaxProfit(t *testing.T) {
	for _, c := range maxProfitCases {
		ans := maxProfitPrefixMin(c.prices)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestMaxProfitDynamicSlidingWindow(t *testing.T) {
	for _, c := range maxProfitCases {
		ans := maxProfitDynamicSlidingWindow(c.prices)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestMaxProfitReviewDP(t *testing.T) {
	for _, c := range maxProfitCases {
		ans := maxProfitReviewDP(c.prices)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestMaxProfitKadane(t *testing.T) {
	for _, c := range maxProfitCases {
		ans := maxProfitKadane(c.prices)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
