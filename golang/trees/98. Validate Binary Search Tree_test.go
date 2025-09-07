package trees

import (
	"log"
	"testing"
)

func Test_IsValidBST(t *testing.T) {
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	ans1 := isValidBST(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	ans2 := isValidBST(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 8}},
	}
	ans3 := isValidBST(root3)
	log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}
	ans4 := isValidBST(root4)
	log.Printf("ans4: %v", ans4)
}

func Test_IsValidBST2(t *testing.T) {
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	ans1 := isValidBST2(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	ans2 := isValidBST2(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 8}},
	}
	ans3 := isValidBST2(root3)
	log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}
	ans4 := isValidBST2(root4)
	log.Printf("ans4: %v", ans4)
}

func Test_IsValidBST3(t *testing.T) {
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	ans1 := isValidBST3(root1)
	log.Printf("ans1: %v", ans1)

	root2 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	ans2 := isValidBST3(root2)
	log.Printf("ans2: %v", ans2)

	root3 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 8}},
	}
	ans3 := isValidBST3(root3)
	log.Printf("ans3: %v", ans3)

	root4 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}},
	}
	ans4 := isValidBST3(root4)
	log.Printf("ans4: %v", ans4)
}
