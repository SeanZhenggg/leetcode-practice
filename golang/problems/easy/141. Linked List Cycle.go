package easy

import (
	"log"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if slow != nil && fast != nil && slow == fast {
			return true
		}
	}

	return false
}

func Test_hasCycle() {
	case1H4 := &ListNode{-4, nil}
	case1H3 := &ListNode{0, case1H4}
	case1H2 := &ListNode{2, case1H3}
	case1H1 := &ListNode{3, case1H2}
	case1H4.Next = case1H2
	ans1 := hasCycle(case1H1)
	log.Printf("ans1: %v", ans1)

	case2H2 := &ListNode{2, nil}
	case2H1 := &ListNode{1, case2H2}
	case2H2.Next = case2H1
	ans2 := hasCycle(case2H1)
	log.Printf("ans2: %v", ans2)

	case3H1 := &ListNode{1, nil}
	ans3 := hasCycle(case3H1)
	log.Printf("ans3: %v", ans3)
}
