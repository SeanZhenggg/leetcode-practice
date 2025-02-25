package slidingWindow

import (
	"log"
)

// max pq solution - O(n * logk)
func maxSlidingWindow(nums []int, k int) []int {
	pq := NewPriorityQueue(func(a, b interface{}) int {
		return b.([2]int)[0] - a.([2]int)[0]
	})

	output := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		pq.push([2]int{nums[i], i})

		for pq.top().([2]int)[1] <= i-k {
			pq.poll()
		}

		if i+1 >= k {
			output = append(output, pq.top().([2]int)[0])
		}
	}

	return output
}

func Test_MaxSlidingWindow() {
	case1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	ans1 := maxSlidingWindow(case1, k1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{1}
	k2 := 1
	ans2 := maxSlidingWindow(case2, k2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{1, -1}
	k3 := 1
	ans3 := maxSlidingWindow(case3, k3)
	log.Printf("ans3: %v", ans3)
}

type PriorityQueue struct {
	heap        []interface{}
	compareFunc func(a, b interface{}) int
}

func NewPriorityQueue(compareFunc func(a, b interface{}) int) *PriorityQueue {
	pq := &PriorityQueue{
		heap:        make([]interface{}, 0),
		compareFunc: compareFunc,
	}

	return pq
}

func (pq *PriorityQueue) build() {
	for i := len(pq.heap) / 2; i >= 1; i-- {
		pq.heapify(i)
	}
}

func (pq *PriorityQueue) heapify(index int) {
	length := pq.getLength()

	minIndex := index
	l := 2*index + 1
	r := 2*index + 2
	if l < length && pq.compareFunc(pq.heap[minIndex], pq.heap[l]) > 0 {
		minIndex = l
	}

	if r < length && pq.compareFunc(pq.heap[minIndex], pq.heap[r]) > 0 {
		minIndex = r
	}

	if pq.compareFunc(pq.heap[index], pq.heap[minIndex]) > 0 {
		pq.swap(minIndex, index)
	}

	if minIndex != index {
		pq.heapify(minIndex)
	}
}

func (pq *PriorityQueue) getLength() int {
	return len(pq.heap)
}

func (pq *PriorityQueue) swap(idx1 int, idx2 int) {
	pq.heap[idx1], pq.heap[idx2] = pq.heap[idx2], pq.heap[idx1]
}

func (pq *PriorityQueue) poll() interface{} {
	minIndex := 0
	minNode := pq.heap[minIndex]

	pq.heap[minIndex] = pq.heap[len(pq.heap)-1]

	pq.heap = pq.heap[:len(pq.heap)-1]

	pq.heapify(minIndex)

	return minNode
}

func (pq *PriorityQueue) top() interface{} {
	return pq.heap[0]
}

func (pq *PriorityQueue) push(n interface{}) {
	pq.heap = append(pq.heap, n)
	index := len(pq.heap) - 1
	for index > 0 && pq.compareFunc(pq.heap[(index-1)/2], pq.heap[index]) > 0 {
		pq.swap((index-1)/2, index)
		index = (index - 1) / 2
	}
}

func maxSlidingWindowReview(nums []int, k int) []int {
	deque := make([]int, 0, len(nums))
	ret := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i-k+1 > deque[0] {
			deque = deque[1:]
		}

		if i+1 >= k {
			ret = append(ret, nums[deque[0]])
		}
	}

	return ret
}

func maxSlidingWindowReview2(nums []int, k int) []int {
	pq := NewPriorityQueue(func(a, b interface{}) int {
		return b.([2]int)[1] - a.([2]int)[1]
	})
	ret := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		pq.push([2]int{i, nums[i]})

		for pq.top().([2]int)[0] <= i-k {
			pq.poll()
		}

		if i+1 >= k {
			ret = append(ret, pq.top().([2]int)[1])
		}
	}
	return ret
}

func Test_MaxSlidingWindowReview() {
	case1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	ans1 := maxSlidingWindowReview(case1, k1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{1}
	k2 := 1
	ans2 := maxSlidingWindowReview(case2, k2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{1, -1}
	k3 := 1
	ans3 := maxSlidingWindowReview(case3, k3)
	log.Printf("ans3: %v", ans3)
}

func Test_MaxSlidingWindowReview2() {
	case1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	ans1 := maxSlidingWindowReview2(case1, k1)
	log.Printf("ans1: %v", ans1)
	case2 := []int{1}
	k2 := 1
	ans2 := maxSlidingWindowReview2(case2, k2)
	log.Printf("ans2: %v", ans2)
	case3 := []int{1, -1}
	k3 := 1
	ans3 := maxSlidingWindowReview2(case3, k3)
	log.Printf("ans3: %v", ans3)
	case4 := []int{1, 3, 1, 2, 0, 5}
	k4 := 3
	ans4 := maxSlidingWindowReview2(case4, k4)
	log.Printf("ans4: %v", ans4)
	case5 := []int{9, 10, 9, -7, -4, -8, 2, -6}
	k5 := 5
	ans5 := maxSlidingWindowReview2(case5, k5)
	log.Printf("ans5: %v", ans5)
}
