package trees

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var ret = make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	lvl := 0

	treeHeight := this.getTreeHeight(root)
	for len(queue) > 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			top := queue[0]
			queue = queue[1:]
			if top == nil {
				if lvl < treeHeight {
					ret = append(ret, "null")
				}
				continue
			}
			ret = append(ret, strconv.Itoa(top.Val))
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
		}
		lvl++
	}
	return strings.Join(ret, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	split := strings.Split(data, ",")

	// build array of pointer of integer first (becuz could be composed of empty nodes, represented by nil ptr)
	arr := make([]*int, 0, len(split))
	for _, v := range split {
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

func (this *Codec) getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(this.getTreeHeight(root.Left), this.getTreeHeight(root.Right))
}

func Test_Constructor() {
	constructor := Constructor()
	root1 := generateTree([]string{"1", "2", "3", "null", "null", "4", "5"})
	data := constructor.serialize(root1)
	ans1 := constructor.deserialize(data)
	printTree(ans1)

	fmt.Println()

	root2 := generateTree([]string{})
	data2 := constructor.serialize(root2)
	ans2 := constructor.deserialize(data2)
	printTree(ans2)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
