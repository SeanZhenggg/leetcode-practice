package linkedList

import (
	"log"
)

type DoublyLinkedListNode struct {
	Key  int
	Val  int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

type LRUCache struct {
	Capacity int
	Map      map[int]*DoublyLinkedListNode
	Head     *DoublyLinkedListNode
	Tail     *DoublyLinkedListNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	head := &DoublyLinkedListNode{}
	tail := &DoublyLinkedListNode{}

	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*DoublyLinkedListNode, capacity),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, found := this.Map[key]
	if !found {
		return -1
	}

	this.removeNode(node)
	this.addNode(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, found := this.Map[key]
	if found {
		node.Val = value
		this.removeNode(node)
	} else {
		node = &DoublyLinkedListNode{Key: key, Val: value}
		this.Map[key] = node
	}
	this.addNode(node)

	if len(this.Map) > this.Capacity {
		prevNode := this.Tail.Prev
		this.removeNode(prevNode)
		delete(this.Map, prevNode.Key)
	}
}

func (this *LRUCache) addNode(node *DoublyLinkedListNode) {
	nextNode := this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
	node.Next = nextNode
	nextNode.Prev = node
}

func (this *LRUCache) removeNode(node *DoublyLinkedListNode) {
	prevNode := node.Prev
	nextNode := node.Next
	prevNode.Next = nextNode
	nextNode.Prev = prevNode
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
