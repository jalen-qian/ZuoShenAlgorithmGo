package class_10

import "fmt"

// 递归方式遍历二叉树

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type RecursiveTraversalBT struct{}

// Pre 先序遍历
func (r *RecursiveTraversalBT) Pre(root *Node) {
	if root == nil {
		return
	}
	// 先打印根节点
	fmt.Printf("%d ", root.Value)
	// 再打印左子树
	r.Pre(root.Left)
	// 再打印右子树
	r.Pre(root.Right)
}

// In 中序遍历
func (r *RecursiveTraversalBT) In(root *Node) {
	if root == nil {
		return
	}
	// 先打印左子树
	r.In(root.Left)
	// 再打印自己
	fmt.Printf("%d ", root.Value)
	// 最后打印右子树
	r.In(root.Right)
}

// Pos 后序遍历
func (r *RecursiveTraversalBT) Pos(root *Node) {
	if root == nil {
		return
	}
	// 先打印左子树
	r.Pos(root.Left)
	// 再打印右子树
	r.Pos(root.Right)
	// 最后打印自己
	fmt.Printf("%d ", root.Value)
}
