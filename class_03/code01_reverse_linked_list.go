package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
)

// reverseLinkedList 反转单向链表
func reverseLinkedList(head *utils.Node) *utils.Node {
	var pre *utils.Node
	var next *utils.Node
	for head.Next != nil {
		// 先记住下一个节点的位置
		next = head.Next
		// 当前节点往前置
		head.Next = pre
		// pre来到当前节点
		pre = head
		// 当前节点往下一个节点跳
		head = next
	}
	return head
}

func copyLinkedList(node *utils.Node) *utils.Node {
	if node == nil {
		return node
	}
	newNode := &utils.Node{Value: node.Value}
	old := node.Next
	cur := newNode
	for old != nil {
		cur.Next = old
		cur = cur.Next
		old = old.Next
	}
	return newNode
}

// checkLinkedListEqual 校验两个链表是否全等
// 全等条件：长度相同，且每个位置的值相同
// 如果一个为nil，则另一个也必须为nil
func checkLinkedListEqual(head1 *utils.Node, head2 *utils.Node) bool {
	if head1 == nil && head2 != nil {
		return false
	}
	if head1 != nil && head2 == nil {
		return false
	}
	if head1 == nil && head2 == nil {
		return true
	}
	var cur1 = head1
	var cur2 = head2
	for cur1 != nil {
		if cur2 == nil {
			return false
		}
		if cur1.Value != cur2.Value {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return true
}

func main() {
	testTimes := 1
	for i := 0; i < testTimes; i++ {
		// 测试方式：反转两次，测试反转后的结果是否与原始链表全等
		linkedList := utils.GenerateRandomLinkedList(1000, -100, 100)
		copyList := copyLinkedList(linkedList)
		reversedList := reverseLinkedList(linkedList)
		restore := reverseLinkedList(reversedList)
		if !checkLinkedListEqual(copyList, restore) {
			fmt.Println("Fucking fucked!!!")
			return
		}
	}
	fmt.Println("Great!!!")
}
