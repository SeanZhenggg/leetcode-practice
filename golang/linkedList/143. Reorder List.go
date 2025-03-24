package linkedList

import (
	"fmt"
	"log"
	"strings"
)

// hash map solution - O(n), O(n)
func reorderList(head *ListNode) {
	m := make(map[int]*ListNode)
	cur := head
	idx := 0
	for cur != nil {
		m[idx] = cur
		cur = cur.Next
		idx++
	}

	l, r := 0, len(m)-1

	for l <= r {
		m[l].Next = m[r]
		m[r].Next = m[l]
		l++
		r--
	}

	cur.Next = nil
}

// reverse half and merge solution - O(n), O(1)
func reorderList2(head *ListNode) {
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	node := slow.Next
	slow.Next = nil

	var rev *ListNode = nil
	for node != nil {
		next := node.Next
		node.Next = rev
		rev = node
		node = next
	}

	slow = head
	for slow != nil && rev != nil {
		next := slow.Next
		rNext := rev.Next

		slow.Next = rev
		rev.Next = next

		slow = next
		rev = rNext
	}
}

// recursive solution
func reorderList3(head *ListNode) {
	var reverse func(current *ListNode) *ListNode
	root := head
	reverse = func(current *ListNode) *ListNode {
		if current == nil {
			return root
		}

		root = reverse(current.Next)
		if root == nil {
			return nil
		}
		if root == current || root.Next == current {
			current.Next = nil
			return nil
		} else {
			temp := root.Next
			root.Next = current
			current.Next = temp
			return temp
		}
	}

	reverse(root)
}

func Test_reorderList() {
	case1H4 := &ListNode{4, nil}
	case1H3 := &ListNode{3, case1H4}
	case1H2 := &ListNode{2, case1H3}
	case1H1 := &ListNode{1, case1H2}

	reorderList2(case1H1)
	ans1Str := ""
	for case1H1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", case1H1.Val)
		case1H1 = case1H1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	case2H5 := &ListNode{5, nil}
	case2H4 := &ListNode{4, case2H5}
	case2H3 := &ListNode{3, case2H4}
	case2H2 := &ListNode{2, case2H3}
	case2H1 := &ListNode{1, case2H2}

	reorderList2(case2H1)
	ans2Str := ""
	for case2H1 != nil {
		ans2Str += fmt.Sprintf("%d -> ", case2H1.Val)
		case2H1 = case2H1.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)
}

func Test_reorderList2() {
	case1H4 := &ListNode{4, nil}
	case1H3 := &ListNode{3, case1H4}
	case1H2 := &ListNode{2, case1H3}
	case1H1 := &ListNode{1, case1H2}

	reorderList2(case1H1)
	ans1Str := ""
	for case1H1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", case1H1.Val)
		case1H1 = case1H1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	case2H5 := &ListNode{5, nil}
	case2H4 := &ListNode{4, case2H5}
	case2H3 := &ListNode{3, case2H4}
	case2H2 := &ListNode{2, case2H3}
	case2H1 := &ListNode{1, case2H2}

	reorderList2(case2H1)
	ans2Str := ""
	for case2H1 != nil {
		ans2Str += fmt.Sprintf("%d -> ", case2H1.Val)
		case2H1 = case2H1.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	case3H1 := &ListNode{1, nil}

	reorderList2(case3H1)
	ans3Str := ""
	for case3H1 != nil {
		ans3Str += fmt.Sprintf("%d -> ", case3H1.Val)
		case3H1 = case3H1.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)
}

func Test_reorderList3() {
	case1H4 := &ListNode{4, nil}
	case1H3 := &ListNode{3, case1H4}
	case1H2 := &ListNode{2, case1H3}
	case1H1 := &ListNode{1, case1H2}

	reorderList3(case1H1)
	ans1Str := ""
	for case1H1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", case1H1.Val)
		case1H1 = case1H1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	case2H5 := &ListNode{5, nil}
	case2H4 := &ListNode{4, case2H5}
	case2H3 := &ListNode{3, case2H4}
	case2H2 := &ListNode{2, case2H3}
	case2H1 := &ListNode{1, case2H2}

	reorderList3(case2H1)
	ans2Str := ""
	for case2H1 != nil {
		ans2Str += fmt.Sprintf("%d -> ", case2H1.Val)
		case2H1 = case2H1.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)
}
