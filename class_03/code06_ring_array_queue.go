package class_03

import "errors"

/*
使用环状结构的数组实现队列
基本思路：定义一个数组，并有两个下标，头下标和尾下标，分别表示下一个要弹出的数的位置，和要添加的数的位置。
还有一个size变量记录当前队列中的数，初始状态下数组是空的，头坐标追赶尾部坐标
*/

type IQueueErr[T any] interface {
	Push(T) error     // 入队
	Poll() (T, error) // 出队
	IsEmpty() bool    // 是否为空
	Size() int        // 返回队列大小
	Peek() T          // 只返回队尾的值，不实际弹出
}

// RingArrayQueue 环形数组实现的队列
type RingArrayQueue[T any] struct {
	items     []T // 数组
	size      int // 队列大小
	headIndex int // 头坐标，标记“下一次要弹出数的位置”
	tailIndex int // 尾坐标，标记“下一次要加入数的位置”
	limit     int
}

func NewRingArrayQueue[T any](limit int) *RingArrayQueue[T] {
	return &RingArrayQueue[T]{
		items: make([]T, limit),
		limit: limit,
	}
}

func (r *RingArrayQueue[T]) Push(value T) error {
	if r.size == r.limit {
		return errors.New("队列已满，无法加入数据了")
	}
	// 往 endIndex 处添加数
	r.items[r.tailIndex] = value
	// 如果当前 endIndex 已经在数组最后了，则跳到0位置
	if r.tailIndex == r.limit-1 {
		r.tailIndex = 0
	} else {
		r.tailIndex++
	}
	r.size++
	return nil
}

func (r *RingArrayQueue[T]) Poll() (T, error) {
	// 如果队列空了，也报错
	var ans T
	if r.size == 0 {
		return ans, errors.New("队列是空的，无法弹出")
	}
	// 队列不是空的，则将 headIndex 处的值返回
	ans = r.items[r.headIndex]
	if r.headIndex == r.limit-1 {
		r.headIndex = 0
	} else {
		r.headIndex++
	}
	r.size--
	return ans, nil
}

func (r *RingArrayQueue[T]) IsEmpty() bool {
	return r.size == 0
}

func (r *RingArrayQueue[T]) Size() int {
	return r.size
}

func (r *RingArrayQueue[T]) Peek() T {
	var ans T
	if r.IsEmpty() {
		return ans
	}
	return r.items[r.headIndex]
}

var _ IQueueErr[any] = (*RingArrayQueue[any])(nil)
