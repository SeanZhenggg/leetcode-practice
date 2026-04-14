package stack

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

func largestRectangleAreaMonoticStack(heights []int) int {
	// 用 monotonic increasing stack 來處理
	// why? TODO
	// 找到下一個更小的 bar 時:
	// pop stack 頂端的元素
	// pop 出來的 bar 去計算它延伸的最大寬度, 因為 area = w * h , 而 h 在每個 bar 固定, 因此 area 跟 w 成正比
	// 計算完了以後, 將當前 bar 的延伸範圍連同它的 val 存進去 stack
	if len(heights) == 0 {
		return 0
	}

	st := Stack[[2]int]{}
	maxArea := heights[0]
	for i := 0; i < len(heights); i++ {
		leftMost := i
		//maxArea = max(maxArea, heights[i]*1)
		for st.Len() > 0 && st.Top()[0] >= heights[i] {
			top := st.Pop()
			maxArea = max(maxArea, top[0]*(i-top[1]))
			leftMost = top[1]
		}
		st.Push([2]int{heights[i], leftMost})
	}

	for st.Len() > 0 {
		top := st.Pop()
		maxArea = max(maxArea, top[0]*(len(heights)-top[1]))
	}

	return maxArea
}
