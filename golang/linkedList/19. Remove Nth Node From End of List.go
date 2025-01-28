package linkedList

import (
	"fmt"
	"log"
	"strings"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for n > 0 {
		fast = fast.Next
		n--
	}

	removed := head
	var prev *ListNode = nil
	for fast != nil {
		prev = removed
		removed = removed.Next
		fast = fast.Next
	}

	if removed == head {
		return head.Next
	}
	prev.Next = removed.Next
	return head
}

func Test_removeNthFromEnd() {
	case1H5 := &ListNode{5, nil}
	case1H4 := &ListNode{4, case1H5}
	case1H3 := &ListNode{3, case1H4}
	case1H2 := &ListNode{2, case1H3}
	case1H1 := &ListNode{1, case1H2}

	ans1 := removeNthFromEnd(case1H1, 5)
	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	case2H1 := &ListNode{1, nil}

	ans2 := removeNthFromEnd(case2H1, 1)
	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	case3H2 := &ListNode{2, nil}
	case3H1 := &ListNode{1, case3H2}

	ans3 := removeNthFromEnd(case3H1, 1)
	ans3Str := ""
	for ans3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", ans3.Val)
		ans3 = ans3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}
