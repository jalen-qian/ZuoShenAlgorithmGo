package class_07

import (
	"ZuoShenAlgorithmGo/utils"
	"container/heap"
)

/*
已知一个几乎有序的数组，几乎有序是指，如果把数组调整排好序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说是比较小的。

请选择一个合适的排序策略，对这个数组进行排序。
*/

// MinHeap 使用系统的api实现一个小根堆
type MinHeap struct {
	arr []int
}

func (h *MinHeap) Len() int {
	return len(h.arr)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.arr[i] < h.arr[j] // 小的排前面，小根堆
}

// Swap 实现交换
func (h *MinHeap) Swap(i, j int) {
	tmp := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = tmp
}

func (h *MinHeap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

func (h *MinHeap) Pop() any {
	n := len(h.arr)
	ans := h.arr[n-1]
	h.arr = h.arr[0 : n-1]
	return ans
}

func SortedArrDistanceLessK(arr []int, k int) {
	index := 0
	minHeap := &MinHeap{arr: make([]int, 0)}
	// 从堆中弹出再写入arr的下个位置
	popIndex := 0
	// [0 ~ k]范围创建一个小根堆
	for ; index <= utils.Min(k, len(arr)-1); index++ {
		minHeap.Push(arr[index])
	}
	// k + 1位置开始，弹出一个写回原数组，再加入一个，直到所有数都加入完了
	for ; index < len(arr); index++ {
		// 弹出最小值，放到popIndex位置
		arr[popIndex] = minHeap.Pop().(int)
		popIndex++
		// 将index位置的数push到堆中
		minHeap.Push(arr[index])
	}
	// 最后将堆弹空并写入arr
	for minHeap.Len() > 0 {
		arr[popIndex] = minHeap.Pop().(int)
		popIndex++
	}
}

var _ heap.Interface = (*MinHeap)(nil)
