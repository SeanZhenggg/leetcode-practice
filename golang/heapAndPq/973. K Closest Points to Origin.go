package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"log"
)

// min heap, TC: O(n * logn), SC: O(n)
func kClosest(points [][]int, k int) [][]int {
	ret := make([][]int, 0, k)
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		x1, y1 := a.([]int)[0], a.([]int)[1]
		x2, y2 := b.([]int)[0], b.([]int)[1]
		dis1 := x1*x1 + y1*y1
		dis2 := x2*x2 + y2*y2

		if dis1 < dis2 {
			return -1
		} else if dis1 > dis2 {
			return 1
		} else {
			return 0
		}
	})

	for i := 0; i < len(points); i++ {
		pq.Enqueue(points[i])
	}

	for ; k > 0; k-- {
		value, _ := pq.Dequeue()
		ret = append(ret, value.([]int))
	}

	return ret
}

// max heap, TC: O(n * logk), SC: O(k)
func kClosest2(points [][]int, k int) [][]int {
	ret := make([][]int, 0, k)
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		x1, y1 := a.([]int)[0], a.([]int)[1]
		x2, y2 := b.([]int)[0], b.([]int)[1]
		dis1 := x1*x1 + y1*y1
		dis2 := x2*x2 + y2*y2

		if dis1 > dis2 {
			return -1
		} else if dis1 < dis2 {
			return 1
		} else {
			return 0
		}
	})

	for i := 0; i < len(points); i++ {
		pq.Enqueue(points[i])

		if pq.Size() > k {
			pq.Dequeue()
		}
	}

	for ; k > 0; k-- {
		value, _ := pq.Dequeue()
		ret = append(ret, value.([]int))
	}

	return ret
}

func Test_KClosest() {
	case1 := [][]int{{1, 3}, {-2, 2}}
	k1 := 1
	ans1 := kClosest(case1, k1)
	log.Printf("ans1: %v", ans1)

	case2 := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k2 := 2
	ans2 := kClosest(case2, k2)
	log.Printf("ans2: %v", ans2)

}

func Test_KClosest2() {
	case1 := [][]int{{1, 3}, {-2, 2}}
	k1 := 1
	ans1 := kClosest2(case1, k1)
	log.Printf("ans1: %v", ans1)

	case2 := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k2 := 2
	ans2 := kClosest2(case2, k2)
	log.Printf("ans2: %v", ans2)

}
