package medium

import "log"

func dailyTemperatures(temperatures []int) []int {
	// monotonic decreasing stack
	st := make([]int, 0, len(temperatures))
	ret := make([]int, len(temperatures))

	for i := range temperatures {
		for len(st) > 0 && temperatures[st[len(st)-1]] < temperatures[i] {
			top := st[len(st)-1]
			st = st[:len(st)-1]

			ret[top] = i - top
		}

		st = append(st, i)
	}

	return ret
}

func Test_DailyTemperatures() {
	case1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans1 := dailyTemperatures(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{30, 40, 50, 60}
	ans2 := dailyTemperatures(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{30, 60, 90}
	ans3 := dailyTemperatures(case3)
	log.Printf("ans3: %v", ans3)

}
