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
func rightSideView(root *TreeNode) []int {
	// traverse every level
	// at each level, make a temp node representing the rightest node we want
	// at each level, we replace the node with the most right one if we met
	// after each level, push the temp node val into ret

	ret := make([]int, 0)

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ret) == level {
			ret = append(ret, 0)
		}
		ret[level] = root.Val

		dfs(root.Left, level+1)

		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return ret
}

func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	lvl := 0
	for len(queue) > 0 {
		lvlLength := len(queue)
		for i := 0; i < lvlLength; i++ {
			top := queue[0]
			queue = queue[1:]
			if len(ret) == lvl {
				ret = append(ret, top.Val)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
		}
		lvl++
	}

	return ret
}

func Test_RightSideView() {
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	ans1 := rightSideView(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	ans2 := rightSideView(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 3},
	}
	ans3 := rightSideView(root3)
	log.Printf("ans3: %v", ans3)

	ans4 := rightSideView(nil)
	log.Printf("ans4: %v", ans4)
}

func Test_RightSideView2() {
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	ans1 := rightSideView2(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	ans2 := rightSideView2(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 3},
	}
	ans3 := rightSideView2(root3)
	log.Printf("ans3: %v", ans3)

	ans4 := rightSideView2(nil)
	log.Printf("ans4: %v", ans4)
}
