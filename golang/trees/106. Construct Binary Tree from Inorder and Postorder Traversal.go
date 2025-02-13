package trees

import "slices"

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	mid := slices.Index(inorder, rootVal)

	root.Right = buildTree106(inorder[mid+1:], postorder[mid:len(postorder)-1])
	root.Left = buildTree106(inorder[:mid], postorder[:mid])
	return root
}

func buildTree106_2(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for k, v := range inorder {
		inorderMap[v] = k
	}

	postIdx := len(postorder) - 1
	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := postorder[postIdx]
		postIdx--
		root := &TreeNode{Val: rootVal}
		mid := inorderMap[rootVal]

		root.Right = dfs(mid+1, right)
		root.Left = dfs(left, mid-1)
		return root
	}
	return dfs(0, len(postorder)-1)
}

func Test_BuildTree106() {
	inorder1 := []int{9, 3, 15, 20, 7}
	preorder1 := []int{9, 15, 7, 20, 3}
	ans1 := buildTree106(inorder1, preorder1)
	printTree(ans1)
}

func Test_BuildTree106_2() {
	inorder1 := []int{9, 3, 15, 20, 7}
	preorder1 := []int{9, 15, 7, 20, 3}
	ans1 := buildTree106_2(inorder1, preorder1)
	printTree(ans1)
}
