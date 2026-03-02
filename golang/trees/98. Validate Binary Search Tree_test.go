package trees

import "testing"

type validateBSTCase struct {
	root     *TreeNode
	expected bool
}

var validateBSTCases = []validateBSTCase{
	{
		root:     &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		expected: true,
	},
	{
		root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}},
		expected: false,
	},
	{
		root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 8}}},
		expected: false,
	},
	{
		root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}},
		expected: false,
	},
}

func Test_IsValidBST(t *testing.T) {
	for _, c := range validateBSTCases {
		ans := isValidBST(c.root)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func Test_IsValidBST2(t *testing.T) {
	for _, c := range validateBSTCases {
		ans := isValidBST2(c.root)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}

func Test_IsValidBST3(t *testing.T) {
	for _, c := range validateBSTCases {
		ans := isValidBST3(c.root)
		if ans != c.expected {
			t.Errorf("answer is %v, want %v", ans, c.expected)
		}
	}
}
