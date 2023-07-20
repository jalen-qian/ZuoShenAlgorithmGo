package class_06

// ============ 随机快排的非递归调用，使用栈来实现

// Op 用来记录每个子过程的区间，从L到R
type Op struct {
	L int
	R int
}

// QuickSort4 快排4.0，不用递归方式实现随机快排
func QuickSort4(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 初始化一个栈，用来存放子过程
	stack := MyStackWithLinkedList[Op]{}
	left, right := NetherlandsFlag5(nums, 0, len(nums)-1)
	// 先创建两个子过程，并放到栈中
	stack.Push(Op{L: 0, R: left - 1})
	stack.Push(Op{L: right + 1, R: len(nums) - 1})
	// 如果栈没有清空，则不断调用
	for !stack.IsEmpty() {
		// 取出一个子过程
		sub := stack.Pop()
		// 子过程没有>=两个数，则不用处理，直接跳下一个
		if sub.L >= sub.R {
			continue
		}
		// 子区间执行荷兰国旗问题，如果产生了更小的子区间，继续放到栈中
		left, right = NetherlandsFlag5(nums, sub.L, sub.R)
		stack.Push(Op{L: sub.L, R: left - 1})
		stack.Push(Op{L: right + 1, R: sub.R})
	}
}
