package class_07

import "sort"

/*
堆排序
*/

func HeapSort(arr []int) {
	// 1. 先将整个数组调整成大根堆
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	// 2. 不断将0位置与堆最后的位置进行交换，并重新调整成大根堆
	heapSize := len(arr)
	for heapSize > 0 {
		// 将0位置与堆的最后位置交换
		swap(arr, 0, heapSize-1)
		// 堆大小减1，固定最后位置的数
		heapSize--
		heapify(arr, 0, heapSize)
	}
}

// Comparator 对数器，系统排序
func Comparator(arr []int) {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
