package medium

import (
	"fmt"
	"log"
	"strings"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// hash map solution - O(n), O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node)

	newNode := new(Node)
	oldCur := head
	newCur := newNode
	for oldCur != nil {
		newN := &Node{Val: oldCur.Val}
		m[oldCur] = newN
		newCur.Next = newN
		oldCur = oldCur.Next
		newCur = newCur.Next
	}

	oldCur = head
	newCur = newNode.Next
	for oldCur != nil {
		if newN, found := m[oldCur.Random]; found {
			newCur.Random = newN
		}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}

	return newNode.Next
}

func Test_copyRandomList() {
	// case 1
	case1H5 := &Node{Val: 1, Next: nil}
	case1H4 := &Node{Val: 10, Next: case1H5}
	case1H3 := &Node{Val: 11, Next: case1H4}
	case1H2 := &Node{Val: 13, Next: case1H3}
	case1H1 := &Node{Val: 7, Next: case1H2}

	case1H2.Random = case1H1
	case1H3.Random = case1H5
	case1H4.Random = case1H3
	case1H5.Random = case1H1

	ans1 := copyRandomList(case1H1)
	ans1Str := ""
	curAns1 := ans1
	for curAns1 != nil {
		ans1Str += fmt.Sprintf("%d -> ", curAns1.Val)
		curAns1 = curAns1.Next
	}
	ans1Str = strings.TrimRight(ans1Str, " -> ")
	log.Printf(ans1Str)

	curAns1 = ans1
	for curAns1 != nil {
		if curAns1.Random != nil {
			log.Printf("Node %d Random -> %d", curAns1.Val, curAns1.Random.Val)
		} else {
			log.Printf("Node %d Random -> nil", curAns1.Val)
		}
		curAns1 = curAns1.Next
	}

	// case 2
	case2H2 := &Node{Val: 2, Next: nil}
	case2H1 := &Node{Val: 1, Next: case2H2}

	case2H1.Random = case2H2
	case2H2.Random = case2H2

	ans2 := copyRandomList(case2H1)
	ans2Str := ""
	curAns2 := ans2
	for curAns2 != nil {
		ans2Str += fmt.Sprintf("%d -> ", curAns2.Val)
		curAns2 = curAns2.Next
	}
	ans2Str = strings.TrimRight(ans2Str, " -> ")
	log.Printf(ans2Str)

	curAns2 = ans2
	for curAns2 != nil {
		if curAns2.Random != nil {
			log.Printf("Node %d Random -> %d", curAns2.Val, curAns2.Random.Val)
		} else {
			log.Printf("Node %d Random -> nil", curAns2.Val)
		}
		curAns2 = curAns2.Next
	}

	// case 3
	case3H3 := &Node{Val: 3, Next: nil}
	case3H2 := &Node{Val: 3, Next: case3H3}
	case3H1 := &Node{Val: 3, Next: case3H2}

	case3H2.Random = case3H1

	ans3 := copyRandomList(case3H1)
	ans3Str := ""
	curAns3 := ans3
	for curAns3 != nil {
		ans3Str += fmt.Sprintf("%d -> ", curAns3.Val)
		curAns3 = curAns3.Next
	}
	ans3Str = strings.TrimRight(ans3Str, " -> ")
	log.Printf(ans3Str)

	curAns3 = ans3
	for curAns3 != nil {
		if curAns3.Random != nil {
			log.Printf("Node %d Random -> %d", curAns3.Val, curAns3.Random.Val)
		} else {
			log.Printf("Node %d Random -> nil", curAns3.Val)
		}
		curAns3 = curAns3.Next
	}
}
