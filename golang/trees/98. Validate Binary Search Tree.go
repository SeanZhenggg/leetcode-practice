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

type info struct {
	Min   int
	Max   int
	Valid bool
}

func isValidBST(root *TreeNode) bool {
	// check each node if it's valid BST
	// check left subtree
	// check right subtree
	var dfs func(root *TreeNode, a int, b int) *info
	dfs = func(root *TreeNode, a int, b int) *info {
		if root == nil {
			return &info{a, b, true}
		}

		lInfo := dfs(root.Left, a, b)
		rInfo := dfs(root.Right, a, b)

		if root.Val < lInfo.Max || root.Val > rInfo.Min {
			return &info{min(lInfo.Min, rInfo.Min, root.Val), max(lInfo.Max, rInfo.Max, b), false}
		}

		return &info{min(root.Val, lInfo.Min, rInfo.Min), max(root.Val, lInfo.Max, rInfo.Max), lInfo.Valid && rInfo.Valid}
	}

	ans := dfs(root, math.MaxInt16, math.MinInt16)
	return ans.Valid
}

//func isValidBST(root *TreeNode) bool {
//	// check each node if it's valid BST
//	// check left subtree
//	// check right subtree
//
//	if root == nil {
//		return true
//	}
//
//	lValid := isValidBST(root.Left)
//	rValid := isValidBST(root.Right)
//
//	if (root.Left != nil && root.Val <= root.Left.Val) || (root.Right != nil && root.Val >= root.Right.Val) {
//		return false
//	}
//
//	return lValid && rValid
//}

func Test_IsValidBST() {
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	ans1 := isValidBST(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	ans2 := isValidBST(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 8}},
	}
	ans3 := isValidBST(root3)
	log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}
	ans4 := isValidBST(root4)
	log.Printf("ans4: %v", ans4)
}
