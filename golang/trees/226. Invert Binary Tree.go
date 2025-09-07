package trees

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS, post-order traversal, tc: O(n), sc: O(n) (call stack)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = left

	return root
}

// BFS, level order traversal, tc: O(n), sc: O(n)
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		left := top.Left
		top.Left = top.Right
		top.Right = left

		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}

	return root
}

// DFS, post-order traversal, tc: O(n), sc: O(n) (call stack)
func invertTreeReview1(root *TreeNode) *TreeNode {
	// root == nil, return
	// invert left and right
	// dfs recursively to invert
	// post-order traversal

	if root == nil {
		return nil
	}

	leftTree := invertTreeReview1(root.Left)
	root.Left = invertTreeReview1(root.Right)
	root.Right = leftTree
	return root
}

// BFS, level order traversal, tc: O(n), sc: O(n)
func invertTreeReview2(root *TreeNode) *TreeNode {
	// root == nil, return
	// invert left and right
	// use bfs to invert
	// level order traversal

	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		left := top.Left
		top.Left = top.Right
		top.Right = left

		if top.Left != nil {
			queue = append(queue, top.Left)
		}

		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return root
}

func invertTreeReview3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tempLeft := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tempLeft)

	return root
}

// BFS, level order traversal, tc: O(n), sc: O(n)
func invertTreeReview4(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	st := make([]*TreeNode, 0)
	st = append(st, root)

	for len(st) > 0 {
		last := st[0]
		st = st[1:]

		last.Left, last.Right = last.Right, last.Left
		if last.Left != nil {
			st = append(st, last.Left)
		}

		if last.Right != nil {
			st = append(st, last.Right)
		}
	}

	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	st := make([]*TreeNode, 0)
	st = append(st, root)
	for len(st) > 0 {
		node := st[0]
		st = st[1:]

		fmt.Printf("%d\t", node.Val)
		if node.Left != nil {
			st = append(st, node.Left)
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
	}
}

func generateTree(strArr []string) *TreeNode {
	if len(strArr) == 0 {
		return nil
	}

	// build array of pointer of integer first (becuz could be composed of empty nodes, represented by nil ptr)
	arr := make([]*int, 0, len(strArr))
	for _, v := range strArr {
		if v == "null" {
			arr = append(arr, nil)
		} else {
			val, _ := strconv.ParseInt(v, 10, 64)
			val2 := int(val)
			arr = append(arr, &val2)
		}
	}

	top := arr[0]
	if top == nil {
		return nil
	}
	arr = arr[1:]
	root := &TreeNode{Val: *top}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 && len(arr) > 0 {
		node := queue[0]
		queue = queue[1:]
		top = arr[0]
		arr = arr[1:]
		if top != nil {
			node.Left = &TreeNode{Val: *top}
			queue = append(queue, node.Left)
		}
		if len(arr) > 0 {
			top = arr[0]
			arr = arr[1:]
			if top != nil {
				node.Right = &TreeNode{Val: *top}
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
