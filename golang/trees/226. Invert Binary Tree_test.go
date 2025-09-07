package trees

import (
	"fmt"
	"testing"
)

func Test_InvertTree(t *testing.T) {
	root1 := generateTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	tree := invertTree(root1)

	printTree(tree)

	fmt.Println()

	root2 := generateTree([]string{"2", "1", "3"})
	tree2 := invertTree(root2)

	printTree(tree2)

	root3 := generateTree([]string{})
	tree3 := invertTree(root3)

	printTree(tree3)
}

func Test_InvertTree2(t *testing.T) {
	root1 := generateTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	tree := invertTree2(root1)

	printTree(tree)

	fmt.Println()

	root2 := generateTree([]string{"2", "1", "3"})
	tree2 := invertTree2(root2)

	printTree(tree2)

	root3 := generateTree([]string{})
	tree3 := invertTree2(root3)

	printTree(tree3)
}

func Test_InvertTreeReview1(t *testing.T) {
	root1 := generateTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	tree := invertTreeReview1(root1)

	printTree(tree)

	fmt.Println()

	root2 := generateTree([]string{"2", "1", "3"})
	tree2 := invertTreeReview1(root2)

	printTree(tree2)

	root3 := generateTree([]string{})
	tree3 := invertTreeReview1(root3)

	printTree(tree3)
}

func Test_InvertTreeReview2(t *testing.T) {
	root1 := generateTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	tree := invertTreeReview2(root1)

	printTree(tree)

	fmt.Println()

	root2 := generateTree([]string{"2", "1", "3"})
	tree2 := invertTreeReview2(root2)

	printTree(tree2)

	root3 := generateTree([]string{})
	tree3 := invertTreeReview2(root3)

	printTree(tree3)
}

func Test_InvertTreeReview3(t *testing.T) {
	root1 := generateTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	tree := invertTreeReview3(root1)

	printTree(tree)

	fmt.Println()

	root2 := generateTree([]string{"2", "1", "3"})
	tree2 := invertTreeReview3(root2)

	printTree(tree2)

	root3 := generateTree([]string{})
	tree3 := invertTreeReview3(root3)

	printTree(tree3)
}

func Test_InvertTreeReview4(t *testing.T) {
	root1 := generateTree([]string{"4", "2", "7", "1", "3", "6", "9"})
	tree := invertTreeReview4(root1)

	printTree(tree)

	fmt.Println()

	root2 := generateTree([]string{"2", "1", "3"})
	tree2 := invertTreeReview4(root2)

	printTree(tree2)

	root3 := generateTree([]string{})
	tree3 := invertTreeReview4(root3)

	printTree(tree3)
}
