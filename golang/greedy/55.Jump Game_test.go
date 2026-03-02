package greedy

import "testing"

type canJumpCase struct {
	nums     []int
	expected bool
}

var canJumpCases = []canJumpCase{
	{nums: []int{2, 3, 1, 1, 4}, expected: true},
	{nums: []int{3, 2, 1, 0, 4}, expected: false},
	{nums: []int{2, 1, 0, 1, 4}, expected: false},
	{nums: []int{3, 2, 4, 0, 5, 0, 2, 1, 0, 3, 10}, expected: true},
}

func Test_canJump(t *testing.T) {
	for _, c := range canJumpCases {
		ans := canJump(c.nums)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func Test_canJump2(t *testing.T) {
	for _, c := range canJumpCases {
		ans := canJump2(c.nums)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func Test_canJump3(t *testing.T) {
	for _, c := range canJumpCases {
		ans := canJump3(c.nums)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func Test_canJump4(t *testing.T) {
	for _, c := range canJumpCases {
		ans := canJump4(c.nums)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
