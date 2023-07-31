package class_07

import (
	"container/heap"
	"math/rand"
	"testing"
	"time"
)

// 根据系统提供的接口，实现一个大根堆
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// Less 定义比较器，大根堆
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap 交换
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 压入
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop 弹出
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TestMaxHeap 测试大根堆，同时熟悉下系统的堆的用法
func TestMaxHeap(t *testing.T) {
	myMaxHeap := MaxHeap{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intHeap := make(IntHeap, 0)
	for i := 0; i < 100000; i++ {
		// 生成一个随机数 [-1000, 1000]
		randomNum := r.Intn(1001) - r.Intn(1001)
		randFloat := r.Float64()
		if randFloat < 0.33 {
			// 1/3 概率加入随机数
			myMaxHeap.Push(randomNum)
			heap.Push(&intHeap, randomNum)
		} else if randFloat < 0.66 {
			// 1/3 概率弹出一个数（如果堆不是空的）
			var myHeapPop int
			var intHeapPop int
			if !myMaxHeap.IsEmpty() {
				myHeapPop = myMaxHeap.Pop()
			}
			if intHeap.Len() != 0 {
				intHeapPop = heap.Pop(&intHeap).(int)
			}
			if myHeapPop != intHeapPop {
				t.Errorf("测试失败，myMaxHeap弹出的数是：%d, 系统堆弹出的数是：%d", myHeapPop, intHeapPop)
				break
			}
		} else {
			// 1/3 概率判断是否为空
			myHeapIsEmpty := myMaxHeap.IsEmpty()
			intHeapIsEmpty := intHeap.Len() == 0
			if myHeapIsEmpty != intHeapIsEmpty {
				t.Errorf("测试失败，myMaxHeap是空的：%v, 系统堆是空的：%v", myHeapIsEmpty, intHeapIsEmpty)
				break
			}
		}
	}
}

// TestMyHeap 测试泛型的堆
func TestMyHeap(t *testing.T) {
	// 泛型的大根堆
	myMaxHeap := NewMyHeap[int](func(a int, b int) bool {
		return a > b
	})
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intHeap := make(IntHeap, 0)
	for i := 0; i < 100000; i++ {
		// 生成一个随机数 [-1000, 1000]
		randomNum := r.Intn(1001) - r.Intn(1001)
		randFloat := r.Float64()
		if randFloat < 0.33 {
			// 1/3 概率加入随机数
			myMaxHeap.Push(randomNum)
			heap.Push(&intHeap, randomNum)
		} else if randFloat < 0.66 {
			// 1/3 概率弹出一个数（如果堆不是空的）
			var myHeapPop int
			var intHeapPop int
			if !myMaxHeap.IsEmpty() {
				myHeapPop = myMaxHeap.Pop()
			}
			if intHeap.Len() != 0 {
				intHeapPop = heap.Pop(&intHeap).(int)
			}
			if myHeapPop != intHeapPop {
				t.Errorf("测试失败，myMaxHeap弹出的数是：%d, 系统堆弹出的数是：%d", myHeapPop, intHeapPop)
				break
			}
		} else {
			// 1/3 概率判断是否为空
			myHeapIsEmpty := myMaxHeap.IsEmpty()
			intHeapIsEmpty := intHeap.Len() == 0
			if myHeapIsEmpty != intHeapIsEmpty {
				t.Errorf("测试失败，myMaxHeap是空的：%v, 系统堆是空的：%v", myHeapIsEmpty, intHeapIsEmpty)
				break
			}
		}
	}
}
