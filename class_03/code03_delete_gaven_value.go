package class_03

// 题目2：删除链表中给定的数
func deleteNumberInLinkedList(head *ListNode, value int) *ListNode {
	// 1. 找到第一个不等于number的节点作为头结点，并将前面的节点都释放
	for head != nil {
		if head.Value != value {
			break
		}
		head = head.Next
	}
	// 2. 接下来的链表中，如果遇到number，则删除
	// pre 和 cur都来到head的位置，head肯定不等于value
	pre := head
	cur := head
	for cur != nil {
		// 如果cur的值是给定的value，则跳过
		if cur.Value == value {
			pre.Next = cur.Next
		} else {
			// pre来到cur的位置
			pre = cur
		}
		// cur往下一个条
		cur = cur.Next
	}
	return head
}

// testDeleteNumberInLinkedList 对数器，分配额外的空间来实现
func testDeleteNumberInLinkedList(head *ListNode, number int) *ListNode {
	// 1. 生成一个节点数组
	nodeArr := make([]*ListNode, 0)
	// 2. 将链表中的节点填入数组中
	cur := head
	for cur != nil {
		nodeArr = append(nodeArr, cur)
		cur = cur.Next
	}
	// 3. 遍历数组，将是number的节点置为空
	for i, node := range nodeArr {
		if node.Value == number {
			nodeArr[i] = nil
		}
	}
	// 4. 将节点重新连接起来
	var foundHead bool
	var pre *ListNode
	nilCount := 0
	for _, node := range nodeArr {
		// 值为空了，说明是number被删除掉了，则跳过
		if node == nil {
			nilCount++
			continue
		}
		// 如果还没指定过头节点，则指定，并将pre设定到此位置
		if !foundHead {
			head = node
			pre = head
			foundHead = true
		} else {
			// 往下遍历，遇到的节点都置为pre的next
			pre.Next = node
			pre = node
		}
	}
	// 全删除了，则返回nil
	if nilCount == len(nodeArr) {
		return nil
	}
	return head
}
