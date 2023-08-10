package class_03

/*
使用两个栈实现队列
思路：和使用两个队列实现栈类似，使用两个栈也可以实现一个队列
同样准备 stack 和 help 两个栈
1. 当加入数据时，stack 压入数据，假如压入了 1 2 3 4 5，共5个数，则情况为（注意右边是栈顶）
stack 1 2 3 4 5
help
2. 当要出队时，依次将4个数弹出，压入help
stack 1
help  5 4 3 2
并将stack中的数1弹出，之后，将help中的数再次批量压入stack中
stack 2 3 4 5
help

3.假如调用的是Peek()，则直接将stack中的1Peek返回后，直接将help中的数批量倒入stack
*/

// TwoStackQueue 使用两个栈实现的队列
type TwoStackQueue[T any] struct {
	stack IStack[T]
	help  IStack[T]
}

func NewTwoStackQueue[T any]() *TwoStackQueue[T] {
	return &TwoStackQueue[T]{
		stack: NewMyStack[T](),
		help:  NewMyStack[T](),
	}
}

func (q *TwoStackQueue[T]) Push(value T) {
	// 加入数据时，正常往stack中压入数据
	q.stack.Push(value)
}

func (q *TwoStackQueue[T]) Poll() T {
	var ans T
	// 将栈顶的前 size -1 个数弹出来，压入help栈
	for q.stack.Size() > 1 {
		q.help.Push(q.stack.Pop())
	}
	// 要返回的数就是 stack 中的最后一个数
	ans = q.stack.Pop()
	// 最后将help中的所有数再次批量压入stack中
	for !q.help.IsEmpty() {
		q.stack.Push(q.help.Pop())
	}
	return ans
}

func (q *TwoStackQueue[T]) IsEmpty() bool {
	return q.stack.IsEmpty()
}

func (q *TwoStackQueue[T]) Size() int {
	return q.stack.Size()
}

func (q *TwoStackQueue[T]) Peek() T {
	var ans T
	// 将栈顶的前 size -1 个数弹出来，压入help栈
	for q.stack.Size() > 1 {
		q.help.Push(q.stack.Pop())
	}
	// 要返回的数就是 stack 中的最后一个数
	ans = q.stack.Peek()
	// 最后将help中的所有数再次批量压入stack中
	for !q.help.IsEmpty() {
		q.stack.Push(q.help.Pop())
	}
	return ans
}

var _ IQueue[any] = (*TwoStackQueue[any])(nil)
