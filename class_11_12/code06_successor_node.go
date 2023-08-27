package class_11_12

import "ZuoShenAlgorithmGo/class_03"

// 给定一个有父指针的二叉树的某个节点，返回该节点的后继节点

type TreeNodeP struct {
	Val    int
	Parent *TreeNodeP
	Left   *TreeNodeP
	Right  *TreeNodeP
}

func GetSuccessorNode(node *TreeNodeP) *TreeNodeP {
	if node == nil {
		return nil
	}
	// 有右子树，则获取右子树最左侧的节点
	if node.Right != nil {
		return getLeftMost(node.Right)
	} else {
		// 没有右子树
		parent := node.Parent
		for parent != nil && parent.Right == node {
			node = parent
			parent = node.Parent
		}
		return parent
	}
}

// 找一棵树最左侧的节点
func getLeftMost(node *TreeNodeP) *TreeNodeP {
	// 一直往左划，直到划到Left为空
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// GetSuccessorNodeNormal 普通方法找到后继节点，同时做对数器测试
func GetSuccessorNodeNormal(node *TreeNodeP) *TreeNodeP {
	if node == nil {
		return nil
	}
	// 找到根节点
	root := node
	for root.Parent != nil {
		root = root.Parent
	}
	// 中序遍历
	inQueue := class_03.NewMyQueue[*TreeNodeP]()
	inTreeNodeP(root, inQueue)
	// 中序遍历顺序弹出，并找下一个
	for !inQueue.IsEmpty() && inQueue.Peek() != node {
		inQueue.Poll()
	}
	// 弹出当前节点
	inQueue.Poll()
	// 弹出当前节点的下一个
	return inQueue.Poll()
}

func inTreeNodeP(root *TreeNodeP, inQueue *class_03.MyQueue[*TreeNodeP]) {
	if root == nil {
		return
	}
	inTreeNodeP(root.Left, inQueue)
	inQueue.Push(root)
	inTreeNodeP(root.Right, inQueue)
}
