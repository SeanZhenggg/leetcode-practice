package backtracking

import (
	"log"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ret := make([][]int, 0)

	sort.Ints(candidates)
	current := make([]int, 0)
	sum := 0

	var backtracking func(idx int)
	backtracking = func(idx int) {
		if sum > target {
			return
		}
		if sum == target {
			p := make([]int, len(current))
			copy(p, current)
			ret = append(ret, p)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i-1] == candidates[i] {
				continue
			}

			sum += candidates[i]
			current = append(current, candidates[i])
			backtracking(i + 1)
			current = current[:len(current)-1]
			sum -= candidates[i]
		}
	}

	backtracking(0)

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
