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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

type levelNode struct {
	node  *TreeNode
	level int
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDep := 0
	queue := make([]levelNode, 0)
	queue = append(queue, levelNode{node: root, level: 1})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.node != nil {
			maxDep = max(maxDep, top.level)

			if top.node.Left != nil {
				queue = append(queue, levelNode{node: top.node.Left, level: top.level + 1})
			}
			if top.node.Right != nil {
				queue = append(queue, levelNode{node: top.node.Right, level: top.level + 1})
			}
		}
	}
	return maxDep
}

func Test_MaxDepth() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := maxDepth(root1)
	log.Printf("ans1: %v", ans1)
	root2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	ans2 := maxDepth(root2)
	log.Printf("ans2: %v", ans2)
}

func Test_MaxDepth2() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := maxDepth2(root1)
	log.Printf("ans1: %v", ans1)
	root2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	ans2 := maxDepth2(root2)
	log.Printf("ans2: %v", ans2)
}
