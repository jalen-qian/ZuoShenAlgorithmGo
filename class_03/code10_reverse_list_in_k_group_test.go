package main

import (
	"fmt"
	"testing"
)

func TestReverseListInKGroup(t *testing.T) {
	head := GenerateLinkedListBySlice([]int{1, 2, 3, 4, 5})
	ans := reverseKGroup(head, 3)
	fmt.Printf("结果:%s\n", SPrintLinkedList(ans))
	//head = GenerateLinkedListBySlice([]int{1, 2, 3, 4, 5})
	//ans = reverseKGroup(head, 2)
	//fmt.Printf("结果:%s\n", SPrintLinkedList(ans))
}
