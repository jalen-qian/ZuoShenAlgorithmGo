package class_09_10

import "fmt"

// Node 单向链表
type Node struct {
	Value int
	Next  *Node
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
