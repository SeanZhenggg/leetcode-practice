package hard

import "log"

type node struct {
	Height      int
	LeftMostIdx int
}

// first solution
func largestRectangleArea(heights []int) int {
	// monotonic increasing stack
	st := make([]node, 0, len(heights))
	maxArea := 0
	for i, height := range heights {
		leftMostIdx := i
		for len(st) > 0 && st[len(st)-1].Height >= height {
			top := st[len(st)-1]
			leftMostIdx = top.LeftMostIdx
			st = st[:len(st)-1]
		}
		st = append(st, node{LeftMostIdx: leftMostIdx, Height: height})

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

type node2 struct {
	Height int
	Start  int
}

// better solution
func largestRectangleArea2(heights []int) int {
	// monotonic increasing stack
	st := make([]node2, 0, len(heights))
	maxArea := 0
	for i, height := range heights {
		start := i
		for len(st) > 0 && st[len(st)-1].Height > height {
			top := st[len(st)-1]
			maxArea = max(maxArea, (i-top.Start)*top.Height)
			st = st[:len(st)-1]
			start = top.Start
		}
		st = append(st, node2{Start: start, Height: height})
	}

	if len(st) > 0 {
		for j := 0; j < len(st); j++ {
			maxArea = max(maxArea, (len(heights)-st[j].Start)*st[j].Height)
		}
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
