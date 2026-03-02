package arraysAndHashing

import "testing"

type longestConsecutiveCase struct {
	nums     []int
	expected int
}

var longestConsecutiveCases = []longestConsecutiveCase{
	{nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
	{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
	{nums: []int{159, 86, 72, 85, 87, 160, 73, 73, 158, 161}, expected: 4},
	{nums: []int{1, 0, 1, 2}, expected: 3},
	{nums: []int{3, 4, 5, 4}, expected: 3},
}

func TestLongestConsecutive(t *testing.T) {
	for _, c := range longestConsecutiveCases {
		nums := make([]int, len(c.nums))
		copy(nums, c.nums)
		ans := longestConsecutive(nums)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestLongestConsecutive2(t *testing.T) {
	for _, c := range longestConsecutiveCases {
		nums := make([]int, len(c.nums))
		copy(nums, c.nums)
		ans := longestConsecutive2(nums)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestLongestConsecutive3(t *testing.T) {
	for _, c := range longestConsecutiveCases {
		ans := longestConsecutive3(c.nums)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}

func TestLongestConsecutive4(t *testing.T) {
	for _, c := range longestConsecutiveCases {
		ans := longestConsecutive4(c.nums)
		if ans != c.expected {
			t.Errorf("answer is %d, want %d", ans, c.expected)
		}
	}
}
