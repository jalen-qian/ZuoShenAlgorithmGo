package class_09

// 找到两个链表第一个相交的节点

func GetIntersectNode(head1 *Node, head2 *Node) *Node {
	// 1. 先找到各自的入环节点，同时判断有环还是没环
	loop1 := getLoopNode(head1)
	loop2 := getLoopNode(head2)
	// 2. 如果都无环，则走无环逻辑
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}
	// 3. 两个都有环，走都有环的逻辑
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, loop1, head2, loop2)
	}
	// 4. 走到这，肯定是一个有环一个无环，则必定不相交，返回nil
	return nil
}

// getLoopNode 找到第一个入环节点，如果没有环，返回空
func getLoopNode(head *Node) *Node {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	fast, slow := head, head // 快慢指针都指向头
	for slow.Next != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		// 如果追上了slow，则fast回到开头，并停止
		if fast == slow {
			fast = head
			break
		}
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// noLoop 两个无环但可能相交的链表，找到第一个相交的节点
func noLoop(head1 *Node, head2 *Node) *Node {
	// 统计数量
	n := 0
	cur1 := head1
	for cur1.Next != nil {
		n++
		cur1 = cur1.Next
	}
	cur2 := head2
	for cur2.Next != nil {
		n--
		cur2 = cur2.Next
	}
	// n变成正数，并且cur1指向较长的
	if n >= 0 {
		cur1 = head1
		cur2 = head2
	} else {
		n = -n
		cur1 = head2
		cur2 = head1
	}
	// 先让cur1跳n步
	for i := 0; i < n; i++ {
		cur1 = cur1.Next
	}
	// 同步往下走，如果相遇，则返回相交的节点
	for cur1.Next != nil {
		if cur1 == cur2 {
			return cur1
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	// 走到底都没相遇，则说明不相交，返回空
	return nil
}

// bothLoop 两个有环的链表，找到第一个相交的节点
func bothLoop(head1 *Node, loop1 *Node, head2 *Node, loop2 *Node) *Node {
	// 入环节点相同，转换成noLoop问题
	if loop1 == loop2 {
		n := 0
		cur1 := head1
		for cur1.Next != loop1 {
			n++
			cur1 = cur1.Next
		}
		cur2 := head2
		for cur2.Next != loop2 {
			n--
			cur2 = cur2.Next
		}
		// n变成正数，并且cur1指向较长的
		if n >= 0 {
			cur1 = head1
			cur2 = head2
		} else {
			n = -n
			cur1 = head2
			cur2 = head1
		}
		// 先让cur1跳n步
		for i := 0; i < n; i++ {
			cur1 = cur1.Next
		}
		// 同步往下走，直到相遇，则返回相交的节点（一定会相遇）
		for cur1 != cur2 {
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else {
		cur1 := loop1.Next
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.Next
		}
		return nil
	}
}
