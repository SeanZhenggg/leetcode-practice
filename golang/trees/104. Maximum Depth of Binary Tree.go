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

func maxDepthReview1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepthReview1(root.Left), maxDepthReview1(root.Right))
}

func maxDepthReview2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	lvl := 0
	for len(queue) > 0 {
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			top := queue[0]
			queue = queue[1:]
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		lvl++
	}

	return lvl
}

func Test_MaxDepth() {
	root1 := generateTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	ans1 := maxDepth(root1)
	log.Printf("ans1: %v", ans1)

	root2 := generateTree([]string{"1", "null", "2"})
	ans2 := maxDepth(root2)
	log.Printf("ans2: %v", ans2)
}

func Test_MaxDepth2() {
	root1 := generateTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	ans1 := maxDepth2(root1)
	log.Printf("ans1: %v", ans1)

	root2 := generateTree([]string{"1", "null", "2"})
	ans2 := maxDepth2(root2)
	log.Printf("ans2: %v", ans2)
}

func Test_MaxDepthReview1() {
	root1 := generateTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	ans1 := maxDepthReview1(root1)
	log.Printf("ans1: %v", ans1)

	root2 := generateTree([]string{"1", "null", "2"})
	ans2 := maxDepthReview1(root2)
	log.Printf("ans2: %v", ans2)
}

func Test_MaxDepthReview2() {
	root1 := generateTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	ans1 := maxDepthReview2(root1)
	log.Printf("ans1: %v", ans1)

	root2 := generateTree([]string{"1", "null", "2"})
	ans2 := maxDepthReview2(root2)
	log.Printf("ans2: %v", ans2)
}
