package class_09

// 链表荷兰国旗问题

// NetherlandsFlagList1 链表的荷兰国旗问题，使用辅助数组实现
//
//	@param head 原始头节点
//	@param pivot 划分值
//	@return newHead 新的头节点
func NetherlandsFlagList1(head *Node, pivot int) (newHead *Node) {
	var help []*Node // 准备一个辅助数组
	// 将链表中的节点全部放入数组中
	cur := head
	for cur != nil {
		help = append(help, cur)
		cur = cur.Next
	}
	// 执行help数组的荷兰国旗问题
	netherlandsFlagArr(help, pivot)
	// 再将数组中的额元素一个个串起来返回
	for i := 0; i < len(help); i++ {
		if i == 0 {
			newHead = help[i]
		}
		if i < len(help)-1 {
			help[i].Next = help[i+1]
		} else {
			help[i].Next = nil
		}
	}
	return newHead
}

func netherlandsFlagArr(arr []*Node, pivot int) {
	if len(arr) == 0 {
		return
	}
	L, R := -1, len(arr) // 小于区域和大于区域 	i
	i := 0
	// i与右侧区域碰上，就停止循环
	for i < R {
		if arr[i].Value < pivot {
			swap(arr, i, L+1)
			L++
			i++
		} else if arr[i].Value > pivot {
			swap(arr, i, R-1)
			R--
		} else {
			i++
		}
	}
}

func swap(arr []*Node, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

// NetherlandsFlagList2 荷兰国旗2，不使用数组实现
func NetherlandsFlagList2(head *Node, pivot int) (newHead *Node) {
	// 准备6个变量，分别表示小于区域头和尾，等于区域头和尾，大于区域头和尾，初始都是空
	var sH, sT *Node = nil, nil // 小于区域头尾
	var eH, eT *Node = nil, nil // 等于区域头尾
	var bH, bT *Node = nil, nil // 大于区域头尾
	// next 用来先记住下一个环境，因为当前环境可能会重新指向新的节点
	var next *Node
	// 如果当前节点有下一个节点，则继续
	for head.Next != nil {
		next = head.Next
		head.Next = nil // 当前节点必然会挂到某个区域，挂之前先和后面的断连
		// 比特定值小，挂到小于区域
		if head.Value < pivot {
			if sH == nil {
				sH = head
			} else {
				sT.Next = head
			}
			sT = head
		} else if head.Value == pivot {
			// 与特定值相等，挂到等于区域
			if eH == nil {
				eH = head
			} else {
				eT.Next = head
			}
			eT = head
		} else {
			// 比特定值大，挂到大于区域
			if bH == nil {
				bH = head
			} else {
				bT.Next = head
			}
			bT = head
		}
		head = next
	}
	// 全部走完，将小于区域尾部连接等于区域头部，等于区域尾部连接大于区域头部
	if sT != nil {
		sT.Next = eH
	}
	// 保证让eT去连接bH
	if eT == nil {
		eT = sT
	}
	if eT != nil {
		eT.Next = bH
	}
	// 返回新的头节点
	if sH != nil {
		newHead = sH
	} else if eH != nil {
		newHead = eH
	} else {
		newHead = bH
	}
	return
}
