package stack

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
		//days := 1
		//j = i + days
		for j < n && temperatures[j] <= temperatures[i] {
			if ret[j] == 0 {
				j = n
				break
			}
			j += ret[j]
			//i + days += ret[i+days]
		}

		if j < n {
			ret[i] = j - i
			//ret[i] = i + days - i = days
		}
	}

	return ret
}

func dailyTemperaturesReview(temperatures []int) []int {
	ret := make([]int, len(temperatures))
	st := make([]int, 0, len(temperatures))
	for i, temperature := range temperatures {
		for len(st) > 0 && temperatures[st[len(st)-1]] < temperature {
			top := st[len(st)-1]
			ret[top] = i - top
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	return ret
}

func dailyTemperaturesReview2(temperatures []int) []int {
	ret := make([]int, 0, len(temperatures))

	hottestFromRight := temperatures[len(temperatures)-1]
	for i := len(temperatures) - 2; i >= 0; i-- {
		if hottestFromRight <= temperatures[i] {
			hottestFromRight = temperatures[i]
			continue
		}

		days := 1
		for i+days < len(temperatures)-1 && temperatures[i] >= temperatures[i+days] {
			days += ret[i+days]
		}

		ret[i] = days
	}

	return ret
}

type item struct {
	Val int
	Idx int
}

func dailyTemperaturesMonotonicStack(temperatures []int) []int {
	monotonicDescSt := make([]item, 0, len(temperatures))
	ans := make([]int, len(temperatures))

	if len(temperatures) == 0 {
		return ans
	}

	for i := 0; i < len(temperatures); i++ {
		for len(monotonicDescSt) > 0 && temperatures[i] > monotonicDescSt[len(monotonicDescSt)-1].Val {
			pop := monotonicDescSt[len(monotonicDescSt)-1]
			ans[pop.Idx] = i - pop.Idx
			monotonicDescSt = monotonicDescSt[:len(monotonicDescSt)-1]
		}
		monotonicDescSt = append(monotonicDescSt, item{Val: temperatures[i], Idx: i})
	}

	if len(monotonicDescSt) > 0 {
		for i := 0; i < len(monotonicDescSt); i++ {
			ans[monotonicDescSt[i].Idx] = 0
		}
	}

	return ans
}

func dailyTemperaturesArray(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	n := len(temperatures)
	if n == 0 {
		return ans
	}

	for i := n - 2; i >= 0; i-- {
		j := i + 1

		for j < n && temperatures[j] <= temperatures[i] {
			if ans[j] == 0 {
				j = n
				break
			}
			j += ans[j]
		}

		if j < n {
			ans[i] = j - i
		}
	}

	return ans
}
