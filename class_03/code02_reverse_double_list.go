package class_03

func reverseDoubleList(head *ListDoubleNode) *ListDoubleNode {
	if head == nil {
		return nil
	}
	var pre *ListDoubleNode
	var next *ListDoubleNode
	for head != nil {
		// 先记住下一个节点
		next = head.Next
		// 当前节点的next往前指
		head.Next = pre
		// 当前节点的Pre指针往后指
		head.Last = next
		// pre跳到当前位置
		pre = head
		// 当前位置跳到下一个
		head = next
	}
	return pre
}
