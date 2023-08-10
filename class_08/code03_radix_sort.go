package class_08

import (
	"math"

	"ZuoShenAlgorithmGo/class_03"
	"ZuoShenAlgorithmGo/utils"
)

// 基数排序

// RadixSort1 基数排序方式1，使用桶
func RadixSort1(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 从 0~N-1上排好序，并传入最大的位数（十进制）
	radixSort(arr, 0, len(arr)-1, getMaxDecimalBit(arr))
}

// radixSort 从L~R的范围上排好序 digit是最大的十进制位数，比如 897 是3位数
func radixSort(arr []int, L, R int, digit int) {
	// 准备10个桶（队列）
	var buckets [10]*class_03.MyQueue[int]
	for i := 0; i < 10; i++ {
		buckets[i] = class_03.NewMyQueue[int]()
	}
	// 最大有多少位，就需要入桶和出桶多少次
	for i := 1; i <= digit; i++ {
		// 所有数入桶
		for j := L; j <= R; j++ {
			buckets[getDigit(arr[j], i)].Push(arr[j])
		}
		// 重新出桶
		index := L
		for _, bucket := range buckets {
			for !bucket.IsEmpty() {
				arr[index] = bucket.Poll()
				index++
			}
		}
	}
}

// getMaxDecimalBit 获取数组中最大的十进制位数 比如：100 对应3位，返回3
func getMaxDecimalBit(arr []int) int {
	// 先获取最大的数
	maxNum := math.MinInt
	for _, num := range arr {
		maxNum = utils.Max(maxNum, num)
	}
	// 计算最大
	decimal := 0
	for maxNum%10 != 0 {
		decimal++
		maxNum = maxNum / 10
	}
	return decimal
}

// 获取num在第d位上的数字，比如 357 d=2 则返回5
func getDigit(num int, d int) int {
	return int(float64(num)/math.Pow(10, float64(d-1))) % 10
}

// === 上面的基数排序是基于容器实现的，也可以不基于容器实现，实现代码如下

// RadixSort2 基数排序方式2，优化版，不真正使用容器
func RadixSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 从 0~N-1上排好序，并传入最大的位数（十进制）
	radixSor2(arr, 0, len(arr)-1, getMaxDecimalBit(arr))
}

// 从L~R的范围上排好序，并且不准备10个桶，只准备一个辅助数组
func radixSor2(arr []int, L, R int, digit int) {
	// 准备10长度的int数组，统计每次入桶出桶过程中数字0-9出现的次数
	counter := [10]int{}
	// 有多少个数，就需要准备多少个辅助空间
	help := make([]int, R-L+1)
	// 最大有多少位，就需要入桶和出桶多少次
	for i := 1; i <= digit; i++ {
		// 统计当前位的数，出现的次数
		for j := L; j <= R; j++ {
			counter[getDigit(arr[j], i)]++
		}
		// 将counter数组转换为前缀数组，此时counter数组的含义就变了
		// 之前比如 counter[5] == 100 表示数字5出现了100次
		// 现在假如 counter[5] == 200，由于counter变成了前缀数组，表示 <= 5的数，出现了200次
		for j := 1; j < 10; j++ {
			counter[j] = counter[j] + counter[j-1]
		}
		// 从原始数组的右往左遍历，比如遍历到375（个位数是5，如果是桶，则会放入5号桶的最后一个）
		// 然后假如获取到counter[5]==10，说明当前个位<=5的有10个数，则直接将这个数放到help数组的第9位（0~9位的最后一个位置）
		// 解释：用桶的方式，这个数也会被倒回到原数组的第9位，因为如果 counter[5]==10，说明即使用桶的方式，0-5号桶中的数的个数
		// 总共也必定是10个，而当前是倒序遍历的，这个 375 一定在5号桶的最后一个，则一定是0-5号桶中最后一个被倒出来的，则一定会被倒
		// 到位置9。这里使用一个辅助数组实现了和用10个桶完全一样的效果。
		for j := R; j >= L; j-- {
			help[counter[getDigit(arr[j], i)]-1] = arr[j]
		}
		// 将help倒回arr
		for j := L; j <= R; j++ {
			arr[j] = help[j]
		}
	}
}
