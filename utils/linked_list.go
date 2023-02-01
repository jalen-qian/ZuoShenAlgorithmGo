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

func genNodeValue(minValue int, maxValue int) int {
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
	head := &Node{Value: genNodeValue(minValue, maxValue)}
	cur := head
	// 还剩length-1个节点
	for i := 1; i < length; i++ {
		value := genNodeValue(minValue, maxValue)
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
		ans += fmt.Sprintf("%d->", cur.Value)
		cur = cur.Next
	}
	ans += "null}"
	return ans
}
