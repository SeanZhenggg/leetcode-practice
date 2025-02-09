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

	//p31 := &TreeNode{
	//	Val:   1,
	//	Left:  &TreeNode{Val: 2},
	//	Right: &TreeNode{Val: 1},
	//}
	//p32 := &TreeNode{
	//	Val:   1,
	//	Left:  &TreeNode{Val: 1},
	//	Right: &TreeNode{Val: 2},
	//}
	//ans3 := isSameTree(p31, p32)
	//log.Printf("ans3: %v", ans3)
	//
	//p41 := &TreeNode{
	//	Val: 1,
	//}
	//ans4 := isSameTree(p41, nil)
	//log.Printf("ans4: %v", ans4)
	//
	//ans5 := isSameTree(nil, nil)
	//log.Printf("ans5: %v", ans5)
}
