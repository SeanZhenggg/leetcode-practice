package heapAndPq

import (
	"log"
)

type KthLargest struct {
	Nums []int
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	instance := KthLargest{
		Nums: nums,
		K:    k,
	}
	for i := (len(instance.Nums) - 1) / 2; i >= 0; i-- {
		instance.heapify(i, len(instance.Nums))
	}
	instance.sort()
	return instance
}

func (this *KthLargest) Add(val int) int {
	l, r := 0, len(this.Nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if this.Nums[mid] >= val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	a := make([]int, len(this.Nums)-l)
	copy(a, this.Nums[l:])
	this.Nums = append(this.Nums[:l], val)
	this.Nums = append(this.Nums, a...)

	return this.Nums[len(this.Nums)-this.K]
}

func (this *KthLargest) sort() {
	index := len(this.Nums) - 1
	top := 0

	for index > 0 {
		this.Nums[top], this.Nums[index] = this.Nums[index], this.Nums[top]
		this.heapify(top, index)
		index--
	}
}

func (this *KthLargest) add(val int) {
	this.Nums = append(this.Nums, val)
	index := len(this.Nums) - 1
	parent := (index - 1) / 2
	for index > 0 && this.Nums[index] > this.Nums[parent] {
		this.Nums[index], this.Nums[parent] = this.Nums[parent], this.Nums[index]
		index = parent
		parent /= 2
	}
}

func (this *KthLargest) heapify(index int, length int) {
	parent := index
	for index < length {
		left := 2*index + 1
		right := 2*index + 2

		if left < length && this.Nums[index] < this.Nums[left] {
			index = left
		}

		if right < length && this.Nums[index] < this.Nums[right] {
			index = right
		}

		if this.Nums[parent] == this.Nums[index] {
			break
		}

		this.Nums[parent], this.Nums[index] = this.Nums[index], this.Nums[parent]
		parent = index
	}
}

type KthLargest2 struct {
	Nums []int
	K    int
}

func Constructor2(k int, nums []int) KthLargest2 {
	instance := KthLargest2{
		Nums: nums,
		K:    k,
	}

	for i := (len(instance.Nums) - 2) / 2; i >= 0; i-- {
		instance.heapify(i)
	}

	for len(instance.Nums) > k {
		instance.delete()
	}
	return instance
}

func (this *KthLargest2) Add(val int) int {
	this.add(val)

	if len(this.Nums) > this.K {
		this.delete()
	}

	return this.Nums[0]
}

func (this *KthLargest2) add(value int) {
	this.Nums = append(this.Nums, value)
	index := len(this.Nums) - 1
	parent := (index - 1) / 2
	for index > 0 && this.Nums[index] < this.Nums[parent] {
		this.Nums[index], this.Nums[parent] = this.Nums[parent], this.Nums[index]
		index = parent
		parent = (parent - 1) / 2
	}
}

func (this *KthLargest2) heapify(index int) {
	parent := index
	for {
		left := 2*index + 1
		right := 2*index + 2

		if left < len(this.Nums) && this.Nums[index] > this.Nums[left] {
			index = left
		}

		if right < len(this.Nums) && this.Nums[index] > this.Nums[right] {
			index = right
		}

		if this.Nums[parent] == this.Nums[index] {
			break
		}

		this.Nums[parent], this.Nums[index] = this.Nums[index], this.Nums[parent]
		parent = index
	}
}

func (this *KthLargest2) delete() {
	this.Nums[0] = this.Nums[len(this.Nums)-1]
	this.Nums = this.Nums[:len(this.Nums)-1]

	this.heapify(0)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestKthLargest() {
	k1 := 3
	nums1 := []int{4, 5, 8, 2}
	obj1 := Constructor(k1, nums1)
	ans11 := obj1.Add(3)
	log.Println("ans1: ")
	log.Println(ans11)
	ans12 := obj1.Add(5)
	log.Println(ans12)
	ans13 := obj1.Add(10)
	log.Println(ans13)
	ans14 := obj1.Add(9)
	log.Println(ans14)
	ans15 := obj1.Add(4)
	log.Println(ans15)

	k2 := 4
	nums2 := []int{7, 7, 7, 7, 8, 3}
	obj2 := Constructor(k2, nums2)
	ans21 := obj2.Add(2)
	log.Println("ans2: ")
	log.Println(ans21)
	ans22 := obj2.Add(10)
	log.Println(ans22)
	ans23 := obj2.Add(9)
	log.Println(ans23)
	ans24 := obj2.Add(9)
	log.Println(ans24)
}

func TestKthLargest2() {
	k1 := 3
	nums1 := []int{4, 5, 8, 2}
	obj1 := Constructor2(k1, nums1)
	ans11 := obj1.Add(3)
	log.Println("ans1: ")
	log.Println(ans11)
	ans12 := obj1.Add(5)
	log.Println(ans12)
	ans13 := obj1.Add(10)
	log.Println(ans13)
	ans14 := obj1.Add(9)
	log.Println(ans14)
	ans15 := obj1.Add(4)
	log.Println(ans15)

	k2 := 4
	nums2 := []int{7, 7, 7, 7, 8, 3}
	obj2 := Constructor2(k2, nums2)
	ans21 := obj2.Add(2)
	log.Println("ans2: ")
	log.Println(ans21)
	ans22 := obj2.Add(10)
	log.Println(ans22)
	ans23 := obj2.Add(9)
	log.Println(ans23)
	ans24 := obj2.Add(9)
	log.Println(ans24)
}
