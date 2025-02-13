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

// first try - not very efficient
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

// neetcode solution
func maxPathSum2(root *TreeNode) int {
	res := math.MinInt
	dfs2(root, &res)
	return res
}

func dfs2(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	l := getMax(root.Left)
	r := getMax(root.Right)
	*res = max(*res, root.Val+l+r)

	dfs2(root.Left, res)
	dfs2(root.Right, res)
}

func getMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getMax(root.Left)
	r := getMax(root.Right)

	return max(0, root.Val+max(l, r))
}

func Test_MaxPathSum() {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	ans1 := maxPathSum(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	ans2 := maxPathSum(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{Val: -3}
	ans3 := maxPathSum(root3)
	log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{Val: 2, Left: &TreeNode{Val: -1}}
	ans4 := maxPathSum(root4)
	log.Printf("ans4: %v", ans4)

	root5 := generateTree([]string{"9", "6", "-3", "null", "null", "-6", "2", "null", "null", "2", "null", "-6", "-6", "-6"})
	ans5 := maxPathSum(root5)
	log.Printf("ans5: %v", ans5)
}

func Test_MaxPathSum2() {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	ans1 := maxPathSum2(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	ans2 := maxPathSum2(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{Val: -3}
	ans3 := maxPathSum2(root3)
	log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{Val: 2, Left: &TreeNode{Val: -1}}
	ans4 := maxPathSum2(root4)
	log.Printf("ans4: %v", ans4)

	root5 := generateTree([]string{"9", "6", "-3", "null", "null", "-6", "2", "null", "null", "2", "null", "-6", "-6", "-6"})
	ans5 := maxPathSum2(root5)
	log.Printf("ans5: %v", ans5)
}
