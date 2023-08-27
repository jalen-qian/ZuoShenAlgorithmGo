package class_11_12

import "fmt"

// 递归方式遍历二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type RecursiveTraversalBT struct{}

// Pre 先序遍历
func (r *RecursiveTraversalBT) Pre(root *TreeNode) {
	if root == nil {
		return
	}
	// 先打印根节点
	fmt.Printf("%d ", root.Val)
	// 再打印左子树
	r.Pre(root.Left)
	// 再打印右子树
	r.Pre(root.Right)
}

// In 中序遍历
func (r *RecursiveTraversalBT) In(root *TreeNode) {
	if root == nil {
		return
	}
	// 先打印左子树
	r.In(root.Left)
	// 再打印自己
	fmt.Printf("%d ", root.Val)
	// 最后打印右子树
	r.In(root.Right)
}

// Pos 后序遍历
func (r *RecursiveTraversalBT) Pos(root *TreeNode) {
	if root == nil {
		return
	}
	// 先打印左子树
	r.Pos(root.Left)
	// 再打印右子树
	r.Pos(root.Right)
	// 最后打印自己
	fmt.Printf("%d ", root.Val)
}
