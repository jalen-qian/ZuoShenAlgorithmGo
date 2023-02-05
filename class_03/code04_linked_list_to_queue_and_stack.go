package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
使用单链表实现栈和队列
*/

// Node 单链表节点，使用泛型，value可以接收任意类型
type Node[T any] struct {
	value T
	next  *Node[T]
}

// IQueue 队列接口，主要提供4种方法
type IQueue[T any] interface {
	Push(T)        // 入队
	Poll() T       // 出队
	IsEmpty() bool // 是否为空
	Size() int     // 返回队列大小
	Peek() T       // 只返回队尾的值，不实际弹出
}

// MyQueueWithLinkedList 使用单向链表实现队列
type MyQueueWithLinkedList[T any] struct {
	// 头指针
	head *Node[T]
	// 尾指针
	tail *Node[T]
	// 队列大小
	size int
}

// Push 加入值
func (q *MyQueueWithLinkedList[T]) Push(value T) {
	// 1. 先创建一个节点
	node := &Node[T]{value: value}
	// 2. 如果当前队列是空的，则将头指针和尾指针都指向此节点
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		// 3. 否则，将节点添加到尾指针的next，并将尾指针指向当前节点
		q.tail.next = node
		q.tail = node
	}
	// 4. 标记size + 1
	q.size++
}

// Poll 弹出值，如果没有值，则返回对应类型的0值
func (q *MyQueueWithLinkedList[T]) Poll() T {
	var ans T
	// 如果头指针不是空的，则返回头指针
	if q.head != nil {
		ans = q.head.value
		q.head = q.head.next
		q.size--
	}
	// 如果只剩一个值且被弹出了，则head会变成nil，此时标记尾指针也是nil
	if q.head == nil {
		q.tail = nil
	}
	return ans
}

// IsEmpty 返回队列是否为空
func (q *MyQueueWithLinkedList[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *MyQueueWithLinkedList[T]) Size() int {
	return q.size
}

func (q *MyQueueWithLinkedList[T]) Peek() T {
	var ans T
	// 如果头指针不是空的，则返回头指针
	if q.head != nil {
		ans = q.head.value
	}
	return ans
}

var _ IQueue[any] = (*MyQueueWithLinkedList[any])(nil)

// MyQueueWithSlice golang没有内置的栈和队列这种数据结构，这里为了对数器测试，使用切片实现一个相同功能的队列
// 这里同样使用泛型实现
type MyQueueWithSlice[T any] struct {
	items []T
}

func (q *MyQueueWithSlice[T]) Push(value T) {
	// 入队很简单，往后面拼接就好了
	q.items = append(q.items, value)
}

// Poll 出队
func (q *MyQueueWithSlice[T]) Poll() T {
	var ans T
	// 如果队列是空的，则直接返回
	if q.IsEmpty() {
		return ans
	}
	// 使用切片表达式删除最前面的元素
	ans = q.items[0]
	q.items = q.items[1:]
	return ans
}

func (q *MyQueueWithSlice[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *MyQueueWithSlice[T]) Size() int {
	return len(q.items)
}

func (q *MyQueueWithSlice[T]) Peek() T {
	var ans T
	if q.IsEmpty() {
		return ans
	}
	return q.items[0]
}

var _ IQueue[any] = (*MyQueueWithSlice[any])(nil)

// 使用单链表实现栈

// IStack 栈接口，提供
type IStack[T any] interface {
	Push(T)        // 入栈
	Pop() T        // 弹出
	IsEmpty() bool // 是否为空
	Size() int     // 返回栈大小
	Peek() T       // 只返回栈顶的值，不实际弹出
}

// MyStackWithLinkedList 使用单向链表实现的栈
type MyStackWithLinkedList[T any] struct {
	head *Node[T]
	size int
}

// Push 入栈
func (s *MyStackWithLinkedList[T]) Push(value T) {
	// 1. 先创建一个节点
	node := &Node[T]{value: value}
	// 2. 判断是否head是空点，如果是，说明栈是空的，让head指向此节点
	if s.head == nil {
		s.head = node
	} else {
		// 3. head不是空的，说明栈里面有值，则将node添加到链表的头部
		// node的next指向原来的头部
		node.next = s.head
		// 头部指向现在的节点
		s.head = node
	}
	s.size++
}

func (s *MyStackWithLinkedList[T]) Pop() T {
	var ans T
	if s.head == nil {
		return ans
	}
	// 弹出头部的元素
	ans = s.head.value
	// head指向head的下一个
	s.head = s.head.next
	s.size--
	return ans
}

func (s *MyStackWithLinkedList[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *MyStackWithLinkedList[T]) Size() int {
	return s.size
}

func (s *MyStackWithLinkedList[T]) Peek() T {
	var ans T
	if s.head != nil {
		return s.head.value
	}
	return ans
}

var _ IStack[any] = (*MyStackWithLinkedList[any])(nil)

// 使用切片实现一个栈

func main() {
	fmt.Println("测试开始")
	// 测试500000次
	testTimes := 500000
	// 每次测试操作的次数（push或者poll)100次
	oneTestOperatorNum := 100
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		// 创建一个队列，确定泛型类型为int类型
		queue1 := &MyQueueWithLinkedList[int]{}
		queue2 := &MyQueueWithSlice[int]{}
		for j := 0; j < oneTestOperatorNum; j++ {
			// 随机生成一个数
			value := rand.Intn(1001) - rand.Intn(1001) // [-1000, 1000]
			// 如果队列是空的，则必然加入队列
			if queue1.IsEmpty() {
				// 同时加入队列
				queue1.Push(value)
				queue2.Push(value)
			} else {
				// 队列不是空的，则50%的概率决定是入队还是出队
				p := rand.Float32()
				if p < 0.5 {
					// 同时加入队列
					queue1.Push(value)
					queue2.Push(value)
				} else {
					// 同时弹出，并判断弹出的数是否相同
					ans1 := queue1.Poll()
					ans2 := queue2.Poll()
					if ans1 != ans2 {
						fmt.Printf("出错了，队列1弹出%d,队列2弹出%d\n", ans1, ans2)
						return
					}
				}
			}
		}
	}
	fmt.Println("测试结束")
}
