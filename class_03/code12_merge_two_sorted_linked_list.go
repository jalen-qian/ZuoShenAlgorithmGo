package main

/**
两个有序链表的合并
给定两个有序链表，合并成一个有序链表
示例：
2->4->5->8->null
1->3->7->9->null
最后返回
1->2->3->4->5->7->8->9->null

思路：
1. 如果两个链表中，有一个是空，则返回非空的那个
2. 两个链表中，头部较小的那个head，是最终要返回的head，先记住这个head并最后返回
3. 定义一个pre指针，来到head的next位置
*/

// mergeTwoSortedLinkedList 算法实现
func mergeTwoSortedLinkedList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	// 找到小的那个，作为最终的头结点
	var head *ListNode = nil
	var cur2 *ListNode = nil
	if head1.Value < head2.Value {
		head = head1
		cur2 = head2
	} else {
		head = head2
		cur2 = head1
	}
	// cur1在head的下一个位置
	cur1 := head.Next
	// 定义一个pre节点，来到head的位置
	pre := head
	for cur1 != nil && cur2 != nil {
		if cur1.Value < cur2.Value {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	if cur1 != nil {
		pre.Next = cur1
	} else {
		pre.Next = cur2
	}
	return head
}
