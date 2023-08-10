package class_09

/**
快慢指针相关面试题
*/

// MidOrUpMidNode 输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func MidOrUpMidNode(head *Node) *Node {
	fast, slow := head, head // 快慢两个指针，初始都指向头节点
	// fast一直往下跳到尾部
	for fast != nil && fast.Next != nil {
		// 如果能往下跳两步，就跳两步，否则跳1步
		if fast.Next.Next != nil {
			fast = fast.Next.Next
			// 快指针能走两步，则慢指针走1步，就一定是中点或者上中点停止
			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}
	return slow
}

// MidOrDownMidNode 输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func MidOrDownMidNode(head *Node) *Node {
	fast, slow := head, head // 快慢两个指针，初始都指向头节点
	// fast一直往下跳到尾部
	for fast != nil && fast.Next != nil {
		// 如果能往下跳两步，就跳两步，否则跳1步
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		// 不管快指针走1步还是两步，慢指针都走1步，就一定是中点或者下中点停止
		slow = slow.Next
	}
	// 1->2->3->4 ==> 返回3的指针
	return slow
}

// MidOrUpMidPreNode 输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
// 1->2->3 返回1   1->2 返回null   1->2->3->4 返回1
func MidOrUpMidPreNode(head *Node) *Node {
	// 如果节点个数少于3个，则一定是空的
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	// 至少有3个节点
	fast, slow := head.Next.Next, head // 快慢两个指针，快指针先跳两步，慢指针指向头
	// fast一直往下跳到尾部
	for fast.Next != nil {
		// 如果快指针能往下跳两步，就跳两步，否则跳1步
		if fast.Next.Next != nil {
			fast = fast.Next.Next
			// 快指针能跳两步，慢指针才跳1步
			// 全部跳完后，因为快指针先跳了1步，所以慢指针一定在是中点或者上中点前的一个节点
			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}
	return slow
}

// MidOrDownMidPreNode 输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
// 1->2 返回1  1->2->3 返回1  1->2->3->4 返回2
func MidOrDownMidPreNode(head *Node) *Node {
	// 如果节点个数少于两个，则一定是空的
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head // 快慢两个指针，都先指向头部
	// 如果有3个节点，则快指针先走两步
	if fast.Next.Next != nil {
		fast = fast.Next.Next
	}
	// fast一直往下跳到尾部
	for fast.Next != nil {
		// 如果快指针能往下跳两步，就跳两步，否则跳1步
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
		// 不管快指针走1步还是两步，慢指针都走1步，就一定是中点或者下中点停止
		// 由于有3个以及以上节点时，快指针先走了2步，所以慢指针会落在下中点或者中点前一个节点
		slow = slow.Next
	}
	return slow
}
