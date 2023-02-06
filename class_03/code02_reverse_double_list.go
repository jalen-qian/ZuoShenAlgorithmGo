package main

import (
	"ZuoShenAlgorithmGo/utils"
)

func reverseDoubleList(head *utils.DoubleNode) *utils.DoubleNode {
	if head == nil {
		return nil
	}
	var pre *utils.DoubleNode
	var next *utils.DoubleNode
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