package medium

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	jing := 0
	cur := dummy
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + jing
		jing = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}

	for l1 != nil {
		sum := l1.Val + jing
		jing = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		sum := l2.Val + jing
		jing = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		l2 = l2.Next
		cur = cur.Next
	}
	if jing != 0 {
		cur.Next = &ListNode{Val: jing}
	}

	return dummy.Next
}

func Test_addTwoNumbers() {
	case1L1H3 := &ListNode{3, nil}
	case1L1H2 := &ListNode{4, case1L1H3}
	case1L1H1 := &ListNode{2, case1L1H2}

	case1L2H3 := &ListNode{4, nil}
	case1L2H2 := &ListNode{6, case1L2H3}
	case1L2H1 := &ListNode{5, case1L2H2}
	ans1 := addTwoNumbers(case1L1H1, case1L2H1)
	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	case2L1H1 := &ListNode{0, nil}

	case2L2H1 := &ListNode{0, nil}

	ans2 := addTwoNumbers(case2L1H1, case2L2H1)
	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	case3L1H7 := &ListNode{9, nil}
	case3L1H6 := &ListNode{9, case3L1H7}
	case3L1H5 := &ListNode{9, case3L1H6}
	case3L1H4 := &ListNode{9, case3L1H5}
	case3L1H3 := &ListNode{9, case3L1H4}
	case3L1H2 := &ListNode{9, case3L1H3}
	case3L1H1 := &ListNode{9, case3L1H2}

	case3L2H4 := &ListNode{9, nil}
	case3L2H3 := &ListNode{9, case3L2H4}
	case3L2H2 := &ListNode{9, case3L2H3}
	case3L2H1 := &ListNode{9, case3L2H2}

	ans3 := addTwoNumbers(case3L1H1, case3L2H1)
	ans3Str := ""
	for ans3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", ans3.Val)
		ans3 = ans3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}
