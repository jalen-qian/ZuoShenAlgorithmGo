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

}

// process4 让nums在L到R上排好序，并且不使用递归
func process4(nums []int, L, R int) {
	// L>=R说明要处理的区域没有数或者只有一个数，直接返回
	if L >= R {
		return
	}
	// 初始化一个栈，用来存放子过程
	stack := MyStackWithLinkedList[Op]{}

}
