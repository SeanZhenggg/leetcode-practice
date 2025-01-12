package hard

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

// first merge solution - O(?), O(?)
func mergeKLists(lists []*ListNode) *ListNode {
	var prev *ListNode = nil
	for _, node := range lists {
		prev = merge1(prev, node)
	}
	return prev
}

func merge1(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

// merge sort solution - O(?), O(?)
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeList2(lists, 0, len(lists)-1)
}

func mergeList2(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := left + (right-left)/2

	lNode := mergeList2(lists, left, mid)
	rNode := mergeList2(lists, mid+1, right)
	return merge2(lNode, rNode)
}

func merge2(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
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
