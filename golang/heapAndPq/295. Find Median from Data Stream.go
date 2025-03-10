package heapAndPq

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"sort"
)

// sort solution, TC: O(n*log n), SC: O(n)
type MedianFinder struct {
	nums []int
}

func MedianFinderConstructor() MedianFinder {
	return MedianFinder{
		nums: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.nums)

	if len(this.nums)%2 == 0 {
		return float64(this.nums[len(this.nums)/2]+this.nums[(len(this.nums)-1)/2]) / 2.0
	} else {
		return float64(this.nums[len(this.nums)/2])
	}
}

// two heaps solution, TC: O(log n), SC: O(n)
type MedianFinder2 struct {
	lo *priorityqueue.Queue
	hi *priorityqueue.Queue
}

func MedianFinder2Constructor() MedianFinder2 {
	return MedianFinder2{
		lo: priorityqueue.NewWith(func(a, b interface{}) int {
			return b.(int) - a.(int)
		}), // max heap
		hi: priorityqueue.NewWith(utils.IntComparator), // min heap
	}
}

func (this *MedianFinder2) AddNum(num int) {
	this.lo.Enqueue(num)
	val, _ := this.lo.Dequeue()
	this.hi.Enqueue(val)
	if this.lo.Size() < this.hi.Size() {
		val, _ := this.hi.Dequeue()
		this.lo.Enqueue(val)
	}
}

func (this *MedianFinder2) FindMedian() float64 {
	if this.lo.Size() > this.hi.Size() {
		val, _ := this.lo.Peek()
		return float64(val.(int))
	} else {
		lVal, _ := this.lo.Peek()
		hVal, _ := this.hi.Peek()
		return float64(lVal.(int)+hVal.(int)) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
