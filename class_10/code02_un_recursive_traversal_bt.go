package class_10

import (
	"ZuoShenAlgorithmGo/class_03"
	"fmt"
)

// 使用非递归的方式

type UnRecursiveTraversalBT struct{}

// Pre 先序遍历
func (r *UnRecursiveTraversalBT) Pre(root *Node) {
	if root == nil {
		return
	}
	// 创建一个栈，并先把根节点压入
	stack := class_03.NewMyStack[*Node]()
	stack.Push(root)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		// 出栈就打印
		fmt.Printf("%d ", cur.Value)
		// 有右子树，就入栈
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		// 有左子树，就入栈
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
	fmt.Println()
}

// In 中序遍历
func (r *UnRecursiveTraversalBT) In(root *Node) {
	if root == nil {
		return
	}
	// 创建一个栈，如果当前节点有左右孩子，就压入，没有就弹出
	stack := class_03.NewMyStack[*Node]()
	// 3 2 1
	cur := root
	for !stack.IsEmpty() || cur != nil {
		if cur.Left != nil || cur.Right != nil {
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		} else {

		}
	}
	fmt.Println()
}

// Pos 后序遍历
func (r *UnRecursiveTraversalBT) Pos(root *Node) {

}
