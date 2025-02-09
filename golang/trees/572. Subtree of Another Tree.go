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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(root *TreeNode, subRoot *TreeNode) bool
	dfs = func(root *TreeNode, subRoot *TreeNode) bool {
		var l, r bool
		if root.Left != nil {
			l = dfs(root.Left, subRoot)
		}
		if root.Right != nil {
			r = dfs(root.Right, subRoot)
		}

		return isSameTree(root, subRoot) || l || r
	}

	return dfs(root, subRoot)
}

func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(root *TreeNode, subRoot *TreeNode) bool
	dfs = func(root *TreeNode, subRoot *TreeNode) bool {
		if root == nil {
			return false
		}

		if isSameTree(root, subRoot) {
			return true
		}

		return dfs(root.Left, subRoot) || dfs(root.Right, subRoot)
	}

	return dfs(root, subRoot)
}

func Test_IsSubtree() {
	p11 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 5},
	}
	p12 := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	ans1 := isSubtree(p11, p12)
	log.Printf("ans1: %v", ans1)

	p21 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}},
		Right: &TreeNode{Val: 5},
	}
	p22 := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	ans2 := isSubtree(p21, p22)
	log.Printf("ans2: %v", ans2)

	p31 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	p32 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	ans3 := isSubtree(p31, p32)
	log.Printf("ans3: %v", ans3)
}

func Test_IsSubtree2() {
	p11 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 5},
	}
	p12 := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	ans1 := isSubtree2(p11, p12)
	log.Printf("ans1: %v", ans1)

	p21 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}},
		Right: &TreeNode{Val: 5},
	}
	p22 := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	ans2 := isSubtree2(p21, p22)
	log.Printf("ans2: %v", ans2)

	p31 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	p32 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	ans3 := isSubtree2(p31, p32)
	log.Printf("ans3: %v", ans3)
}
