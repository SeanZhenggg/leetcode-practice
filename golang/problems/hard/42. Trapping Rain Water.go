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

// prefix maximum/suffix maximum solution - TC: O(n), SC: O(n)
func trap1(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	maxV := 0
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	for i := 0; i < len(height); i++ {
		lMax, rMax := leftMax[i], rightMax[i]

		minH := min(lMax, rMax)
		if minH-height[i] > 0 {
			maxV += min(lMax, rMax) - height[i]
		}
	}

	return maxV
}

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
	ans1 := trap1(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 2, 0, 3, 2, 5}
	ans2 := trap1(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{0, 1, 2, 0, 0, 1, 3, 2, 5, 0}
	ans3 := trap1(case3)
	log.Printf("ans3: %v", ans3)
}

func Test_TrapReview2() {
	case1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans1 := trap2(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{4, 2, 0, 3, 2, 5}
	ans2 := trap2(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{0, 1, 2, 0, 0, 1, 3, 2, 5, 0}
	ans3 := trap2(case3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{0, 1, 0, 3, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans4 := trap2(case4)
	log.Printf("ans4: %v", ans4)
}
