package class_07

/**
实现一个大根堆，与Java的数组实现相比，这个大根堆没有长度限制，因为是用切片实现的，能动态扩展长度
*/

type MaxHeap struct {
	arr      []int // 存放堆的数组
	heapSize int   // 当前堆中数字的多少（堆的大小）
}

// Push 向大根堆中加入一个数
func (heap *MaxHeap) Push(num int) {
	// heapSize初始等于0，添加一个数后heapSize等于1，也就是说，heapSize就是下次要加入的数字的位置
	heap.add(num, heap.heapSize)
	// 执行heapInsert
	heapInsert(heap.arr, heap.heapSize)
	// 加入之后，heapSize++
	heap.heapSize++
}

// heapInsert
func heapInsert(arr []int, index int) {
	// 如果index位置比 (index - 1) /2 位置大，则交换
	// 如果index已经是0，(0 - 1) /2 也是0，不会被交换
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// Pop 从大根堆中弹出0位置的数（最大的数），剩下的数仍然维持大根堆
func (heap *MaxHeap) Pop() int {
	ans := heap.arr[0]
	swap(heap.arr, 0, heap.heapSize-1)
	heap.heapSize--
	heapify(heap.arr, 0, heap.heapSize)
	return ans
}

func heapify(arr []int, index int, heapSize int) {
	// 左孩子 index * 2 + 1 如果有右孩子，则右孩子是 left + 1
	left := index*2 + 1
	// 当前堆最后一个数的位置是 heapSize - 1
	for left < heapSize {
		// 比较左右孩子，找到较大的，将位置给largest
		largest := left
		// 如果有又孩子，且右孩子比左孩子大，则右孩子给largest
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		// 找到index和largest中谁最大，给largest
		if arr[index] > arr[largest] {
			largest = index
		}
		// 如果largest==index，说明左右孩子都不比当前位置的数大，则停止
		if largest == index {
			break
		}
		swap(arr, largest, index)
		index = largest
		// 找到下一个左孩子
		left = index*2 + 1
	}
}

// add 向当前数组的i位置添加一个数，如果超过切片长度，就扩展切片长度
func (heap *MaxHeap) add(num int, index int) {
	if index >= len(heap.arr) {
		heap.arr = append(heap.arr, num)
	}
	heap.arr[index] = num
}

func (heap *MaxHeap) IsEmpty() bool {
	return heap.heapSize == 0
}

// 交换
func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
