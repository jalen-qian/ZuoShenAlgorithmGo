package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		// 测试方式：反转两次，测试反转后的结果是否与原始链表全等
		linkedList := utils.GenerateRandomLinkedList(1000, -100, 100)
		copyList := utils.CopyLinkedList(linkedList)
		reversedList := reverseLinkedList(linkedList)
		reversedList2 := testReverseLinkedList(copyList)
		if !utils.CheckLinkedListEqual(reversedList, reversedList2) {
			fmt.Println("Fucking fucked!!!")
			fmt.Printf("原始链表：%s\n", utils.SPrintLinkedList(copyList))
			fmt.Printf("反转链表：%s\n", utils.SPrintLinkedList(reversedList))
			fmt.Printf("对数器链表：%s\n", utils.SPrintLinkedList(reversedList2))
			return
		}
	}
	fmt.Println("Great!!!")
}
