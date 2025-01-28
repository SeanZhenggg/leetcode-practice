package stack

import (
	"log"
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

func Test_CarFleet() {
	target1 := 12
	pos1 := []int{10, 8, 0, 5, 3}
	speed1 := []int{2, 4, 1, 1, 3}
	ans1 := carFleet(target1, pos1, speed1)
	log.Printf("ans1: %v", ans1)
	target2 := 10
	pos2 := []int{3}
	speed2 := []int{3}
	ans2 := carFleet(target2, pos2, speed2)
	log.Printf("ans2: %v", ans2)
	target3 := 100
	pos3 := []int{0, 2, 4}
	speed3 := []int{4, 2, 1}
	ans3 := carFleet(target3, pos3, speed3)
	log.Printf("ans3: %v", ans3)
	target4 := 10
	pos4 := []int{6, 8}
	speed4 := []int{3, 2}
	ans4 := carFleet(target4, pos4, speed4)
	log.Printf("ans4: %v", ans4)
	target5 := 10
	pos5 := []int{8, 3, 7, 4, 6, 5}
	speed5 := []int{4, 4, 4, 4, 4, 4}
	ans5 := carFleet(target5, pos5, speed5)
	log.Printf("ans5: %v", ans5)
	target6 := 10
	pos6 := []int{2, 4}
	speed6 := []int{3, 2}
	ans6 := carFleet(target6, pos6, speed6)
	log.Printf("ans6: %v", ans6)
	target7 := 31
	pos7 := []int{5, 26, 18, 25, 29, 21, 22, 12, 19, 6}
	speed7 := []int{7, 6, 6, 4, 3, 4, 9, 7, 6, 4}
	ans7 := carFleet(target7, pos7, speed7)
	log.Printf("ans7: %v", ans7)
}
