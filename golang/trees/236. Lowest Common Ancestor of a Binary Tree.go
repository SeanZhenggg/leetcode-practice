package trees

import "log"

func lowestCommonAncestorBT(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, p, q *TreeNode) (*TreeNode, bool)
	dfs = func(root, p, q *TreeNode) (*TreeNode, bool) {
		if root == nil {
			return nil, false
		}

		found, lfound := dfs(root.Left, p, q)
		if found != nil {
			return found, true
		}
		found, rfound := dfs(root.Right, p, q)
		if found != nil {
			return found, true
		}

		var midFound bool
		if root.Val == p.Val || root.Val == q.Val {
			midFound = true
		}
		if (midFound && lfound) || (midFound && rfound) || (lfound && rfound) {
			return root, true
		} else {
			return nil, midFound || lfound || rfound
		}
	}
	d, _ := dfs(root, p, q)
	return d
}

func Test_LowestCommonAncestorBT() {
	root1 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}},
	}
	p1 := &TreeNode{Val: 5}
	q1 := &TreeNode{Val: 1}
	ans1 := lowestCommonAncestorBT(root1, p1, q1)
	log.Printf("ans1: %v", ans1.Val)

	root2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}},
	}
	p2 := &TreeNode{Val: 5}
	q2 := &TreeNode{Val: 4}
	ans2 := lowestCommonAncestorBT(root2, p2, q2)
	log.Printf("ans2: %v", ans2.Val)

	root3 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	p3 := &TreeNode{Val: 1}
	q3 := &TreeNode{Val: 2}
	ans3 := lowestCommonAncestorBT(root3, p3, q3)
	log.Printf("ans3: %v", ans3.Val)
}
