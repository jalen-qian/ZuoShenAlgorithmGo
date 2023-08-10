package class_09

import (
	"ZuoShenAlgorithmGo/class_03"
)

// 判断一个链表是否是回文结构

// IsPalindrome1 使用容器实现（栈）
func IsPalindrome1(head *Node) bool {
	// 创建一个栈
	stack := class_03.NewMyStack[int]()
	cur := head
	// 遍历链表，将所有节点的值压入栈
	for cur != nil {
		stack.Push(cur.Value)
		cur = cur.Next
	}
	// 再次遍历链表，并一个个弹出栈，如果弹出的值与当前值不同，则不是回文
	cur = head
	for cur != nil {
		if cur.Value != stack.Pop() {
			return false
		}
		cur = cur.Next
	}
	// 每个都弹出一致，则是回文链表，null链表两个for都不走，也算回文链表
	return true
}

// IsPalindrome2 不使用容器实现
func IsPalindrome2(head *Node) bool {
	// 空链表或者只有1个节点，算回文结构
	if head == nil || head.Next == nil {
		return true
	}
	// 找到中点或者下中点，作为下半部分的开始
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		slow = slow.Next
	}
	// slow就是中点或者下中点，从slow开始反转链表 1->2->3->2->1  返回3 变成 1->2->3<-2<-1 翻转链表，同时让3指向空
	// 如果是偶数个， 1->2->3->3->2->1 下中点是第2个3，反转成 1->2->3->3<-2<-1 同时也让第2个3指向空
	// 极端情况，只有两个，比如 1->1  slow是最后的1，然后不会发生反转，下面的c2=pre指向最后的1，比对只会发生1次，然后c2变成nil，能跑通
	var pre, next *Node
	cur := slow
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 反转完，pre就是反转后的头节点，也是原链表的尾节点
	// 遍历反转后的两个节点，如果两个链表的next指针都不是空，则说明没遍历到中点位置，有1个为空或者两个都为空，则停止
	c1 := head
	c2 := pre
	isPalindrome := true // 记录是否是回文
	for c1 != nil && c2 != nil {
		if c1.Value != c2.Value {
			isPalindrome = false // 回文对不上，记录不是回文，直接break
			break
		} else {
			//否则，都往下一个跳
			c1 = c1.Next
			c2 = c2.Next
		}
	}
	// 最后，将链表反转恢复回来
	// 从pre的位置开始反转，pre此时的位置是原始链表的尾节点，反转后的第二个链表的头节点
	cur = pre
	pre, next = nil, nil
	next = cur.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return isPalindrome
}
