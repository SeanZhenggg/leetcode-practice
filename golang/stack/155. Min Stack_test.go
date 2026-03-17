package stack

import (
	"log"
	"testing"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func TestMinStack(t *testing.T) {
	obj := MinStackConstructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
}

func TestMinStack2(t *testing.T) {
	//obj := Constructor()
	//obj.Push(-2)
	//obj.Push(0)
	//obj.Push(-3)
	//log.Println(obj.GetMin())
	//obj.Pop()
	//log.Println(obj.Top())
	//log.Println(obj.GetMin())

	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
	obj.Push(2)
	obj.Push(-4)
	obj.Push(3)
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
}

func TestMinStackFailed(t *testing.T) {
	//obj := Constructor()
	//obj.Push(-2)
	//obj.Push(0)
	//obj.Push(-3)
	//log.Println(obj.GetMin())
	//obj.Pop()
	//log.Println(obj.Top())
	//log.Println(obj.GetMin())

	obj := ConstructorFailed()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
	obj.Push(2)
	obj.Push(-4)
	obj.Push(3)
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
	obj.Pop()
	log.Println(obj.Top())
	log.Println(obj.GetMin())
}
