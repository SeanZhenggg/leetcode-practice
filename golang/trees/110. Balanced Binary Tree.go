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
func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) (int, bool)
	dfs = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}

		lHeight, isLeftBalanced := dfs(root.Left)
		rHeight, isRightBalanced := dfs(root.Right)

		if !isLeftBalanced || !isRightBalanced {
			return 0, false
		}

		if math.Abs(float64(lHeight)-float64(rHeight)) > 1 {
			return 0, false
		}

		return 1 + max(lHeight, rHeight), true
	}

	_, isBalance := dfs(root)
	return isBalance
}

func isBalanced2(root *TreeNode) bool {
	// An empty tree satisfies the definition of a balanced tree
	if root == nil {
		return true
	}
	// Check if subtrees have height within 1. If they do, check if the subtrees are balanced
	return math.Abs(float64(height(root.Left))-float64(height(root.Right))) < 2 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}

func isBalancedReview1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(getHeight(root.Left)-getHeight(root.Right))) <= 1 && isBalancedReview1(root.Left) && isBalancedReview1(root.Right)
}

func isBalancedReview2(root *TreeNode) bool {

	var dfs func(root *TreeNode) (int, bool)
	dfs = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}

		hOfLeft, isLeftBalanced := dfs(root.Left)
		hOfRight, isRightBalanced := dfs(root.Right)

		return 1 + max(hOfLeft, hOfRight), isLeftBalanced && isRightBalanced && math.Abs(float64(hOfLeft-hOfRight)) <= 1
	}

	_, isBalance := dfs(root)
	return isBalance
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func height(root *TreeNode) int {
	// An empty tree has height -1
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func Test_IsBalanced() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := isBalanced(root1)
	log.Printf("ans1: %v", ans1)
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	ans2 := isBalanced(root2)
	log.Printf("ans2: %v", ans2)

	ans3 := isBalanced(nil)
	log.Printf("ans3: %v", ans3)
}

func Test_IsBalanced2() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := isBalanced2(root1)
	log.Printf("ans1: %v", ans1)
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	ans2 := isBalanced2(root2)
	log.Printf("ans2: %v", ans2)

	ans3 := isBalanced2(nil)
	log.Printf("ans3: %v", ans3)
}

func Test_IsBalancedReview1() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := isBalancedReview1(root1)
	log.Printf("ans1: %v", ans1)
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	ans2 := isBalancedReview1(root2)
	log.Printf("ans2: %v", ans2)

	ans3 := isBalancedReview1(nil)
	log.Printf("ans3: %v", ans3)
}

func Test_IsBalancedReview2() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	ans1 := isBalancedReview2(root1)
	log.Printf("ans1: %v", ans1)
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	ans2 := isBalancedReview2(root2)
	log.Printf("ans2: %v", ans2)

	ans3 := isBalancedReview2(nil)
	log.Printf("ans3: %v", ans3)
}
