package class_07

import "ZuoShenAlgorithmGo/utils"

/**
实现一个大根堆
*/

type BigHeap struct {
	arr      []int // 存放堆的数组
	heapSize int   // 当前堆中数字的多少（堆的大小）
}

// Add 向大根堆中加入一个数
func (heap *BigHeap) Add(num int) {
	// 下一次要添加的数在的位置是 heapSize - 1
	index := heap.heapSize - 1
	heap.add(num, index)
	// 如果比父亲大，就一直和父亲交换，指导 <= 父亲或者到了根节点
	for (index-1)>>1 >= 0 && heap.arr[index] > heap.arr[(index-1)>>1] {
		utils.Swap(heap.arr, (index-1)/2, index)
		index = (index - 1) >> 1
	}
}

// add 向当前数组的i位置添加一个数
func (heap *BigHeap) add(num int, index int) {
	if index >= len(heap.arr) {
		heap.arr = append(heap.arr, num)
	}
	heap.arr[index] = num
	heap.heapSize++
}
