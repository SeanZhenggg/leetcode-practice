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
// recursive solution
func kthSmallest(root *TreeNode, k int) int {
	kth := 0
	ret := 0
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		found := dfs(root.Left)
		if found {
			return found
		}
		kth += 1
		if kth == k {
			ret = root.Val
		}
		return dfs(root.Right)
	}

	dfs(root)
	return ret
}

// recursive + slice solution
func kthSmallest2(root *TreeNode, k int) int {
	store := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		store = append(store, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return store[k-1]
}

// iterative solution
func kthSmallest3(root *TreeNode, k int) int {
	st := make([]*TreeNode, 0)
	node := root
	for len(st) > 0 || node != nil {
		if node != nil {
			st = append(st, node)
			node = node.Left
		} else {
			node = st[len(st)-1]
			st = st[:len(st)-1]
			k -= 1
			if k == 0 {
				return node.Val
			}
			node = node.Right
		}
	}
	return 0 // unreachable
}

func Test_KthSmallest() {
	root1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}
	k1 := 1
	ans1 := kthSmallest(root1, k1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	k2 := 3
	ans2 := kthSmallest(root2, k2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{Val: 12,
		Left:  &TreeNode{Val: 10, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 9}}}, Right: &TreeNode{Val: 11}},
		Right: &TreeNode{Val: 14, Left: &TreeNode{Val: 13}},
	}
	k3 := 5
	ans3 := kthSmallest(root3, k3)
	log.Printf("ans3: %v", ans3)
}

func Test_KthSmallest2() {
	root1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}
	k1 := 1
	ans1 := kthSmallest2(root1, k1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	k2 := 3
	ans2 := kthSmallest2(root2, k2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{Val: 12,
		Left:  &TreeNode{Val: 10, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 9}}}, Right: &TreeNode{Val: 11}},
		Right: &TreeNode{Val: 14, Left: &TreeNode{Val: 13}},
	}
	k3 := 5
	ans3 := kthSmallest2(root3, k3)
	log.Printf("ans3: %v", ans3)
}

func Test_KthSmallest3() {
	root1 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 4},
	}
	k1 := 1
	ans1 := kthSmallest3(root1, k1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	k2 := 3
	ans2 := kthSmallest3(root2, k2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{Val: 12,
		Left:  &TreeNode{Val: 10, Left: &TreeNode{Val: 8, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 9}}, Right: &TreeNode{Val: 11}},
		Right: &TreeNode{Val: 14, Left: &TreeNode{Val: 13}},
	}
	k3 := 5
	ans3 := kthSmallest3(root3, k3)
	log.Printf("ans3: %v", ans3)
}
