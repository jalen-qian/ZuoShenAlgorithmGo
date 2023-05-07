package class_04

// Bitmap 位图
type Bitmap struct {
	arr []int64
	max int64
}

func NewBitmap(maxValue int64) *Bitmap {
	return &Bitmap{
		arr: make([]int64, (maxValue+64)>>6), // 数组长度：(maxValue + 64) / 64
		max: maxValue,
	}
}

// Add 添加一个数到位图中
func (b *Bitmap) Add(v int64) {
	// 不在限定范围内的，直接忽略
	if v > b.max || v < 0 {
		return
	}
	// 先找到这个数在数组的第几个数
	//index := v / 64
	// 要移动的位数，比如100，在数组的第1个数中的第36位 100 - 1 * 64 = 32
	//step := v - index*64
	// 将0左移step位后，或到数组中的对应位置的数
	//b.arr[index] |= 1 << step
	// 精简写法：
	b.arr[v>>6] |= 1 << (v - v>>6<<6)
}

// Remove 从位图中移除一个数
func (b *Bitmap) Remove(v int64) {
	// 不在限定范围内的，直接忽略
	if v > b.max || v < 0 {
		return
	}
	if !b.Exist(v) {
		return
	}
	// 先找到这个数在数组的第几个数
	//index := v / 64
	// 要移动的位数，比如100，在数组的第1个数中的第36位 100 - 1 * 64 = 32
	//step := v - index*64
	// 将0左移step位后，亦或到数组中的对应位置的数（亦或能抹掉这个位置的1）
	b.arr[v>>6] ^= 1 << (v - v>>6<<6)
}

// Exist 判断一个数是否加入到了位图中
func (b *Bitmap) Exist(v int64) bool {
	// 先找到这个数在数组的第几个数
	//index := v / 64
	// 要移动的位数，比如100，在数组的第1个数中的第36位 100 - 1 * 64 = 32
	//step := v - index*64
	// 判断这个位上是否是1，将这个数与上只有这位是1的数，判断结果是否是0
	return b.arr[v>>6]&(1<<(v-v>>6<<6)) != 0
}
