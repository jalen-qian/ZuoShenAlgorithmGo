package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
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

// test 构造一个新的单向链表，是原始链表反转的形式
func test(head *utils.Node) *utils.Node {
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

func copyLinkedList(node *utils.Node) *utils.Node {
	if node == nil {
		return nil
	}
	newNode := &utils.Node{Value: node.Value}
	oldCur := node
	newCur := newNode
	for oldCur.Next != nil {
		newCur.Next = &utils.Node{Value: oldCur.Next.Value}
		oldCur = oldCur.Next
		newCur = newCur.Next
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
	cur1 := head1
	cur2 := head2
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
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		// 测试方式：反转两次，测试反转后的结果是否与原始链表全等
		linkedList := utils.GenerateRandomLinkedList(1000, -100, 100)
		copyList := copyLinkedList(linkedList)
		reversedList := reverseLinkedList(linkedList)
		reversedList2 := test(copyList)
		if !checkLinkedListEqual(reversedList, reversedList2) {
			fmt.Println("Fucking fucked!!!")
			fmt.Printf("原始链表：%s\n", utils.SPrintLinkedList(copyList))
			fmt.Printf("反转链表：%s\n", utils.SPrintLinkedList(reversedList))
			fmt.Printf("对数器链表：%s\n", utils.SPrintLinkedList(reversedList2))
			return
		}
	}
	fmt.Println("Great!!!")
}
