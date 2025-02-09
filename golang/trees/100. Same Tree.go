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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(p *TreeNode, q *TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		left := dfs(p.Left, q.Left)
		right := dfs(p.Right, q.Right)
		return left && right && p.Val == q.Val
	}

	return dfs(p, q)
}

func Test_IsSameTree() {
	p11 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	p12 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	ans1 := isSameTree(p11, p12)
	log.Printf("ans1: %v", ans1)

	p21 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	p22 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	ans2 := isSameTree(p21, p22)
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
	ans3 := isSameTree(p31, p32)
	log.Printf("ans3: %v", ans3)

	p41 := &TreeNode{
		Val: 1,
	}
	ans4 := isSameTree(p41, nil)
	log.Printf("ans4: %v", ans4)

	ans5 := isSameTree(nil, nil)
	log.Printf("ans5: %v", ans5)
}
