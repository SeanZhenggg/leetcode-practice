package stack

import "log"

// monotonic stack solution - TC: O(n), SC: O(n)
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

// array solution - TC: O(n), SC: O(1)
func dailyTemperatures2(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	hottest := 0
	for i := len(temperatures) - 1; i >= 0; i-- {
		if temperatures[i] >= hottest {
			hottest = temperatures[i]
			continue
		}
		days := 1
		for i+days < len(temperatures) && temperatures[i+days] <= temperatures[i] {
			days += ret[i+days]
		}

		ret[i] = days
	}

	return ret
}

// array solution2 - TC: O(n), SC: O(1)
func dailyTemperatures3(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	n := len(temperatures)
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && temperatures[j] <= temperatures[i] {
			if ret[j] == 0 {
				j = n
				break
			}
			j += ret[j]
		}

		if j < n {
			ret[i] = j - i
		}
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

func Test_DailyTemperatures2() {
	case1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans1 := dailyTemperatures2(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{30, 40, 50, 60}
	ans2 := dailyTemperatures2(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{30, 60, 90}
	ans3 := dailyTemperatures2(case3)
	log.Printf("ans3: %v", ans3)

}

func Test_DailyTemperatures3() {
	case1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans1 := dailyTemperatures3(case1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{30, 40, 50, 60}
	ans2 := dailyTemperatures3(case2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{30, 60, 90}
	ans3 := dailyTemperatures3(case3)
	log.Printf("ans3: %v", ans3)

}
