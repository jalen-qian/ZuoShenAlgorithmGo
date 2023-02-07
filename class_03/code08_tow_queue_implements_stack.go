package main

/*
使用两个队列实现栈结构
思路：准备两个队列，queue 和 help，初始时都是空的
当入栈时，将数不断加入queue，比如先后压入 1 2 3 4，则
queue 1->2->3->4
help
当要弹出时，此时需要弹出4，我们需要将前3个数先倒入help队列，方法是queue出来的数，help加入，直到queue只剩一个数，变成
queue 4
help  1->2->3
然后将queue的数出队，并将queue和help调换
queue
help  1->2->3 调换后
queue 1->2->3
help
依此类推，每次需要弹出数时，都倒一次数据，push的时间复杂度：O(1) pop的时间复杂度 O(N)
基本思想是两个队列互相倒数据，直到只剩下队头的数据，就是要弹出的
*/

type TwoQueueStack[T any] struct {
	queue IQueue[T]
	help  IQueue[T]
}

func (s *TwoQueueStack[T]) Push(value T) {
	// 入栈时，只向queue入队
	s.queue.Push(value)
}

func (s *TwoQueueStack[T]) Pop() T {
	// 弹出，先将前 n-1 个数倒入help队列
	for s.queue.Size() > 1 {
		s.help.Push(s.queue.Poll())
	}
	var ans T
	// queue队列只剩一个数，弹出
	ans = s.queue.Poll()
	// 将help和queue交换
	tmp := s.queue
	s.queue = s.help
	s.help = tmp
	return ans
}

func (s *TwoQueueStack[T]) IsEmpty() bool {
	return s.queue.IsEmpty()
}

func (s *TwoQueueStack[T]) Size() int {
	return s.queue.Size()
}

func (s *TwoQueueStack[T]) Peek() T {
	// Peek的逻辑和Pop一致，区别是最后队列头的数不弹出，而是继续加入help队列
	// 先将前 n-1 个数倒入help队列
	for s.queue.Size() > 1 {
		s.help.Push(s.queue.Poll())
	}
	var ans T
	// 将这个数弹出后，加入到help
	ans = s.queue.Poll()
	s.help.Push(ans)
	// 将help和queue交换
	tmp := s.queue
	s.queue = s.help
	s.help = tmp
	return ans
}

func NewTwoQueueStack[T any]() *TwoQueueStack[T] {
	return &TwoQueueStack[T]{
		queue: NewMyQueue[T](),
		help:  NewMyQueue[T](),
	}
}

// 确保实现栈接口
var _ IStack[any] = (*TwoQueueStack[any])(nil)
