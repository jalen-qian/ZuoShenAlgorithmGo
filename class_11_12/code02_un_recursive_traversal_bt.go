package class_11_12

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
	// 创建一个栈，如果当前节点有左孩子，就压入，并不断往左边界靠
	stack := class_03.NewMyStack[*Node]()
	cur := root
	for !stack.IsEmpty() || cur != nil {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			cur = stack.Pop()
			fmt.Printf("%d ", cur.Value)
			cur = cur.Right
		}
	}
	fmt.Println()
}

// Pos1 后序遍历
func (r *UnRecursiveTraversalBT) Pos1(root *Node) {
	if root == nil {
		return
	}
	s1 := class_03.NewMyStack[*Node]()
	s2 := class_03.NewMyStack[*Node]()
	s1.Push(root)
	for !s1.IsEmpty() {
		head := s1.Pop()
		// 只要出栈，不打印，而是压入另一个栈
		s2.Push(head)
		// 弹出后依次压入左和右，实现整体 头 右 左 的顺序弹出
		if head.Left != nil {
			s1.Push(head.Left)
		}
		if head.Right != nil {
			s1.Push(head.Right)
		}
	}
	// 所有事情做完，依次将s2弹出
	for !s2.IsEmpty() {
		fmt.Printf("%d ", s2.Pop().Value)
	}
}
