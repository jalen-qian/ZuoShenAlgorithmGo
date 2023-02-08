package main

/*
实现一个双向队列，并通过双向队列实现栈和队列
双向队列：可以从头加，从头弹出，也可以从尾部加，从尾部弹出的队列
*/

// IDoubleEndsQueue 双向队列接口
type IDoubleEndsQueue[T any] interface {
	AddFromHead(value T)   // 从头部加
	PopFromHead() T        // 从头部弹出
	AddFromBottom(value T) // 从尾部加
	PopFromBottom() T      // 从尾部弹出
	IsEmpty() bool         // 返回是否为空
	PeekFromHead() T       // 返回头部的值，但是不弹出
	PeekFromBottom() T     // 返回尾部的值，但是不弹出
}

// ListDoubleNode 双链表
type DoubleNode[T any] struct {
	value T
	next  *DoubleNode[T]
	last  *DoubleNode[T]
}

// DoubleEndsQueue 双向队列
type DoubleEndsQueue[T any] struct {
	head *DoubleNode[T]
	tail *DoubleNode[T]
	size int
}

// AddFromHead 从头部加入
func (q *DoubleEndsQueue[T]) AddFromHead(value T) {
	node := &DoubleNode[T]{value: value}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		node.next = q.head
		q.head.last = node
		q.head = node
	}
	q.size++
}

// PopFromHead 从头部弹出
func (q *DoubleEndsQueue[T]) PopFromHead() T {
	var ans T
	if q.head == nil {
		return ans
	}
	ans = q.head.value
	// head往下跳一个节点
	q.head = q.head.next
	// 如果跳到了空，说明队列空了，则要将尾指针也置空
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return ans
}

func (q *DoubleEndsQueue[T]) AddFromBottom(value T) {
	node := &DoubleNode[T]{value: value}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		node.last = q.tail
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

// PopFromBottom 从尾部弹出
func (q *DoubleEndsQueue[T]) PopFromBottom() T {
	var ans T
	if q.tail == nil {
		return ans
	}
	ans = q.tail.value
	// q.tail尾指针往前跳一个节点
	q.tail = q.tail.last
	if q.tail == nil {
		q.head = nil
	}
	q.size--
	return ans
}

func (q *DoubleEndsQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *DoubleEndsQueue[T]) PeekFromHead() T {
	var ans T
	if q.head != nil {
		ans = q.head.value
	}
	return ans
}

func (q *DoubleEndsQueue[T]) PeekFromBottom() T {
	var ans T
	if q.tail != nil {
		ans = q.tail.value
	}
	return ans
}

var _ IDoubleEndsQueue[any] = (*DoubleEndsQueue[any])(nil)

// MyQueue 使用双向队列实现队列
type MyQueue[T any] struct {
	doubleEndsQueue *DoubleEndsQueue[T]
}

// NewMyQueue 构造函数
func NewMyQueue[T any]() *MyQueue[T] {
	return &MyQueue[T]{
		doubleEndsQueue: &DoubleEndsQueue[T]{},
	}
}

func (q *MyQueue[T]) Push(value T) {
	// 入队从头入
	q.doubleEndsQueue.AddFromHead(value)
}

func (q *MyQueue[T]) Poll() T {
	// 出队从尾部出
	return q.doubleEndsQueue.PopFromBottom()
}

func (q *MyQueue[T]) IsEmpty() bool {
	return q.doubleEndsQueue.IsEmpty()
}

func (q *MyQueue[T]) Size() int {
	return q.doubleEndsQueue.size
}

func (q *MyQueue[T]) Peek() T {
	// 返回队列尾部的值
	return q.doubleEndsQueue.PeekFromBottom()
}

var _ IQueue[any] = (*MyQueue[any])(nil)

// MyStack 使用双向队列实现栈
type MyStack[T any] struct {
	doubleEndsQueue *DoubleEndsQueue[T]
}

// NewMyStack 构造函数
func NewMyStack[T any]() *MyStack[T] {
	return &MyStack[T]{
		doubleEndsQueue: &DoubleEndsQueue[T]{},
	}
}

func (s *MyStack[T]) Push(value T) {
	// 入队从头入
	s.doubleEndsQueue.AddFromHead(value)
}

func (s *MyStack[T]) Pop() T {
	// 出队也从头部出
	return s.doubleEndsQueue.PopFromHead()
}

func (s *MyStack[T]) IsEmpty() bool {
	return s.doubleEndsQueue.IsEmpty()
}

func (s *MyStack[T]) Size() int {
	return s.doubleEndsQueue.size
}

func (s *MyStack[T]) Peek() T {
	// 返回队列头部的值
	return s.doubleEndsQueue.PeekFromHead()
}

var _ IStack[any] = (*MyStack[any])(nil)
