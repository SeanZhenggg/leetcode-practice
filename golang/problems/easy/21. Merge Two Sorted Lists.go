package easy

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp := list1.Next
			list1.Next = nil
			head.Next = list1
			list1 = temp
		} else {
			temp := list2.Next
			list2.Next = nil
			head.Next = list2
			list2 = temp
		}

		head = head.Next
	}

	for list1 != nil {
		head.Next = list1
		list1 = list1.Next
		head = head.Next
	}

	for list2 != nil {
		head.Next = list2
		list2 = list2.Next
		head = head.Next
	}

	return dummyHead.Next
}

func Test_mergeTwoLists() {
	case1L1H3 := &ListNode{3, nil}
	case1L1H2 := &ListNode{2, case1L1H3}
	case1L1H1 := &ListNode{1, case1L1H2}
	case1L2H3 := &ListNode{4, nil}
	case1L2H2 := &ListNode{3, case1L2H3}
	case1L2H1 := &ListNode{1, case1L2H2}
	ans1 := mergeTwoLists(case1L1H1, case1L2H1)

	ans1Str := ""
	for ans1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", ans1.Val)
		ans1 = ans1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	ans2 := mergeTwoLists(nil, nil)

	ans2Str := ""
	for ans2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", ans2.Val)
		ans2 = ans2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	ans3 := mergeTwoLists(&ListNode{Val: 0}, nil)

	ans3Str := ""
	for ans3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", ans3.Val)
		ans3 = ans3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}
