package arraysAndHashing

import (
	"reflect"
	"testing"
)

type productExceptSelfCase struct {
	nums     []int
	expected []int
}

var productExceptSelfCases = []productExceptSelfCase{
	{nums: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
	{nums: []int{-1, 1, 0, -3, 3}, expected: []int{0, 0, 9, 0, 0}},
	{nums: []int{8, 2, -1, 3, 5}, expected: []int{-30, -120, 240, -80, -48}},
}

func TestProductExceptSelf(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelf(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestProductExceptSelf2(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelf2(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestProductExceptSelf3(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelf3(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestProductExceptSelf4(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelf4(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestProductExceptSelf5(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelf5(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestProductExceptSelfReview(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelfReview(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func TestProductExceptSelfReview2(t *testing.T) {
	for _, c := range productExceptSelfCases {
		ans := productExceptSelfReview2(c.nums)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
