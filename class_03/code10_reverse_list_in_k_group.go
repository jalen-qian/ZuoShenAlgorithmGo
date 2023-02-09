package main

/*
leetCode第25题：K 个一组反转链表
给你链表的头结点 head ,每k个节点一组进行翻转，返回修改后的链表
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
实例1：k = 2
1->2->3->4->5->null
2->1->4->3->5->null
实例2：k = 3
1->2->3->4->5->null
3->2->1->4->5->null
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-nodes-in-k-group

思路：反转单项链表的算法已经实现过，这里需要分成多个小组，每个小组反转后，还能和前后的小组头尾相接
1. 首先实现一个函数，传入一个start节点，返回小组的end节点，如果不够了，则返回空。比如实例2中，传入
	节点1，则返回节点3，传入节点4则返回空（因为4、5不够一组）
2. 首先获取第一个尾结点，如果获取到nil，说明整个链表都不够一组，则直接返回head，如果获取到，则这个
	尾结点就是整个要返回链表的头结点，将head指向这个节点。

*/

// findGroupEnd 找到分组的尾结点
func findGroupEnd(start *ListNode, k int) *ListNode {
	var end *ListNode
	cur := start
	for ; k > 0; k-- {
		if cur != nil {
			end = cur
			cur = cur.Next
		} else {
			end = nil
		}
	}
	return end
}

// reverseKGroup K个一组反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 先找到第一个分组的end节点，作为整个新链表的头结点
	end := findGroupEnd(head, k)
	// 如果返回nil，说明整个链表都不够一组，直接返回当前头结点
	if end == nil {
		return head
	}
	// 分组开始的节点，最开始是头结点
	start := head
	// 找到第一个分组的结束节点后，由于会反转，这个结束节点就会是最终的头部
	head = end
	// 用户反转链表的pre和next以及cur
	var pre *ListNode
	var next *ListNode
	// cur先来到第一个分组的头结点
	cur := start
	// preEnd用来记录上一个分组的尾部，对于第一个分组来说，尾部就是头结点（会反转）
	preEnd := start
	// 2. 不断找分组的end节点，如果不是空的，说明找到一个分组
	for end != nil {
		// endNext用来先记住下一个分组的头结点，反转后好连过去
		endNext := end.Next
		// 反转组内链表，一共进行次数是k次
		pre = nil
		for i := 0; i < k; i++ {
			// 先记住下一个节点
			next = cur.Next
			// 当前节点往前指
			cur.Next = pre
			// pre来到当前节点
			pre = cur
			// 当前来到下一个位置
			cur = next
		}
		// 当前分组的最后一个节点（反转后start来到尾部）指向下一个分组的头节点
		start.Next = endNext
		// 上一个分组的尾结点指向当前分组的头结点（反转后end变成头结点）
		if preEnd != nil {
			preEnd.Next = end
		}
		// preEnd跳到当前分组的尾部
		preEnd = start
		// start来到下一个分组的头结点
		start = endNext
		// end 继续获取下一个分组的尾结点
		end = findGroupEnd(start, k)
	}
	return head
}
