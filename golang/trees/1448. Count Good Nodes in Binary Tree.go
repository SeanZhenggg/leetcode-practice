package trees

import (
	"log"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	count := 0
	var dfs func(root *TreeNode, maxValFromRoot int)
	dfs = func(root *TreeNode, maxValFromRoot int) {
		if root == nil {
			return
		}
		if root.Val >= maxValFromRoot {
			count++
			maxValFromRoot = root.Val
		}

		dfs(root.Left, maxValFromRoot)
		dfs(root.Right, maxValFromRoot)
	}

	dfs(root, math.MinInt16)

	return count
}

func Test_GoodNodes() {
	root1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}},
	}
	ans1 := goodNodes(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}},
	}
	ans2 := goodNodes(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{Val: 1}
	ans3 := goodNodes(root3)
	log.Printf("ans3: %v", ans3)

}
