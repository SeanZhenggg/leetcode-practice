package greedy

import "log"

func canJump(nums []int) bool {
	pathMap := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		pathMap[i] = true
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		// log.Printf("i now: %d", i)
		if i == len(nums)-1 {
			return true
		}

		if nums[i] == 0 {
			// log.Printf("i: %d failed, record false", i)
			pathMap[i] = false
			return false
		}

		for j := nums[i]; j >= 1; j-- {
			// log.Printf("try i+j: %d", i+j)
			if !pathMap[i+j] {
				continue
			}

			if dfs(i + j) {
				return true
			}
		}
		// log.Printf("i: %d failed, record false", i)
		pathMap[i] = false
		return false
	}

	return dfs(0)
}

func Test_canJump() {
	case1 := []int{2, 3, 1, 1, 4}
	ans1 := canJump(case1)
	log.Printf("ans1: %v", ans1)
	log.Println()
	case2 := []int{3, 2, 1, 0, 4}
	ans2 := canJump(case2)
	log.Printf("ans2: %v", ans2)
	log.Println()
	case3 := []int{2, 1, 0, 1, 4}
	ans3 := canJump(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{3, 2, 4, 0, 5, 0, 2, 1, 0, 3, 10}
	ans4 := canJump(case4)
	log.Printf("ans4: %v", ans4)
}
