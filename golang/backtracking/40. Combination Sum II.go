package backtracking

import (
	"log"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	current := make([]int, 0)
	var backtracking func(idx int, sum int)
	backtracking = func(idx int, sum int) {
		if sum > target {
			return
		} else if sum == target {
			p := make([]int, len(current))
			copy(p, current)
			ret = append(ret, p)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i-1] == candidates[i] {
				continue
			}
			if len(current) > 0 && current[len(current)-1] > candidates[i] {
				continue
			}
			current = append(current, candidates[i])
			sum += candidates[i]
			backtracking(i+1, sum)
			sum -= candidates[i]
			current = current[:len(current)-1]
		}
	}

	backtracking(0, 0)

	return ret
}

// 1 1 2 2 4 5 6 7 10
// 10 7 6 5 2 1 1
func Test_combinationSum2() {
	candidates1 := []int{10, 1, 2, 7, 6, 1, 5, 2, 4}
	target1 := 8
	ans1 := combinationSum2(candidates1, target1)
	log.Printf("ans1: %v", ans1)

	candidates2 := []int{2, 5, 2, 1, 2}
	target2 := 5
	ans2 := combinationSum2(candidates2, target2)
	log.Printf("ans2: %v", ans2)

}
