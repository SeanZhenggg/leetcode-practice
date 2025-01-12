package medium

import (
	"log"
)

type DoublePtrNode struct {
	Key  int
	Val  int
	Next *DoublePtrNode
	Prev *DoublePtrNode
}

type LRUCache struct {
	hashMap  map[int]*DoublePtrNode
	head     *DoublePtrNode
	tail     *DoublePtrNode
	capacity int
}

func LRUCacheConstructor(capacity int) LRUCache {
	head := &DoublePtrNode{}
	tail := &DoublePtrNode{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		hashMap:  make(map[int]*DoublePtrNode, capacity),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, found := this.hashMap[key]
	if !found {
		return -1
	}

	if node.Prev != this.head {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		next := this.head.Next
		this.head.Next = node
		node.Prev = this.head
		node.Next = next
		next.Prev = node
	}
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if node, found := this.hashMap[key]; found {
		node.Val = value
		this.Get(key)
		return
	}

	if len(this.hashMap) == this.capacity {
		lruNode := this.tail.Prev
		delete(this.hashMap, lruNode.Key)
		this.tail.Prev = lruNode.Prev
		lruNode.Prev.Next = this.tail
	}

	node := &DoublePtrNode{Key: key, Val: value}
	this.hashMap[key] = node
	next := this.head.Next
	this.head.Next = node
	node.Prev = this.head
	node.Next = next
	next.Prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Test_LRUCache() {
	//obj := LRUCacheConstructor(2)
	//
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//log.Println(obj.Get(1))
	//obj.Put(3, 3)
	//log.Println(obj.Get(2))
	//obj.Put(4, 4)
	//log.Println(obj.Get(1))
	//log.Println(obj.Get(3))
	//log.Println(obj.Get(4))

	obj2 := LRUCacheConstructor(2)

	obj2.Put(1, 0)
	obj2.Put(2, 2)
	log.Println(obj2.Get(1))
	obj2.Put(3, 3)
	log.Println(obj2.Get(2))
	obj2.Put(4, 4)
	log.Println(obj2.Get(1))
	log.Println(obj2.Get(3))
	log.Println(obj2.Get(4))
}
