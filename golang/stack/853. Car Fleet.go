package stack

import (
	"slices"
)

type node struct {
	Pos   int
	Speed int
	Time  float64
}

func carFleet(target int, position []int, speed []int) int {
	speedMap := make(map[int]int)
	st := make([]node, 0, len(position))

	for i, v := range position {
		speedMap[v] = speed[i]
	}

	slices.SortFunc(position, func(a, b int) int {
		return a - b
	})

	for i := len(position) - 1; i >= 0; i-- {
		dis := target - position[i]
		spe := speedMap[position[i]]
		t := float64(dis) / float64(spe)
		if len(st) == 0 || st[len(st)-1].Time < t {
			st = append(st, node{position[i], spe, t})
		}
	}
	return len(st)
}

func carFleet2(target int, position []int, speed []int) int {
	// 大到小排，看看距離遠的能不能追上前面的人
	m := make(map[int]int)
	for i, v := range position {
		m[v] = speed[i]
	}

	slices.SortFunc(position, func(a, b int) int {
		return a - b
	})
	speed = make([]int, 0, len(position))
	for _, v := range position {
		speed = append(speed, m[v])
	}

	timeStack := new(Stack[float64])
	for i := range position {
		ti := float64(target-position[i]) / float64(speed[i])
		for timeStack.Len() > 0 && timeStack.Top() <= ti {
			timeStack.Pop()
		}
		timeStack.Push(ti)
	}

	return timeStack.Len()
}

func carFleetForIteration(target int, position []int, speed []int) int {
	// 大到小排，看看距離遠的能不能追上前面的人
	var posWithSpeed = make([][2]int, 0, len(position))

	for i, v := range position {
		posWithSpeed = append(posWithSpeed, [2]int{v, speed[i]})
	}

	var curCostTime float64
	var fleetCount int
	slices.SortFunc(posWithSpeed, func(a, b [2]int) int {
		return b[0] - a[0]
	})

	for _, item := range posWithSpeed {
		nextCostTime := float64(target-item[0]) / float64(item[1])
		if curCostTime == 0 || curCostTime < nextCostTime {
			curCostTime = nextCostTime
			fleetCount += 1
		}
	}

	return fleetCount
}
