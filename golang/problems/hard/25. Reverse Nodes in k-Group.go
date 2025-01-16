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

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	begin := dummy
	for i := 1; head != nil; i += 1 {
		if i%k == 0 {
			begin = reverse(begin, head.Next)
			head = begin.Next
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

// 0 -> 1 -> 2 -> 3
// 0 = start
// 1 = first
// 3 = end
func reverse(start *ListNode, end *ListNode) *ListNode {
	first := start.Next
	cur := first
	prev := end

	for cur != end {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	start.Next = prev
	return first
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	i := 0
	cur := head
	for i = 0; cur != nil && i < k; i++ {
		cur = cur.Next
	}

	if i == k {
		// nextHead = 5, head = 3
		nextHead := reverseKGroup2(cur, k)

		prev := nextHead
		for i = 0; i < k; i++ {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}
		return prev
	} else {
		return head
	}

}

func Test_reverseKGroup() {
	head15 := &ListNode{Val: 5, Next: nil}
	head14 := &ListNode{Val: 4, Next: head15}
	head13 := &ListNode{Val: 3, Next: head14}
	head12 := &ListNode{Val: 2, Next: head13}
	head11 := &ListNode{Val: 1, Next: head12}

	ans1 := reverseKGroup(head11, 2)
	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	head25 := &ListNode{Val: 5, Next: nil}
	head24 := &ListNode{Val: 4, Next: head25}
	head23 := &ListNode{Val: 3, Next: head24}
	head22 := &ListNode{Val: 2, Next: head23}
	head21 := &ListNode{Val: 1, Next: head22}
	ans2 := reverseKGroup(head21, 3)
	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)
}

func Test_reverseKGroup2() {
	head15 := &ListNode{Val: 5, Next: nil}
	head14 := &ListNode{Val: 4, Next: head15}
	head13 := &ListNode{Val: 3, Next: head14}
	head12 := &ListNode{Val: 2, Next: head13}
	head11 := &ListNode{Val: 1, Next: head12}

	ans1 := reverseKGroup2(head11, 2)
	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	head25 := &ListNode{Val: 5, Next: nil}
	head24 := &ListNode{Val: 4, Next: head25}
	head23 := &ListNode{Val: 3, Next: head24}
	head22 := &ListNode{Val: 2, Next: head23}
	head21 := &ListNode{Val: 1, Next: head22}
	ans2 := reverseKGroup2(head21, 3)
	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)
}
