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

func diameterOfBinaryTreeReview1(root *TreeNode) int {
	maxLen := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		l := getMaxReview(root.Left)
		r := getMaxReview(root.Right)
		maxLen = max(maxLen, l+r)

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return maxLen
}

func getMaxReview(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getMaxReview(root.Left), getMaxReview(root.Right))
}

func diameterOfBinaryTreeReview2(root *TreeNode) int {
	maxLen := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := dfs(root.Left)
		r := dfs(root.Right)
		maxLen = max(maxLen, l+r)
		return 1 + max(l, r)
	}

	dfs(root)
	return maxLen
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

func Test_DiameterOfBinaryTreeReview1() {
	root1 := generateTree([]string{"1", "2", "3", "4", "5"})
	ans1 := diameterOfBinaryTreeReview1(root1)
	log.Printf("ans1: %v", ans1)

	root2 := generateTree([]string{"1", "2"})
	ans2 := diameterOfBinaryTreeReview1(root2)
	log.Printf("ans2: %v", ans2)

	root3 := generateTree([]string{"1", "2", "3", "4", "5", "null", "null", "7", "null", "9", "null", "8", "null", "10"})
	ans3 := diameterOfBinaryTreeReview1(root3)
	log.Printf("ans3: %v", ans3)

}

func Test_DiameterOfBinaryTreeReview2() {
	root1 := generateTree([]string{"1", "2", "3", "4", "5"})
	ans1 := diameterOfBinaryTreeReview2(root1)
	log.Printf("ans1: %v", ans1)

	root2 := generateTree([]string{"1", "2"})
	ans2 := diameterOfBinaryTreeReview2(root2)
	log.Printf("ans2: %v", ans2)

	root3 := generateTree([]string{"1", "2", "3", "4", "5", "null", "null", "7", "null", "9", "null", "8", "null", "10"})
	ans3 := diameterOfBinaryTreeReview2(root3)
	log.Printf("ans3: %v", ans3)

}
