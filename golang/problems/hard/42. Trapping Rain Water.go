package hard

import "log"

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
