package class_09

import (
	"testing"
)

// 测试链表回文练习题

// 实现方式1，使用栈实现的方式
func TestIsPalindrome1(t *testing.T) {
	// 生成一个自定义的链表，1->2->3->2->1
	head, _ := generateRandomLinkListAndArr(0, 0, 0, 1, 2, 3, 2, 1)
	if !IsPalindrome1(head) {
		t.Errorf("测试失败，应该是回文，返回false")
	}
	t.Logf("执行后打印链表：%s\n", SPrintLinkedList(head))

	head, _ = generateRandomLinkListAndArr(0, 0, 0, 1, 20, 20, 1)
	if !IsPalindrome1(head) {
		t.Errorf("测试失败，应该是回文，返回false")
	}
	t.Logf("执行后打印链表：%s\n", SPrintLinkedList(head))

	head, _ = generateRandomLinkListAndArr(0, 0, 0, 1, 20, 30, 10, 1)
	if IsPalindrome1(head) {
		t.Errorf("测试失败，应该不是回文，返回true")
	}
	t.Logf("执行后打印链表：%s\n", SPrintLinkedList(head))
}

// 实现方式2，使用有限几个变量实现
func TestIsPalindrome2(t *testing.T) {
	// 生成一个自定义的链表，1->2->3->2->1
	head, _ := generateCustomLinkedListByArr([]int{0, 0, 0, 1, 2, 3, 2, 1})
	if !IsPalindrome2(head) {
		t.Errorf("测试失败，应该是回文，返回false")
	}
	t.Logf("执行后打印链表：%s\n", SPrintLinkedList(head))

	head, _ = generateCustomLinkedListByArr([]int{0, 0, 0, 1, 2, 3, 2, 1})
	if !IsPalindrome2(head) {
		t.Errorf("测试失败，应该是回文，返回false")
	}
	t.Logf("执行后打印链表：%s\n", SPrintLinkedList(head))

	head, _ = generateCustomLinkedListByArr([]int{0, 0, 0, 1, 2, 3, 2, 1})
	if IsPalindrome2(head) {
		t.Errorf("测试失败，应该不是回文，返回true")
	}
	t.Logf("执行后打印链表：%s\n", SPrintLinkedList(head))
}
