package main

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		// 测试方式：反转两次，测试反转后的结果是否与原始链表全等
		linkedList := GenerateRandomLinkedList(1000, -100, 100)
		copyList := CopyLinkedList(linkedList)
		reversedList := reverseLinkedList(linkedList)
		reversedList2 := testReverseLinkedList(copyList)
		if !CheckLinkedListEqual(reversedList, reversedList2) {
			fmt.Println("Fucking fucked!!!")
			fmt.Printf("原始链表：%s\n", SPrintLinkedList(copyList))
			fmt.Printf("反转链表：%s\n", SPrintLinkedList(reversedList))
			fmt.Printf("对数器链表：%s\n", SPrintLinkedList(reversedList2))
			return
		}
	}
	fmt.Println("Great!!!")
}
