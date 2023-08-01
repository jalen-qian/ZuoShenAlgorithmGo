package class_07

// MyHeapGreater 手写堆，提供比系统堆更多的功能
// 注意：因为泛型T要放到哈希表维护索引位置，所以T要是能比较的，所以泛型类型是 comparable
type MyHeapGreater[T comparable] struct {
	arr        []T             // 存放堆的数组
	heapSize   int             // 当前堆中数字的多少（堆的大小）
	comparator MyComparator[T] // 提供一个比较器，用于确定类型如何比较大小
	indexMap   map[T]int       // T的哈希表，用于维护T在arr中的位置
}

func NewMyHeapGreater[T comparable](comparator MyComparator[T]) *MyHeapGreater[T] {
	return &MyHeapGreater[T]{
		comparator: comparator,
	}
}

// Contains 额外功能1：提供判断某个对象是否存在的功能
func (heap *MyHeapGreater[T]) Contains(obj T) bool {
	_, ok := heap.indexMap[obj]
	return ok
}

// Remove 额外功能2：从堆中移除某个对象
func (heap *MyHeapGreater[T]) Remove(obj T) {
	i, ok := heap.indexMap[obj]
	// 如果堆中没有这个元素，则移除
	if !ok {
		return
	}
	// 将这个元素和最后位置交换
	heap.swap(heap.arr, i, heap.heapSize-1)
	// 删除最后一个元素
	delete(heap.indexMap, obj)
	heap.heapSize--
	// 从i位置开始调整
	//if heap.arr[i] >=
}

// Push 向大根堆中加入一个数
func (heap *MyHeapGreater[T]) Push(num T) {
	// heapSize初始等于0，添加一个数后heapSize等于1，也就是说，heapSize就是下次要加入的数字的位置
	heap.add(num, heap.heapSize)
	// 执行heapInsert
	heap.heapInsert(heap.arr, heap.heapSize)
	// 加入之后，heapSize++
	heap.heapSize++
}

// heapInsert
func (heap *MyHeapGreater[T]) heapInsert(arr []T, index int) {
	// 如果index位置比 (index - 1) /2 排前面，则交换
	// 如果index已经是0，(0 - 1) /2 也是0，不会被交换
	for heap.comparator(arr[index], arr[(index-1)/2]) {
		heap.swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// Pop 从大根堆中弹出0位置的数（最大的数），剩下的数仍然维持大根堆
func (heap *MyHeapGreater[T]) Pop() T {
	ans := heap.arr[0]
	heap.swap(heap.arr, 0, heap.heapSize-1)
	heap.heapSize--
	heap.heapify(heap.arr, 0, heap.heapSize)
	return ans
}

// Peek 只返回堆顶的值，但是不弹出
func (heap *MyHeapGreater[T]) Peek() T {
	return heap.arr[0]
}

func (heap *MyHeapGreater[T]) Size() int {
	return heap.heapSize
}

func (heap *MyHeapGreater[T]) heapify(arr []T, index int, heapSize int) {
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
func (heap *MyHeapGreater[T]) add(num T, index int) {
	if index >= len(heap.arr) {
		heap.arr = append(heap.arr, num)
	}
	heap.arr[index] = num
}

func (heap *MyHeapGreater[T]) IsEmpty() bool {
	return heap.heapSize == 0
}

// 交换
func (heap *MyHeapGreater[T]) swap(arr []T, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
