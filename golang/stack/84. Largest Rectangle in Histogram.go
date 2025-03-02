package stack

import "log"

type n struct {
	Height      int
	LeftMostIdx int
}

// first solution
func largestRectangleArea(heights []int) int {
	// monotonic increasing stack
	st := make([]n, 0, len(heights))
	maxArea := 0
	for i, height := range heights {
		leftMostIdx := i
		for len(st) > 0 && st[len(st)-1].Height >= height {
			top := st[len(st)-1]
			leftMostIdx = top.LeftMostIdx
			st = st[:len(st)-1]
		}
		st = append(st, n{LeftMostIdx: leftMostIdx, Height: height})

		if len(st) > 0 {
			for j := len(st) - 1; j >= 0; j-- {
				el := st[j]
				width := i - el.LeftMostIdx + 1
				area := width * el.Height
				maxArea = max(maxArea, area)
			}
		}
	}
	return maxArea
}

type n2 struct {
	Height int
	Start  int
}

// better solution
func largestRectangleArea2(heights []int) int {
	// monotonic increasing stack
	st := make([]n2, 0, len(heights))
	maxArea := 0
	for i, height := range heights {
		start := i
		for len(st) > 0 && st[len(st)-1].Height > height {
			top := st[len(st)-1]
			maxArea = max(maxArea, (i-top.Start)*top.Height)
			st = st[:len(st)-1]
			start = top.Start
		}
		st = append(st, n2{Start: start, Height: height})
	}

	if len(st) > 0 {
		for j := 0; j < len(st); j++ {
			maxArea = max(maxArea, (len(heights)-st[j].Start)*st[j].Height)
		}
	}
	return maxArea
}

func largestRectangleAreaReview(heights []int) int {
	// push (idx where current height could extend to, height) into stack
	// if current height is lesser than top of stack, pop out each height greater than current height
	// calculate the max area

	st := make([][2]int, 0)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		start := i
		for len(st) > 0 && st[len(st)-1][1] >= heights[i] {
			start = st[len(st)-1][0]
			height := st[len(st)-1][1]
			maxArea = max(maxArea, (i-start)*height)
			st = st[:len(st)-1]
		}

		st = append(st, [2]int{start, heights[i]})
	}

	for i := 0; i < len(st); i++ {
		maxArea = max(maxArea, (len(heights)-st[i][0])*st[i][1])
	}

	return maxArea
}

func Test_largestRectangleArea() {
	case1 := []int{2, 1, 5, 6, 2, 3}
	ans1 := largestRectangleArea(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{2, 4}
	ans2 := largestRectangleArea(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{2, 2, 5, 6, 2, 3}
	ans3 := largestRectangleArea(case3)
	log.Printf("ans3: %v", ans3)
}

func Test_largestRectangleArea2() {
	case1 := []int{2, 1, 5, 6, 2, 3}
	ans1 := largestRectangleArea2(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{2, 4}
	ans2 := largestRectangleArea2(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{2, 2, 5, 6, 2, 3}
	ans3 := largestRectangleArea2(case3)
	log.Printf("ans3: %v", ans3)
}

func Test_largestRectangleAreaReview() {
	case1 := []int{2, 1, 5, 6, 2, 3}
	ans1 := largestRectangleAreaReview(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{2, 4}
	ans2 := largestRectangleAreaReview(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{2, 2, 5, 6, 2, 3}
	ans3 := largestRectangleAreaReview(case3)
	log.Printf("ans3: %v", ans3)
}
