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
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return -1001
	}

	total := root.Val
	l := dfs(root.Left)
	if isValCount(l) {
		total = max(total, total+l)
	}
	r := dfs(root.Right)
	if isValCount(r) {
		total = max(total, total+r)
	}

	leftPathSum := maxPathSum(root.Left)
	if isValCount(leftPathSum) {
		total = max(total, leftPathSum)
	}
	rightPathSum := maxPathSum(root.Right)
	if isValCount(rightPathSum) {
		total = max(total, rightPathSum)
	}
	return total
}

func dfs(root *TreeNode) int {
	if root == nil {
		return -1001
	}

	l := dfs(root.Left)
	lSum := root.Val
	if isValCount(l) {
		lSum += l
	}
	r := dfs(root.Right)
	rSum := root.Val
	if isValCount(r) {
		rSum += r
	}

	return max(root.Val, lSum, rSum)
}

func isValCount(val int) bool {
	return val > -1001
}

func Test_MaxPathSum() {
	//root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	//ans1 := maxPathSum(root1)
	//log.Printf("ans1: %v", ans1)
	//
	//root2 := &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	//ans2 := maxPathSum(root2)
	//log.Printf("ans2: %v", ans2)
	//
	//root3 := &TreeNode{Val: -3}
	//ans3 := maxPathSum(root3)
	//log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{Val: 2, Left: &TreeNode{Val: -1}}
	ans4 := maxPathSum(root4)
	log.Printf("ans4: %v", ans4)
}
