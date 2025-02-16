package trees

import "log"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	queue := make([]struct {
		Node  *TreeNode
		Depth int
	}, 0)
	queue = append(queue, struct {
		Node  *TreeNode
		Depth int
	}{
		Node:  root,
		Depth: 1,
	})

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if len(ret) < first.Depth {
			ret = append(ret, make([]int, 0))
		}
		ret[first.Depth-1] = append(ret[first.Depth-1], first.Node.Val)

		if first.Node.Left != nil {
			queue = append(queue, struct {
				Node  *TreeNode
				Depth int
			}{
				Node:  first.Node.Left,
				Depth: first.Depth + 1,
			})
		}
		if first.Node.Right != nil {
			queue = append(queue, struct {
				Node  *TreeNode
				Depth int
			}{
				Node:  first.Node.Right,
				Depth: first.Depth + 1,
			})
		}
	}
	return ret
}

func levelOrderReview1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	ret := make([][]int, 0)
	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			top := queue[0]
			queue = queue[1:]

			if level == len(ret) {
				ret = append(ret, make([]int, 0))
			}
			ret[level] = append(ret[level], top.Val)

			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		level++
	}

	return ret
}

func Test_LevelOrder() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := levelOrder(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: 1}
	ans2 := levelOrder(root2)
	log.Printf("ans2: %v", ans2)

	ans3 := levelOrder(nil)
	log.Printf("ans3: %v", ans3)
}

func Test_LevelOrderReview1() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := levelOrderReview1(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: 1}
	ans2 := levelOrderReview1(root2)
	log.Printf("ans2: %v", ans2)

	ans3 := levelOrderReview1(nil)
	log.Printf("ans3: %v", ans3)
}
