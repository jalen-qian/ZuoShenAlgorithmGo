package class_07

import (
	"ZuoShenAlgorithmGo/utils"
)

/*
已知一个几乎有序的数组，几乎有序是指，如果把数组调整排好序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说是比较小的。

请选择一个合适的排序策略，对这个数组进行排序。
*/

// SortedArrDistanceLessK 几乎有序数组排序
func SortedArrDistanceLessK(arr []int, k int) {
	// 初始化一个小根堆
	minHeap := NewMyHeap[int](func(a int, b int) bool {
		return a < b
	})
	index := 0
	// 从堆中弹出再写入arr的下个位置
	popIndex := 0
	// [0 ~ k]范围创建一个小根堆
	for ; index <= utils.Min(k, len(arr)-1); index++ {
		minHeap.Push(arr[index])
	}
	// k + 1位置开始，弹出一个写回原数组，再加入一个，直到所有数都加入完了
	for ; index < len(arr); index++ {
		// 弹出最小值，放到popIndex位置
		arr[popIndex] = minHeap.Pop()
		popIndex++
		// 将index位置的数push到堆中
		minHeap.Push(arr[index])
	}
	// 最后将堆弹空并写入arr
	for !minHeap.IsEmpty() {
		arr[popIndex] = minHeap.Pop()
		popIndex++
	}
}
