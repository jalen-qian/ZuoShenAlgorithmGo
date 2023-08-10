package class_09

import "ZuoShenAlgorithmGo/class_03"

// 判断一个链表是否是回文结构

// IsPalindrome1 使用容器实现（栈）
func IsPalindrome1(head *Node) bool {
	// 创建一个栈
	stack := class_03.NewMyStack[int]()
	cur := head
	// 遍历链表，将所有节点的值压入栈
	for cur != nil {
		stack.Push(cur.Value)
		cur = cur.Next
	}
	// 再次遍历链表，并一个个弹出栈，如果弹出的值与当前值不同，则不是回文
	cur = head
	for cur != nil {
		if cur.Value != stack.Pop() {
			return false
		}
	}
	// 每个都弹出一致，则是回文链表，null链表两个for都不走，也算回文链表
	return true
}

// IsPalindrome2 不使用容器实现
func IsPalindrome2(head *Node) bool {

	return true
}
