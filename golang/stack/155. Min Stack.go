package stack

import "log"

type MinStack struct {
	st    []int
	minSt []int
}

func MinStackConstructor() MinStack {
	return MinStack{
		st:    make([]int, 0),
		minSt: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)

	if len(this.minSt) == 0 || this.minSt[len(this.minSt)-1] >= val {
		this.minSt = append(this.minSt, val)
	}
}

func (this *MinStack) Pop() {
	popEl := this.st[len(this.st)-1]
	if this.minSt[len(this.minSt)-1] == popEl {
		this.minSt = this.minSt[:len(this.minSt)-1]
	}
	this.st = this.st[:len(this.st)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}

type MinStack2 struct {
	st    []int
	minSt []int
}

func Constructor2() MinStack {
	return MinStack{
		st:    make([]int, 0),
		minSt: make([]int, 0),
	}
}

func (this *MinStack) Push2(val int) {
	this.st = append(this.st, val)

	if len(this.minSt) == 0 {

	} else {
		if this.minSt[len(this.minSt)-1] >= val {
			this.minSt = append(this.minSt, val)
		} else {
			this.minSt = append(this.minSt, this.minSt[len(this.minSt)-1])
		}
	}
}

func (this *MinStack) Pop2() {
	this.st = this.st[:len(this.st)-1]
	this.minSt = this.minSt[:len(this.minSt)-1]
}

func (this *MinStack) Top2() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin2() int {
	return this.minSt[len(this.minSt)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func Test_MinStack() {
	obj := MinStackConstructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
}
