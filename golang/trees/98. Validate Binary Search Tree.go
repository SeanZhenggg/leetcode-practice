package trees

import (
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

// failed
//func isValidBST(root *TreeNode) bool {
//	// check each node if it's valid BST
//	// check left subtree
//	// check right subtree
//	var dfs func(root *TreeNode, a int, b int) *info
//	dfs = func(root *TreeNode, a int, b int) *info {
//		if root == nil {
//			return &info{a, b, true}
//		}
//
//		lInfo := dfs(root.Left, a, b)
//		rInfo := dfs(root.Right, a, b)
//
//		if root.Val < lInfo.Max || root.Val > rInfo.Min {
//			return &info{min(lInfo.Min, rInfo.Min, root.Val), max(lInfo.Max, rInfo.Max, b), false}
//		}
//
//		return &info{min(root.Val, lInfo.Min, rInfo.Min), max(root.Val, lInfo.Max, rInfo.Max), lInfo.Valid && rInfo.Valid}
//	}
//
//	ans := dfs(root, math.MaxInt16, math.MinInt16)
//	return ans.Valid
//}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValid(root.Left, root.Val, func(val, limit int) bool { return val < limit }) ||
		!isValid(root.Right, root.Val, func(val, limit int) bool { return val > limit }) {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func isValid(root *TreeNode, limit int, check func(a int, b int) bool) bool {
	if root == nil {
		return true
	}
	if !check(root.Val, limit) {
		return false
	}
	return isValid(root.Left, limit, check) && isValid(root.Right, limit, check)
}

func isValidBST2(root *TreeNode) bool {
	return isValid2(root, math.MinInt64, math.MaxInt64)
}

func isValid2(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val < lower || root.Val > upper {
		return false
	}

	return isValid2(root.Left, lower, root.Val) && isValid2(root.Right, root.Val, upper)
}

func isValidBST3(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val < minVal || root.Val > maxVal {
		return false
	}

	return valid(root.Left, minVal, root.Val) && valid(root.Right, root.Val, maxVal)
}
