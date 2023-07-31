package class_07

type MyComparator[T any] func(a T, b T) bool

// MyHeap 实现一个泛型的堆，能接收任意类型
type MyHeap[T any] struct {
	arr        []T // 存放堆的数组
	heapSize   int // 当前堆中数字的多少（堆的大小）
	comparator MyComparator[T]
}

func NewMyHeap[T any](comparator MyComparator[T]) *MyHeap[T] {
	return &MyHeap[T]{
		comparator: comparator,
	}
}

// Push 向大根堆中加入一个数
func (heap *MyHeap[T]) Push(num T) {
	// heapSize初始等于0，添加一个数后heapSize等于1，也就是说，heapSize就是下次要加入的数字的位置
	heap.add(num, heap.heapSize)
	// 执行heapInsert
	heap.heapInsert(heap.arr, heap.heapSize)
	// 加入之后，heapSize++
	heap.heapSize++
}

// heapInsert
func (heap *MyHeap[T]) heapInsert(arr []T, index int) {
	// 如果index位置比 (index - 1) /2 排前面，则交换
	// 如果index已经是0，(0 - 1) /2 也是0，不会被交换
	for heap.comparator(arr[index], arr[(index-1)/2]) {
		heap.swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// Pop 从大根堆中弹出0位置的数（最大的数），剩下的数仍然维持大根堆
func (heap *MyHeap[T]) Pop() T {
	ans := heap.arr[0]
	heap.swap(heap.arr, 0, heap.heapSize-1)
	heap.heapSize--
	heap.heapify(heap.arr, 0, heap.heapSize)
	return ans
}

func (heap *MyHeap[T]) heapify(arr []T, index int, heapSize int) {
	// 左孩子 index * 2 + 1 如果有右孩子，则右孩子是 left + 1
	left := index*2 + 1
	// 当前堆最后一个数的位置是 heapSize - 1
	for left < heapSize {
		// 比较左右孩子，找到较大的，将位置给largest
		largest := left
		// 如果有又孩子，且右孩子比左孩子大，则右孩子给largest
		if left+1 < heapSize && heap.comparator(arr[left+1], arr[left]) {
			largest = left + 1
		}
		// 找到index和largest中谁最大，给largest
		if heap.comparator(arr[index], arr[largest]) {
			largest = index
		}
		// 如果largest==index，说明左右孩子都不比当前位置的数大，则停止
		if largest == index {
			break
		}
		heap.swap(arr, largest, index)
		index = largest
		// 找到下一个左孩子
		left = index*2 + 1
	}
}

// add 向当前数组的i位置添加一个数，如果超过切片长度，就扩展切片长度
func (heap *MyHeap[T]) add(num T, index int) {
	if index >= len(heap.arr) {
		heap.arr = append(heap.arr, num)
	}
	heap.arr[index] = num
}

func (heap *MyHeap[T]) IsEmpty() bool {
	return heap.heapSize == 0
}

// 交换
func (heap *MyHeap[T]) swap(arr []T, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
