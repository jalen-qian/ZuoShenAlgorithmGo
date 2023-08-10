package class_03

/*
实现一个支持获取最小值的栈，除了实现栈基本功能外，额外提供一个函数，这个函数返回当前栈中最小的数
要求：时间复杂度O(1)，额外空间复杂度O(1)

思路：定义一个结构，内部有两个栈，假设为 s 栈 和 m 栈
s栈正常加入和弹出数字，m栈用来实现最小值，具体逻辑如下：
1. 加入数字a到栈时，如果s和m栈都是空的，则将数字同时加入到s栈和m栈
2. 加入数字a到栈时，如果s栈不是空的，则s栈正常压入，m栈判断栈顶的值是否 <= a，如果是，则继续压入一个栈顶的数，
如果不是，则压入数字a
3. 弹出数字时，s栈和m栈同时弹出
*/

// Comparable 定义可比较大小的泛型
type Comparable interface {
	int | int8 | int32 | int64 | int16 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

// GetMinStack 支持获取最小值的栈
// 注意，因为这里需要获取最小值，设计到比较大小，泛型不能是“any”了，需要限定可以比较大小的类型
type GetMinStack[T Comparable] struct {
	// 数据栈
	dataStack *MyStack[T]
	// 最小值栈
	minStack *MyStack[T]
}

func NewGetMinStack[T Comparable]() *GetMinStack[T] {
	return &GetMinStack[T]{
		dataStack: NewMyStack[T](),
		minStack:  NewMyStack[T](),
	}
}

func (s *GetMinStack[T]) Push(value T) {
	// 是空栈，则都压入
	if s.dataStack.IsEmpty() {
		s.dataStack.Push(value)
		s.minStack.Push(value)
	} else {
		// 不是空的，先获取最小值栈的栈顶
		// 如果栈顶的只比value大，则将value压入栈顶，否则将栈顶的值再次压入栈顶
		minValue := s.minStack.Peek()
		if value <= minValue {
			minValue = value
		}
		s.dataStack.Push(value)
		s.minStack.Push(minValue)
	}
}

func (s *GetMinStack[T]) Pop() T {
	// 弹出时，同时弹出
	s.minStack.Pop()
	return s.dataStack.Pop()
}

func (s *GetMinStack[T]) IsEmpty() bool {
	return s.dataStack.IsEmpty()
}

func (s *GetMinStack[T]) Size() int {
	return s.dataStack.Size()
}

func (s *GetMinStack[T]) Peek() T {
	return s.dataStack.Peek()
}

func (s *GetMinStack[T]) GetMin() T {
	return s.minStack.Peek()
}

var _ IStack[int] = (*GetMinStack[int])(nil)
