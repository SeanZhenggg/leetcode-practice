package trees

import "log"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	ancestorsOfP := appendAncestors(root, p)
	ancestorsOfQ := appendAncestors(root, q)

	j := 0
	if len(ancestorsOfP) > len(ancestorsOfQ) {
		for i := 0; i < len(ancestorsOfP); i++ {
			if ancestorsOfP[i] == ancestorsOfQ[j] {
				return ancestorsOfP[i]
			}

			if i >= (len(ancestorsOfP) - len(ancestorsOfQ)) {
				j++
			}
		}
	} else {
		for i := 0; i < len(ancestorsOfQ); i++ {
			if ancestorsOfQ[i] == ancestorsOfP[j] {
				return ancestorsOfQ[i]
			}
			if i >= (len(ancestorsOfQ) - len(ancestorsOfP)) {
				j++
			}
		}
	}
	return nil
}

func appendAncestors(root *TreeNode, p *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val {
		return []*TreeNode{root}
	}
	nodeOfP := appendAncestors(root.Left, p)
	if len(nodeOfP) > 0 {
		return append(nodeOfP, root)
	}
	nodeOfP = appendAncestors(root.Right, p)
	if nodeOfP != nil {
		return append(nodeOfP, root)
	}
	return nil
}

func lowestCommonAncestorBST2(root, p, q *TreeNode) *TreeNode {
	parentVal := root.Val
	pVal := p.Val
	qVal := q.Val

	if pVal < parentVal && qVal < parentVal {
		return lowestCommonAncestorBST2(root.Left, p, q)
	} else if pVal > parentVal && qVal > parentVal {
		return lowestCommonAncestorBST2(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestorBST3(root, p, q *TreeNode) *TreeNode {
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

func Test_LowestCommonAncestorBST() {
	root1 := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	p1 := &TreeNode{Val: 2}
	q1 := &TreeNode{Val: 8}
	ans1 := lowestCommonAncestorBST(root1, p1, q1)
	log.Printf("ans1: %v", ans1.Val)

	root2 := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	p2 := &TreeNode{Val: 2}
	q2 := &TreeNode{Val: 4}
	ans2 := lowestCommonAncestorBST(root2, p2, q2)
	log.Printf("ans2: %v", ans2.Val)

	root3 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
	}
	p3 := &TreeNode{Val: 2}
	q3 := &TreeNode{Val: 1}
	ans3 := lowestCommonAncestorBST(root3, p3, q3)
	log.Printf("ans3: %v", ans3.Val)

	root4 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}
	p4 := &TreeNode{Val: 2}
	q4 := &TreeNode{Val: 3}
	ans4 := lowestCommonAncestorBST(root4, p4, q4)
	log.Printf("ans4: %v", ans4.Val)

	root5 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	p5 := &TreeNode{Val: 1}
	q5 := &TreeNode{Val: 4}
	ans5 := lowestCommonAncestorBST(root5, p5, q5)
	log.Printf("ans5: %v", ans5.Val)
}

func Test_LowestCommonAncestorBST2() {
	root1 := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	p1 := &TreeNode{Val: 2}
	q1 := &TreeNode{Val: 8}
	ans1 := lowestCommonAncestorBST2(root1, p1, q1)
	log.Printf("ans1: %v", ans1.Val)

	root2 := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	p2 := &TreeNode{Val: 1}
	q2 := &TreeNode{Val: 10}
	ans2 := lowestCommonAncestorBST2(root2, p2, q2)
	log.Printf("ans2: %v", ans2.Val)

	root3 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
	}
	p3 := &TreeNode{Val: 2}
	q3 := &TreeNode{Val: 1}
	ans3 := lowestCommonAncestorBST2(root3, p3, q3)
	log.Printf("ans3: %v", ans3.Val)

	root4 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}
	p4 := &TreeNode{Val: 2}
	q4 := &TreeNode{Val: 3}
	ans4 := lowestCommonAncestorBST2(root4, p4, q4)
	log.Printf("ans4: %v", ans4.Val)

	root5 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	p5 := &TreeNode{Val: 1}
	q5 := &TreeNode{Val: 4}
	ans5 := lowestCommonAncestorBST2(root5, p5, q5)
	log.Printf("ans5: %v", ans5.Val)
}

func Test_LowestCommonAncestorBST3() {
	root1 := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	p1 := &TreeNode{Val: 2}
	q1 := &TreeNode{Val: 8}
	ans1 := lowestCommonAncestorBST3(root1, p1, q1)
	log.Printf("ans1: %v", ans1.Val)

	root2 := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	p2 := &TreeNode{Val: 2}
	q2 := &TreeNode{Val: 4}
	ans2 := lowestCommonAncestorBST3(root2, p2, q2)
	log.Printf("ans2: %v", ans2.Val)

	root3 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
	}
	p3 := &TreeNode{Val: 2}
	q3 := &TreeNode{Val: 1}
	ans3 := lowestCommonAncestorBST3(root3, p3, q3)
	log.Printf("ans3: %v", ans3.Val)

	root4 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}
	p4 := &TreeNode{Val: 2}
	q4 := &TreeNode{Val: 3}
	ans4 := lowestCommonAncestorBST3(root4, p4, q4)
	log.Printf("ans4: %v", ans4.Val)

	root5 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	p5 := &TreeNode{Val: 1}
	q5 := &TreeNode{Val: 4}
	ans5 := lowestCommonAncestorBST3(root5, p5, q5)
	log.Printf("ans5: %v", ans5.Val)
}
