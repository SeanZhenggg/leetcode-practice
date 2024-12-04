package hard

import (
	"log"
	"strconv"
)

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0, len(nums)-k+1)
	m := map[int]int{}

	init := make([]node, 0, k)
	for i := 0; i < k; i++ {
		init = append(init, node{Key: strconv.Itoa(i), Weight: nums[i]})
		m[nums[i]]++
	}

	pq := newPQ(init)
	n := pq.top()
	ret = append(ret, n.Weight)
	l := 0
	for r := k; r < len(nums); r++ {
		m[nums[r]]++
		m[nums[l]]--
		l++
		pq.push(&node{strconv.Itoa(r), nums[r]})
		for m[pq.top().Weight] <= 0 {
			pq.poll()
		}

		ret = append(ret, pq.top().Weight)

	}

	return ret
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

type node struct {
	Key    string
	Weight int
}

type maxPriorityQueue struct {
	heap    []*node
	indexes map[string]int // key -> index
}

func newPQ(list []node) *maxPriorityQueue {
	pq := &maxPriorityQueue{
		heap:    make([]*node, len(list)+1),
		indexes: make(map[string]int, len(list)),
	}

	for i := 0; i < len(list); i++ {
		pq.heap[i+1] = &list[i]
	}

	for i := len(pq.heap) / 2; i >= 1; i-- {
		pq.heapify(i)
	}

	return pq
}

func (pq *maxPriorityQueue) build() {
	for i := len(pq.heap) / 2; i >= 1; i-- {
		pq.heapify(i)
	}
}

func (pq *maxPriorityQueue) heapify(index int) {
	length := pq.getLength()

	minIndex := index
	if 2*index < length && pq.heap[minIndex].Weight < pq.heap[2*index].Weight {
		minIndex = 2 * index
	}

	if 2*index+1 < length && pq.heap[minIndex].Weight < pq.heap[2*index+1].Weight {
		minIndex = 2*index + 1
	}

	pq.swap(minIndex, index)

	if minIndex != index {
		pq.heapify(minIndex)
	}
}

func (pq *maxPriorityQueue) getLength() int {
	return len(pq.heap)
}

func (pq *maxPriorityQueue) swap(idx1 int, idx2 int) {
	pq.heap[idx1], pq.heap[idx2] = pq.heap[idx2], pq.heap[idx1]
}

func (pq *maxPriorityQueue) poll() *node {
	minIndex := 1
	minNode := pq.heap[minIndex]

	pq.heap[minIndex] = pq.heap[len(pq.heap)-1]

	pq.heap = pq.heap[:len(pq.heap)-1]

	pq.heapify(minIndex)

	return minNode
}

func (pq *maxPriorityQueue) top() *node {
	return pq.heap[1]
}

func (pq *maxPriorityQueue) push(n *node) {
	pq.heap = append(pq.heap, n)
	index := len(pq.heap) - 1
	for index > 1 && pq.getParent(index).Weight < pq.heap[index].Weight {
		pq.swap(index/2, index)
		index /= 2
	}
}

func (pq *maxPriorityQueue) increase(index int) {
}

func (pq *maxPriorityQueue) getParent(idx int) *node {
	return pq.heap[idx/2]
}
