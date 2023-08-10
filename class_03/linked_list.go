package class_03

import (
	"fmt"
	"math/rand"
	"time"
)

// ListNode 单向链表
type ListNode struct {
	Value int
	Next  *ListNode
}

func GenNodeValue(minValue int, maxValue int) int {
	return rand.Intn(maxValue-minValue+1) + minValue
}

// GenerateRandomLinkedList 初始化一个随机的链表
func GenerateRandomLinkedList(maxLength int, minValue int, maxValue int) *ListNode {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength + 1)
	// 如果长度是0，则返回空链表
	if length == 0 {
		return nil
	}
	// 先构造一个头节点
	head := &ListNode{Value: GenNodeValue(minValue, maxValue)}
	cur := head
	// 还剩length-1个节点
	for i := 1; i < length; i++ {
		value := GenNodeValue(minValue, maxValue)
		cur.Next = &ListNode{Value: value}
		cur = cur.Next
	}
	return head
}

// GenerateRandomSortedLinkedList 随机生成一个有序的链表
// maxLength 最大链表长度
// minValue 最小值
// maxStep 递增的最大值，[0, maxStep]
func GenerateRandomSortedLinkedList(maxLength int, minValue int, maxStep int) *ListNode {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength + 1)
	if length == 0 {
		return nil
	}
	// 先构造一个头节点
	head := &ListNode{Value: rand.Intn(1001) + minValue}
	cur := head
	for i := 1; i < length; i++ {
		value := cur.Value + rand.Intn(maxStep+1)
		cur.Next = &ListNode{Value: value}
		cur = cur.Next
	}
	return head
}

// SPrintLinkedList 打印链表
func SPrintLinkedList(head *ListNode) string {
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
func CopyLinkedList(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	newNode := &ListNode{Value: node.Value}
	oldCur := node
	newCur := newNode
	for oldCur.Next != nil {
		newCur.Next = &ListNode{Value: oldCur.Next.Value}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	return newNode
}

// CheckLinkedListEqual 校验两个链表是否全等
// 全等条件：长度相同，且每个位置的值相同
// 如果一个为nil，则另一个也必须为nil
func CheckLinkedListEqual(head1 *ListNode, head2 *ListNode) bool {
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

func GenerateLinkedListBySlice(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Value: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{
			Value: arr[i],
		}
		cur = cur.Next
	}
	return head
}

// ListDoubleNode 双向链表
type ListDoubleNode struct {
	Value int
	Last  *ListDoubleNode
	Next  *ListDoubleNode
}

// GenerateRandomDoubleList 初始化一个随机的双向链表
func GenerateRandomDoubleList(maxLength int, minValue int, maxValue int) *ListDoubleNode {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength + 1)
	// 如果长度是0，则返回空链表
	if length == 0 {
		return nil
	}
	// 先构造一个头节点
	head := &ListDoubleNode{Value: GenNodeValue(minValue, maxValue)}
	cur := head
	// 还剩length-1个节点
	for i := 1; i < length; i++ {
		value := GenNodeValue(minValue, maxValue)
		cur.Next = &ListDoubleNode{Value: value, Last: cur}
		cur = cur.Next
	}
	return head
}

// CopyDoubleList 拷贝双向链表
func CopyDoubleList(head *ListDoubleNode) *ListDoubleNode {
	if head == nil {
		return nil
	}
	newHead := &ListDoubleNode{Value: head.Value}
	oldCur := head
	newCur := newHead
	for oldCur.Next != nil {
		newCur.Next = &ListDoubleNode{Value: oldCur.Next.Value, Last: newCur}
		oldCur = oldCur.Next
		newCur = newCur.Next
	}
	return newHead
}
