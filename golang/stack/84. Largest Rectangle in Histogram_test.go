package stack

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case2", []int{2, 4}, 4},
		{"case3", []int{2, 2, 5, 6, 2, 3}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea(tt.input)
			if result != tt.expect {
				t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.input, result, tt.expect)
			}
		})
	}
}

func Test_largestRectangleArea2(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case2", []int{2, 4}, 4},
		{"case3", []int{2, 2, 5, 6, 2, 3}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleArea2(tt.input)
			if result != tt.expect {
				t.Errorf("largestRectangleArea2(%v) = %d, want %d", tt.input, result, tt.expect)
			}
		})
	}
}

func Test_largestRectangleAreaReview(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case2", []int{2, 4}, 4},
		{"case3", []int{2, 2, 5, 6, 2, 3}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleAreaReview(tt.input)
			if result != tt.expect {
				t.Errorf("largestRectangleAreaReview(%v) = %d, want %d", tt.input, result, tt.expect)
			}
		})
	}
}

func Test_largestRectangleAreaMonoticStack(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect int
	}{
		{"case1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case2", []int{2, 4}, 4},
		{"case3", []int{2, 2, 5, 6, 2, 3}, 12},
		{"case3", []int{0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestRectangleAreaMonoticStack(tt.input)
			if result != tt.expect {
				t.Errorf("largestRectangleAreaMonoticStack(%v) = %d, want %d", tt.input, result, tt.expect)
			}
		})
	}
}
