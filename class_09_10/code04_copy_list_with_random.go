package class_09_10

// NodeWithRandom 带有一个随机指针的单链表
type NodeWithRandom struct {
	Value int
	Next  *NodeWithRandom
	Rand  *NodeWithRandom
}

// CopyRandomList1 方式1，使用容器的方式
func CopyRandomList1(head *NodeWithRandom) *NodeWithRandom {
	helpMap := make(map[*NodeWithRandom]*NodeWithRandom)
	cur := head
	// 先将所有节点在map中一一创建出来
	for cur != nil {
		helpMap[cur] = &NodeWithRandom{Value: cur.Value}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		helpMap[cur].Next = helpMap[cur.Next]
		helpMap[cur].Rand = helpMap[cur.Rand]
		cur = cur.Next
	}
	return helpMap[head]
}

// CopyRandomList2 方式2，不使用容器
func CopyRandomList2(head *NodeWithRandom) *NodeWithRandom {
	if head == nil {
		return nil
	}
	cur := head
	var next *NodeWithRandom
	for cur != nil {
		// 先记住下一个环境
		next = cur.Next
		// 生成新的伴随节点，让cur指向它，值与当前节点相同
		cur.Next = &NodeWithRandom{Value: cur.Value}
		// 伴随节点的Next指向原始的下一个，实现插入到中间
		cur.Next.Next = next
		// cur跳下一个
		cur = next
	}
	// 将伴随节点的Rand指针调整对
	cur = head
	for cur != nil {
		if cur.Rand != nil {
			// 找到对应的Rand指针指向的节点，赋值给当前节点的Rand指针
			cur.Next.Rand = cur.Rand.Next
		}
		// 每次跳2步
		cur = cur.Next.Next
	}
	newHead := head.Next //新的头部

	// 将伴随节点从旧的链表中剥离开来，串成新链表，同时恢复原链表
	cur = head
	// 1->1'->2->2'
	for cur != nil {
		// 记住下一个环境 2
		next = cur.Next.Next
		// 当前节点的伴随节点Next，指向下一个环境的伴随节点 1'.Next -> 2'
		if next != nil {
			cur.Next.Next = next.Next
		}
		// 当前节点的Next，指向下一个环境 1->2
		cur.Next = next
		// cur来到2的位置
		cur = next
	}
	return newHead
}

// 检查拷贝得对不对，拷贝得对的标准：
// 1.原链表是空，新链表也是空
// 2.原链表的长度和新链表一致，且新链表不能是环结构（因为原链表也不是）
// 3.原链表遍历，每个节点的值与新链表遍历一一对应
// 4.原链表遍历，Rand指针指向的指针与新链表的Rand指向的位置也一一对应（注意要指向的位置也相同，而不仅仅是指向的值相同）
func checkCopyRight(oldHead *NodeWithRandom, newHead *NodeWithRandom) bool {
	if oldHead == nil && newHead != nil {
		return false
	}
	if oldHead != nil && newHead == nil {
		return false
	}
	if oldHead == nil && newHead == nil {
		return true
	}
	// 建立两个链表的节点一一对应关系
	helpMap := make(map[*NodeWithRandom]*NodeWithRandom)
	cur1 := oldHead
	cur2 := newHead
	for cur1 != nil {
		if cur2 == nil {
			return false
		}
		if cur1.Value != cur2.Value {
			return false
		}
		helpMap[cur1] = cur2
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	// 旧链表遍历完了，新链表还没完，说明长度不一致，或者搞成环了，也算拷贝不成功
	if cur2 != nil {
		return false
	}
	cur1 = oldHead
	cur2 = newHead
	// 通过map来单独检查rand指针的指向关系
	for curOld, curNew := range helpMap {
		if curOld.Rand == nil && curNew.Rand != nil {
			return false
		}
		if curOld.Rand != nil && curNew.Rand == nil {
			return false
		}
		// 如果rand不是空，那么指针指向的节点，也要是对应的节点
		if curOld.Rand != nil && helpMap[curOld.Rand] != curNew.Rand {
			return false
		}
	}
	return true
}
