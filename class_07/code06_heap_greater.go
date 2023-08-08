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
		indexMap:   make(map[T]int),
	}
}

// Contains 额外功能1：提供判断某个对象是否存在的功能
func (heap *MyHeapGreater[T]) Contains(obj T) bool {
	_, ok := heap.indexMap[obj]
	return ok
}

// Remove 额外功能2：从堆中移除某个对象
func (heap *MyHeapGreater[T]) Remove(obj T) {
	// 获取obj对象在堆中的位置
	index, ok := heap.indexMap[obj]
	// 如果堆中没有这个元素，则什么都不用做，返回
	if !ok {
		return
	}
	// 要填补index空缺位置的元素，最后一个元素（实际上就是和最后一个元素交换的另一种写法）
	replace := heap.arr[heap.heapSize-1]
	// 将最后一个元素从堆中删掉
	heap.heapSize--
	// 将obj的索引删掉
	delete(heap.indexMap, obj)

	// 如果要删的obj，正好就是最后一个元素 replace ，则什么都不用干
	if replace != obj {
		// 要删的obj，不是最后一个元素
		// 将replace塞到index的位置
		heap.arr[index] = replace
		heap.indexMap[replace] = index
		// 给replace元素重新Resign调整
		heap.Resign(replace)
	}
}

// Resign 额外功能3，改变了某个对象的值，并且这个值是影响比较大小的，改完就不符合堆结构了，让堆重新调整成堆结构
// 这个功能是系统堆肯定不会提供的
func (heap *MyHeapGreater[T]) Resign(obj T) {
	index, ok := heap.indexMap[obj]
	if ok {
		// 这两个只会发生一个
		heap.heapInsert(index)
		heap.heapify(index)
	}
}

// Push 向堆中加入一个对象
func (heap *MyHeapGreater[T]) Push(obj T) {
	// heapSize初始等于0，添加一个数后heapSize等于1，也就是说，heapSize就是下次要加入的数字的位置
	heap.add(obj, heap.heapSize)
	heap.indexMap[obj] = heap.heapSize
	// 执行heapInsert，从最后一个位置往上调整
	heap.heapInsert(heap.heapSize)
	// 加入之后，heapSize++
	heap.heapSize++
}

// heapInsert
func (heap *MyHeapGreater[T]) heapInsert(index int) {
	// 如果index位置比 (index - 1) /2 排前面，则交换
	for heap.comparator(heap.arr[index], heap.arr[(index-1)/2]) && index != 0 {
		heap.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// Pop 从大根堆中弹出0位置的数（最大的数），剩下的数仍然维持大根堆
func (heap *MyHeapGreater[T]) Pop() T {
	// 弹出堆顶的元素
	ans := heap.arr[0]
	heap.swap(0, heap.heapSize-1)
	heap.heapSize--
	// 堆最后位置的元素换到了0位置，从0位置开始heapify，重新调整成堆
	heap.heapify(0)
	// 删除ans在map中的位置索引
	delete(heap.indexMap, ans)
	return ans
}

// Peek 只返回堆顶的值，但是不弹出
func (heap *MyHeapGreater[T]) Peek() T {
	return heap.arr[0]
}

func (heap *MyHeapGreater[T]) Size() int {
	return heap.heapSize
}

func (heap *MyHeapGreater[T]) heapify(index int) {
	// 左孩子 index * 2 + 1 如果有右孩子，则右孩子是 left + 1
	left := index*2 + 1
	// 当前堆最后一个数的位置是 heapSize - 1
	for left < heap.heapSize {
		// 比较左右孩子，找到较大的，将位置给largest
		largest := left
		// 如果有又孩子，且右孩子比左孩子大，则右孩子给largest
		if left+1 < heap.heapSize && heap.comparator(heap.arr[left+1], heap.arr[left]) {
			largest = left + 1
		}
		// 找到index和largest中谁最大，给largest
		if heap.comparator(heap.arr[index], heap.arr[largest]) {
			largest = index
		}
		// 如果largest==index，说明左右孩子都不比当前位置的数大，则停止
		if largest == index {
			break
		}
		heap.swap(largest, index)
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
func (heap *MyHeapGreater[T]) swap(a, b int) {
	// 交换过程中，两个对象的位置索引也交换
	heap.indexMap[heap.arr[a]] = b
	heap.indexMap[heap.arr[b]] = a
	tmp := heap.arr[a]
	heap.arr[a] = heap.arr[b]
	heap.arr[b] = tmp
}

// GetAllElements 返回所有的元素
func (heap *MyHeapGreater[T]) GetAllElements() []T {
	var ans []T
	for i := 0; i < heap.heapSize; i++ {
		ans = append(ans, heap.arr[i])
	}
	return ans
}
