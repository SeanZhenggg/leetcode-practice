package stack

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

func Constructor() MinStack2 {
	return MinStack2{
		st:    make([]int, 0),
		minSt: make([]int, 0),
	}
}

func (this *MinStack2) Push(val int) {
	this.st = append(this.st, val)
	if len(this.minSt) != 0 {
		top := this.minSt[len(this.minSt)-1]
		this.minSt = append(this.minSt, min(top, val))
	} else {
		this.minSt = append(this.minSt, val)
	}
}

func (this *MinStack2) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.minSt = this.minSt[:len(this.minSt)-1]
}

func (this *MinStack2) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack2) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}

type MinStackFailed struct {
	st    []int
	minSt []int
}

func ConstructorFailed() MinStackFailed {
	return MinStackFailed{
		st:    make([]int, 0),
		minSt: make([]int, 0),
	}
}

func (this *MinStackFailed) Push(val int) {
	this.st = append(this.st, val)
	if len(this.minSt) != 0 {
		top := this.minSt[len(this.minSt)-1]
		if top >= val {
			this.minSt = append(this.minSt, val)
		} else {
			this.minSt = this.minSt[:len(this.minSt)-1]
			this.minSt = append(this.minSt, val)
			this.minSt = append(this.minSt, top)
		}
	} else {
		this.minSt = append(this.minSt, val)
	}
}

func (this *MinStackFailed) Pop() {
	top := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	if len(this.minSt) != 0 {
		if top == this.minSt[len(this.minSt)-1] {
			this.minSt = this.minSt[:len(this.minSt)-1]
		}
	}
}

func (this *MinStackFailed) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStackFailed) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}
