package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Node 单向链表
type Node struct {
	Value int
	Next  *Node
}

func GenNodeValue(minValue int, maxValue int) int {
	return rand.Intn(maxValue-minValue+1) + minValue
}

// GenerateRandomLinkedList 初始化一个随机的链表
func GenerateRandomLinkedList(maxLength int, minValue int, maxValue int) *Node {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength + 1)
	// 如果长度是0，则返回空链表
	if length == 0 {
		return nil
	}
	// 先构造一个头节点
	head := &Node{Value: GenNodeValue(minValue, maxValue)}
	cur := head
	// 还剩length-1个节点
	for i := 1; i < length; i++ {
		value := GenNodeValue(minValue, maxValue)
		cur.Next = &Node{Value: value}
		cur = cur.Next
	}
	return head
}

// SPrintLinkedList 打印链表
func SPrintLinkedList(head *Node) string {
	if head == nil {
		return "null"
	}
	cur := head
	// 遍历链表
	ans := "{"
	for cur != nil {
		ans += fmt.Sprintf("%d -> ", cur.Value)
		cur = cur.Next
	}
	ans += "null}"
	return ans
}

// CopyLinkedList 拷贝一个一模一样的链表
func CopyLinkedList(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{Value: node.Value}
	oldCur := node
	newCur := newNode
	for oldCur.Next != nil {
		newCur.Next = &Node{Value: oldCur.Next.Value}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	return newNode
}

// CheckLinkedListEqual 校验两个链表是否全等
// 全等条件：长度相同，且每个位置的值相同
// 如果一个为nil，则另一个也必须为nil
func CheckLinkedListEqual(head1 *Node, head2 *Node) bool {
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

func GenerateLinkedListBySlice(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{Value: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &Node{
			Value: arr[i],
		}
		cur = cur.Next
	}
	return head
}

// DoubleNode 双向链表
type DoubleNode struct {
	Value int
	Last  *DoubleNode
	Next  *DoubleNode
}

// GenerateRandomDoubleList 初始化一个随机的双向链表
func GenerateRandomDoubleList(maxLength int, minValue int, maxValue int) *DoubleNode {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength + 1)
	// 如果长度是0，则返回空链表
	if length == 0 {
		return nil
	}
	// 先构造一个头节点
	head := &DoubleNode{Value: GenNodeValue(minValue, maxValue)}
	cur := head
	// 还剩length-1个节点
	for i := 1; i < length; i++ {
		value := GenNodeValue(minValue, maxValue)
		cur.Next = &DoubleNode{Value: value, Last: cur}
		cur = cur.Next
	}
	return head
}

// CopyDoubleList 拷贝双向链表
func CopyDoubleList(head *DoubleNode) *DoubleNode {
	if head == nil {
		return nil
	}
	newHead := &DoubleNode{Value: head.Value}
	oldCur := head
	newCur := newHead
	for oldCur.Next != nil {
		newCur.Next = &DoubleNode{Value: oldCur.Next.Value, Last: newCur}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	return newHead
}
