package trees

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	mid := slices.Index(inorder, preorder[0])
	root.Left = buildTree105(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree105(preorder[mid+1:], inorder[mid+1:])

	return root
}

func buildTree105_2(preorder []int, inorder []int) *TreeNode {
	var preIdx = 0
	var inorderMap = make(map[int]int)
	for k, v := range inorder {
		inorderMap[v] = k
	}

	var dfs func(left int, right int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[preIdx]
		preIdx++
		root := &TreeNode{
			Val: rootVal,
		}
		mid := inorderMap[rootVal]

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)
		return root
	}

	return dfs(0, len(preorder)-1)
}

func Test_BuildTree105() {
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	ans1 := buildTree105(preorder1, inorder1)
	printTree(ans1)
}

func Test_BuildTree105_2() {
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	ans1 := buildTree105_2(preorder1, inorder1)
	printTree(ans1)
}
