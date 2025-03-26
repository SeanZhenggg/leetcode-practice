package linkedList

import (
	"fmt"
	"log"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// merge one by one solution - O(?), O(?)
func mergeKLists(lists []*ListNode) *ListNode {
	var prev *ListNode = nil
	for _, node := range lists {
		prev = merge(prev, node)
	}
	return prev
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for list1 != nil && list2 != nil {
		var next *ListNode
		if list1.Val <= list2.Val {
			next = list1
			list1 = list1.Next
		} else {
			next = list2
			list2 = list2.Next
		}
		head.Next = next
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}

	return dummy.Next
}

// merge sort solution - O(?), O(?)
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeList(lists, 0, len(lists)-1)
}

// compare one by one solution - O(k*N), O(n)
func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}

	head := dummy

	for len(lists) > 0 {
		minIdx := 0
		for i := 1; i < len(lists); i++ {
			if lists[i] != nil && lists[minIdx] != nil && lists[i].Val < lists[minIdx].Val {
				minIdx = i
			}
		}

		if lists[minIdx] != nil {
			head.Next = &ListNode{Val: lists[minIdx].Val}
			head = head.Next
			lists[minIdx] = lists[minIdx].Next
		}
		if lists[minIdx] == nil {
			lists = append(lists[:minIdx], lists[minIdx+1:]...)
		}
	}

	return dummy.Next
}

func mergeList(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	mid := l + (r-l)/2

	left := mergeList(lists, l, mid)
	right := mergeList(lists, mid+1, r)

	return merge(left, right)
}

func Test_mergeKLists() {
	//[[1,4,5],[1,3,4],[2,6]]
	node113 := &ListNode{Val: 5, Next: nil}
	node112 := &ListNode{Val: 4, Next: node113}
	node111 := &ListNode{Val: 1, Next: node112}

	node123 := &ListNode{Val: 4, Next: nil}
	node122 := &ListNode{Val: 3, Next: node123}
	node121 := &ListNode{Val: 1, Next: node122}

	node132 := &ListNode{Val: 6, Next: nil}
	node131 := &ListNode{Val: 2, Next: node132}
	ans1 := mergeKLists([]*ListNode{node111, node121, node131})

	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	ans2 := mergeKLists([]*ListNode{})

	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	ans3 := mergeKLists([]*ListNode{nil})

	ans3Str := ""
	for ans3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", ans3.Val)
		ans3 = ans3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}

func Test_mergeKLists2() {
	//[[1,4,5],[1,3,4],[2,6]]
	node113 := &ListNode{Val: 5, Next: nil}
	node112 := &ListNode{Val: 4, Next: node113}
	node111 := &ListNode{Val: 1, Next: node112}

	node123 := &ListNode{Val: 4, Next: nil}
	node122 := &ListNode{Val: 3, Next: node123}
	node121 := &ListNode{Val: 1, Next: node122}

	node132 := &ListNode{Val: 6, Next: nil}
	node131 := &ListNode{Val: 2, Next: node132}
	ans1 := mergeKLists2([]*ListNode{node111, node121, node131})

	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	ans2 := mergeKLists2([]*ListNode{})

	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	ans3 := mergeKLists2([]*ListNode{nil})

	ans3Str := ""
	for ans3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", ans3.Val)
		ans3 = ans3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}

func Test_mergeKLists3() {
	//[[1,4,5],[1,3,4],[2,6]]
	node113 := &ListNode{Val: 5, Next: nil}
	node112 := &ListNode{Val: 4, Next: node113}
	node111 := &ListNode{Val: 1, Next: node112}

	node123 := &ListNode{Val: 4, Next: nil}
	node122 := &ListNode{Val: 3, Next: node123}
	node121 := &ListNode{Val: 1, Next: node122}

	node132 := &ListNode{Val: 6, Next: nil}
	node131 := &ListNode{Val: 2, Next: node132}
	ans1 := mergeKLists3([]*ListNode{node111, node121, node131})

	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	ans2 := mergeKLists3([]*ListNode{})

	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	ans3 := mergeKLists3([]*ListNode{nil, nil})

	ans3Str := ""
	for ans3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", ans3.Val)
		ans3 = ans3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}
