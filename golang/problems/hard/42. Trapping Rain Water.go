package hard

import "log"

// two pointers solution - TC: O(n), SC: O(1)
func trap(height []int) int {
	l, r := 0, len(height)-1
	lMax, rMax := height[l], height[r]
	var ret int

	for l < r {
		if height[l] <= height[r] {
			l++
			if height[l] < lMax {
				ret += lMax - height[l]
			} else {
				lMax = height[l]
			}
		} else {
			r--
			if height[r] < rMax {
				ret += rMax - height[r]
			} else {
				rMax = height[r]
			}
		}
	}

	return ret
}

// prefix/suffix maximum solution - TC: O(n), SC: O(n)
func trap1(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	maxV := 0
	for i := 0; i < len(height); i++ {
		maxV += min(leftMax[i], rightMax[i]) - height[i]
	}

	return maxV
}

// monotonic stack solution - TC: O(n), SC: O(n)
func trap2(height []int) int {
	st := make([]int, 0)
	sum := 0

	for i := 0; i < len(height); i++ {
		for len(st) > 0 && height[st[len(st)-1]] <= height[i] {
			bottom := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			l := st[len(st)-1]
			r := i

			minHeight := min(height[l], height[r])
			bottomHeight := height[bottom]
			width := r - l - 1
			sum += width * (minHeight - bottomHeight)
		}
		st = append(st, i)
	}

	return sum
}

func trapReview1(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if i > 0 {
			leftMax[i] = max(leftMax[i-1], height[i-1])
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		if i < len(height)-1 {
			rightMax[i] = max(rightMax[i+1], height[i+1])
		}
	}

	vol := 0
	for i := 0; i < len(height); i++ {
		vol += max(0, min(leftMax[i], rightMax[i])-height[i])
	}
	return vol
}

func trapReview2(height []int) int {
	leftMax := height[0]
	rightMax := height[len(height)-1]

	vol := 0
	l, r := 0, len(height)-1
	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			vol += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			vol += rightMax - height[r]
		}
	}
	return vol
}

func trapReview3(height []int) int {
	st := make([]int, 0, len(height)-1)
	var l, r = 0, 0
	vol := 0
	for i := 0; i < len(height); i++ {
		for len(st) > 0 && height[st[len(st)-1]] < height[i] {
			r = i
			bottom := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			l = st[len(st)-1]

			vol += (min(height[r], height[l]) - height[bottom]) * (r - l - 1)
		}

		st = append(st, i)
	}
	return vol
}

func Test_Trap() {
	case1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans1 := trap(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 2, 0, 3, 2, 5}
	ans2 := trap(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{0, 1, 2, 0, 0, 1, 3, 2, 5, 0}
	ans3 := trap(case3)
	log.Printf("ans3: %v", ans3)
}

func Test_TrapReview() {
	case1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans1 := trapReview3(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 2, 0, 3, 2, 5}
	ans2 := trapReview3(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{0, 1, 2, 0, 0, 1, 3, 2, 5, 0}
	ans3 := trapReview3(case3)
	log.Printf("ans3: %v", ans3)
}
