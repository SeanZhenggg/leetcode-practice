package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"log"
)

type node struct {
	Task      byte
	Remainder int
}

// failed
func leastInterval(tasks []byte, n int) int {
	var minIntervals int

	// pq (min-heap) to record all the task gap remainder
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		aVal, bVal := a.(node), b.(node)

		return aVal.Remainder - bVal.Remainder
	})

	// hash map to record number of tasks
	m := make(map[byte]int)

	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
	}

	for k := range m {
		pq.Enqueue(node{k, 0})
	}

	// loop until queue has no task left
	for pq.Size() > 0 {
		value, _ := pq.Peek()
		valueNode := value.(node)

		if valueNode.Remainder > minIntervals {
			minIntervals++
			continue
		}

		minIntervals++
		pq.Dequeue()
		m[valueNode.Task]--

		if m[valueNode.Task] > 0 {
			valueNode.Remainder = n + minIntervals
			pq.Enqueue(valueNode)
		}
	}

	return minIntervals
}

// pq + queue solution, TC:O(n * logn)
func leastInterval2(tasks []byte, n int) int {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		aVal, bVal := a.(int), b.(int)

		return bVal - aVal
	})

	m := make(map[int]int)

	for _, task := range tasks {
		m[int(task-'A')]++
	}

	for _, count := range m {
		pq.Enqueue(count)
	}

	queue := make([][2]int, 0)
	time := 0
	for pq.Size() > 0 || len(queue) > 0 {
		time++

		if pq.Size() == 0 {
			time = queue[0][1]
		} else {
			value, _ := pq.Dequeue()
			val := value.(int)

			val--
			if val > 0 {
				queue = append(queue, [2]int{val, time + n})
			}
		}

		if len(queue) > 0 {
			next := queue[0][1]
			if next == time {
				pq.Enqueue(queue[0][0])
				queue = queue[1:]
			}
		}
	}
	return time
}

func Test_LeastInterval() {
	tasks1 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n1 := 2
	ans1 := leastInterval(tasks1, n1)
	log.Printf("ans1: %v", ans1)

	tasks2 := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n2 := 1
	ans2 := leastInterval(tasks2, n2)
	log.Printf("ans2: %v", ans2)

	tasks3 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n3 := 3
	ans3 := leastInterval(tasks3, n3)
	log.Printf("ans3: %v", ans3)

	tasks4 := []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'}
	n4 := 1
	ans4 := leastInterval2(tasks4, n4)
	log.Printf("ans4: %v", ans4)
}

func Test_LeastInterval2() {
	tasks1 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//'A', 'B', 'i', 'i', 'i', 'i', 'A', 'B','i', 'i', 'i', 'i', 'A', 'B'
	n1 := 2
	ans1 := leastInterval2(tasks1, n1)
	log.Printf("ans1: %v", ans1)

	tasks2 := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n2 := 1
	ans2 := leastInterval2(tasks2, n2)
	log.Printf("ans2: %v", ans2)

	tasks3 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n3 := 3
	ans3 := leastInterval2(tasks3, n3)
	log.Printf("ans3: %v", ans3)

	tasks4 := []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'}
	n4 := 1
	ans4 := leastInterval2(tasks4, n4)
	log.Printf("ans4: %v", ans4)

}
