package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"log"
)

func lastStoneWeight(stones []int) int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		aAsserted, bAsserted := a.(int), b.(int)

		return bAsserted - aAsserted
	})

	for _, stone := range stones {
		pq.Enqueue(stone)
	}

	for pq.Size() > 1 {
		y, _ := pq.Dequeue()
		yVal := y.(int)
		x, _ := pq.Dequeue()
		xVal := x.(int)

		if yVal != xVal {
			pq.Enqueue(yVal - xVal)
		}
	}

	last, ok := pq.Peek()
	if !ok {
		return 0
	}
	return last.(int)
}

func Test_LastStoneWeight() {
	case1 := []int{2, 7, 4, 1, 8, 1}
	ans1 := lastStoneWeight(case1)
	log.Printf("ans1: %v", ans1)

	case2 := []int{1}
	ans2 := lastStoneWeight(case2)
	log.Printf("ans2: %v", ans2)

	case3 := []int{2, 7, 4, 1, 8, 3, 9, 12, 15}
	//2, 7, 4, 1, 8, 3, 9, 3
	//2, 7, 4, 1, 1, 3, 3
	//2, 3, 1, 1, 3, 3
	//2, 1, 1, 3
	//1, 1, 1
	//1
	ans3 := lastStoneWeight(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{2, 7, 4, 1, 8, 3, 9, 12, 14}
	//2, 7, 4, 1, 8, 3, 9, 2
	//2, 7, 4, 1, 1, 3, 2
	//2, 3, 1, 1, 3, 2
	//2, 1, 1, 2
	//1, 1
	//0
	ans4 := lastStoneWeight(case4)
	log.Printf("ans4: %v", ans4)
}
