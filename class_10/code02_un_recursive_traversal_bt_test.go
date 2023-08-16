package class_10

import (
	"fmt"
	"testing"
)

func TestRecursiveTraversalBT(t *testing.T) {
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3, Left: &Node{Value: 6}, Right: &Node{Value: 7}},
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
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3, Left: &Node{Value: 6}, Right: &Node{Value: 7}},
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
