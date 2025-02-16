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
//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	var dfs func(p *TreeNode, q *TreeNode) bool
//	dfs = func(p *TreeNode, q *TreeNode) bool {
//		if p == nil && q == nil {
//			return true
//		}
//		if p == nil || q == nil || p.Val != q.Val {
//			return false
//		}
//
//		left := dfs(p.Left, q.Left)
//		right := dfs(p.Right, q.Right)
//		return left && right && p.Val == q.Val
//	}
//
//	return dfs(p, q)
//}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type Pairs struct {
	P *TreeNode
	Q *TreeNode
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue := make([]Pairs, 0)
	queue = append(queue, Pairs{p, q})
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if !check(top.P, top.Q) {
			return false
		}

		if top.P != nil && top.Q != nil {
			queue = append(queue, Pairs{top.P.Left, top.Q.Left})
			queue = append(queue, Pairs{top.P.Right, top.Q.Right})
		}
	}
	return true
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	return p.Val == q.Val
}

func isSameTreeReview1(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	return p.Val == q.Val && isSameTreeReview1(p.Left, q.Left) && isSameTreeReview1(p.Right, q.Right)
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

func Test_IsSameTreeReview1() {
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
	ans1 := isSameTreeReview1(p11, p12)
	log.Printf("ans1: %v", ans1)

	p21 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	p22 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	ans2 := isSameTreeReview1(p21, p22)
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
	ans3 := isSameTreeReview1(p31, p32)
	log.Printf("ans3: %v", ans3)

	p41 := &TreeNode{
		Val: 1,
	}
	ans4 := isSameTreeReview1(p41, nil)
	log.Printf("ans4: %v", ans4)

	ans5 := isSameTreeReview1(nil, nil)
	log.Printf("ans5: %v", ans5)
}

func Test_IsSameTree2() {
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
	ans1 := isSameTree2(p11, p12)
	log.Printf("ans1: %v", ans1)

	p21 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	p22 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}
	ans2 := isSameTree2(p21, p22)
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
	ans3 := isSameTree2(p31, p32)
	log.Printf("ans3: %v", ans3)

	p41 := &TreeNode{
		Val: 1,
	}
	ans4 := isSameTree2(p41, nil)
	log.Printf("ans4: %v", ans4)

	ans5 := isSameTree2(nil, nil)
	log.Printf("ans5: %v", ans5)
}
