package utils

import (
	"math/rand"
	"time"
)

// Node 单向链表
type Node struct {
	Value int
	Next  *Node
}

// GenerateRandomLinkedList 初始化一个随机的链表
func GenerateRandomLinkedList(maxLength int, minValue int, maxValue int) *Node {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength + 1)
	var result *Node
	var cur *Node
	for i := 0; i < length; i++ {
		value := rand.Intn(maxValue-minValue+1) + minValue
		if result == nil {
			result = &Node{Value: value}
			cur = result
		} else {
			cur.Next = &Node{Value: value}
		}
	}
	return result
}
