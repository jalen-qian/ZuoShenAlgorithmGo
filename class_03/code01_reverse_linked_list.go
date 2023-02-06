package main

import (
	"ZuoShenAlgorithmGo/utils"
)

// reverseLinkedList 反转单向链表
func reverseLinkedList(head *utils.Node) *utils.Node {
	var pre *utils.Node
	var next *utils.Node
	for head != nil {
		// 先记住下一个节点的位置
		next = head.Next
		// 当前节点往前指
		head.Next = pre
		// pre来到当前节点
		pre = head
		// 当前节点往下一个节点跳
		head = next
	}
	return pre
}

// testReverseLinkedList 构造一个新的单向链表，是原始链表反转的形式
func testReverseLinkedList(head *utils.Node) *utils.Node {
	if head == nil {
		return nil
	}
	// 用一个数组来存
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Value)
		cur = cur.Next
	}
	// 反向构造
	newHead := &utils.Node{Value: arr[len(arr)-1]}
	cur = newHead
	for i := len(arr) - 2; i >= 0; i-- {
		cur.Next = &utils.Node{Value: arr[i]}
		cur = cur.Next
	}
	return newHead
}
