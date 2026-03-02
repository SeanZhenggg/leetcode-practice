package twoPointers

import "testing"

type maxAreaCase struct {
	height   []int
	expected int
}

var maxAreaCases = []maxAreaCase{
	{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
	{height: []int{1, 1}, expected: 1},
	{height: []int{3, 11, 4, 6, 8, 5, 8, 1, 9}, expected: 63},
}

func TestMaxArea(t *testing.T) {
	for _, c := range maxAreaCases {
		ans := maxArea(c.height)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
