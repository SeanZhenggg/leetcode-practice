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
func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lHeight := dfs(root.Left)
		if root.Left != nil {
			lHeight += 1
		}
		rHeight := dfs(root.Right)
		if root.Right != nil {
			rHeight += 1
		}

		maxLength = max(maxLength, lHeight+rHeight)
		return max(lHeight, rHeight)
	}

	dfs(root)
	return maxLength
}

func diameterOfBinaryTree2(root *TreeNode) int {
	maxLength := 0
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lHeight := dfs(root.Left)
		rHeight := dfs(root.Right)

		maxLength = max(maxLength, lHeight+rHeight)
		return max(lHeight, rHeight) + 1 // highlights
	}

	dfs(root)
	return maxLength
}

func Test_DiameterOfBinaryTree() {
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}

	ans1 := diameterOfBinaryTree(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	ans2 := diameterOfBinaryTree(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 8}}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 10}}}},
		Right: &TreeNode{Val: 3},
	}

	ans3 := diameterOfBinaryTree(root3)
	log.Printf("ans3: %v", ans3)

}
