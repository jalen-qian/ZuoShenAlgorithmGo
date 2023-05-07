package main

import (
	"sort"
	"testing"
)

// for test
// 实现思路：遍历两个链表，将数全部取出放在数组里面，然后对数组排序，最后根据这个数组
// 再生成一个链表
func testMergeTwoSortedLinkedList(head1 *ListNode, head2 *ListNode) *ListNode {
	arr := make([]int, 0)
	cur := head1
	for cur != nil {
		arr = append(arr, cur.Value)
		cur = cur.Next
	}
	cur = head2
	for cur != nil {
		arr = append(arr, cur.Value)
		cur = cur.Next
	}
	// 将数组进行排序
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	// 重新生成一个链表
	return GenerateLinkedListBySlice(arr)
}

func TestMergeTwoSortedLinkedList(t *testing.T) {
	// 测试次数：100000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		head1 := GenerateRandomSortedLinkedList(100, 10, 10)
		head2 := GenerateRandomSortedLinkedList(100, 10, 10)
		head1Copy := CopyLinkedList(head1)
		head2Copy := CopyLinkedList(head2)
		merged1 := mergeTwoSortedLinkedList(head1, head2)
		merged2 := testMergeTwoSortedLinkedList(head1Copy, head2Copy)
		if !CheckLinkedListEqual(merged1, merged2) {
			t.Fatalf("出错了，合并后不相等！！！")
		}
	}
	t.Logf("测试通过！！！\n")
}
