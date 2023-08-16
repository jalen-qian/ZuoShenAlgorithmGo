package class_09_10

import (
	"fmt"
	"testing"
)

func TestNetherlandsFlagList1(t *testing.T) {
	// 测试链表荷兰国旗版本1
	head, _ := generateCustomLinkedListByArr([]int{1, 5, 8, 3, 4, 2, 0, 6, 3, 1})
	fmt.Printf("荷兰国旗划分前的链表：%s\n", SPrintLinkedList(head))
	newHead := NetherlandsFlagList1(head, 3)
	fmt.Printf("荷兰国旗划分后的链表：%s\n", SPrintLinkedList(newHead))
	t.Log("测试成功")
}

func TestNetherlandsFlagList2(t *testing.T) {
	// 测试链表荷兰国旗版本1
	head, _ := generateCustomLinkedListByArr([]int{1, 5, 8, 3, 4, 2, 0, 6, 3, 1})
	fmt.Printf("荷兰国旗划分前的链表：%s\n", SPrintLinkedList(head))
	newHead := NetherlandsFlagList2(head, 3)
	fmt.Printf("荷兰国旗划分后的链表：%s\n", SPrintLinkedList(newHead))
	t.Log("测试成功")
}
