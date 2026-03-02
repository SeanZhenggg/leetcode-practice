package arraysAndHashing

import "testing"

type threeSumClosestCase struct {
	nums     []int
	target   int
	expected int
}

var threeSumClosestCases = []threeSumClosestCase{
	{nums: []int{1, 3, -2, -5, 4, 8}, target: 6, expected: 6},
	{nums: []int{-1, 2, 1, -4}, target: 1, expected: 2},
	{nums: []int{0, 0, 0}, target: 1, expected: 0},
	{nums: []int{1, 1, 1, 1}, target: -100, expected: 3},
}

func TestThreeSumClosest(t *testing.T) {
	for _, c := range threeSumClosestCases {
		ans := threeSumClosest(c.nums, c.target)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
