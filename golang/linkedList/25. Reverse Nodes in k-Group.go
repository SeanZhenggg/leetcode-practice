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

// iterative 1 solution
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
// start: 0, first: 1, end: 3
func reverse(begin *ListNode, end *ListNode) *ListNode {
	first := begin.Next
	cur := first
	prev := end

	for cur != end {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	begin.Next = prev
	return first
}

// recursive 1 solution
func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	i := 0
	cur := head
	for ; cur != nil && i < k; i++ {
		cur = cur.Next
	}

	if i == k {
		// after returning from next recursion, head = 3, cur = 5
		cur = reverseKGroup2(cur, k)

		for i = 0; i < k; i++ {
			next := head.Next
			head.Next = cur
			cur = head
			head = next
		}
		return cur
	} else {
		return head
	}
}

// recursive 2 solution
func reverseKGroup3(head *ListNode, k int) *ListNode {
	curr := head
	i := 0
	for ; i < k && curr != nil; i++ {
		curr = curr.Next
	}

	if i%k == 0 {
		revHead := reverse2(head, k)
		head.Next = reverseKGroup3(curr, k)
		return revHead
	} else {
		return head
	}
}

// iterative 2 solution
func reverseKGroup4(head *ListNode, k int) *ListNode {
	var (
		newHead *ListNode = nil
		kTail   *ListNode = nil
	)
	curr := head

	for curr != nil {
		i := 0
		for ; i < k && curr != nil; i++ {
			curr = curr.Next
		}
		if i == k {
			revHead := reverse2(head, k)
			if newHead == nil {
				newHead = revHead
			}
			if kTail != nil {
				kTail.Next = revHead
			}
			kTail = head
			head = curr
		} else {
			if kTail != nil {
				kTail.Next = head
			}
		}
	}

	if newHead == nil {
		return head
	}
	return newHead
}

func reverse2(head *ListNode, k int) *ListNode {
	var prev *ListNode = nil
	cur := head
	for k > 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		k--
	}
	return prev
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

func Test_reverseKGroup3() {
	head15 := &ListNode{Val: 5, Next: nil}
	head14 := &ListNode{Val: 4, Next: head15}
	head13 := &ListNode{Val: 3, Next: head14}
	head12 := &ListNode{Val: 2, Next: head13}
	head11 := &ListNode{Val: 1, Next: head12}

	ans1 := reverseKGroup3(head11, 2)
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
	ans2 := reverseKGroup3(head21, 3)
	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)
}

func Test_reverseKGroup4() {
	head15 := &ListNode{Val: 5, Next: nil}
	head14 := &ListNode{Val: 4, Next: head15}
	head13 := &ListNode{Val: 3, Next: head14}
	head12 := &ListNode{Val: 2, Next: head13}
	head11 := &ListNode{Val: 1, Next: head12}

	ans1 := reverseKGroup4(head11, 2)
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
	ans2 := reverseKGroup4(head21, 3)
	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)
}
