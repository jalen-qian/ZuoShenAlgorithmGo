package class_11_12

import (
	"fmt"
	"testing"
)

func TestRecursiveTraversalBT(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
	}
	r := &RecursiveTraversalBT{}
	fmt.Print("先序遍历：")
	r.Pre(root)
	fmt.Println()

	fmt.Print("中序遍历：")
	r.In(root)
	fmt.Println()

	fmt.Print("后序遍历：")
	r.Pos(root)
	fmt.Println()
}

func TestUnRecursiveTraversalBT(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
	}
	PrintBT(root)
	fmt.Println()

	r := &UnRecursiveTraversalBT{}
	fmt.Print("先序遍历：")
	r.Pre(root)
	fmt.Println()

	fmt.Print("中序遍历：")
	r.In(root)
	fmt.Println()

	fmt.Print("后序遍历：")
	r.Pos1(root)
	fmt.Println()
}
