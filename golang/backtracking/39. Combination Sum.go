package backtracking

import (
	"log"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	current := make([]int, 0)
	ret := make([][]int, 0)
	sum := 0

	var backtracking func(candidates []int, target int)
	backtracking = func(candidates []int, target int) {
		if sum > target {
			return
		}
		if sum == target {
			p := make([]int, len(current))
			copy(p, current)
			ret = append(ret, p)
			return
		}

		for i := 0; i < len(candidates); i++ {
			if len(current) > 0 && current[len(current)-1] > candidates[i] {
				continue
			}

			sum += candidates[i]
			current = append(current, candidates[i])
			if sum <= target {
				backtracking(candidates, target)
			}
			current = current[:len(current)-1]
			sum -= candidates[i]
		}
	}

	backtracking(candidates, target)
	return ret
}

func Test_combinationSum() {
	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	ans1 := combinationSum(candidates1, target1)
	log.Printf("ans1: %v", ans1)

	candidates2 := []int{2, 3, 5}
	target2 := 8
	ans2 := combinationSum(candidates2, target2)
	log.Printf("ans2: %v", ans2)

	candidates3 := []int{2}
	target3 := 1
	ans3 := combinationSum(candidates3, target3)
	log.Printf("ans3: %v", ans3)
}
