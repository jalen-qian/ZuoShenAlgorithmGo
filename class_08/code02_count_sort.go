package class_08

import (
	"math"

	"ZuoShenAlgorithmGo/utils"
)

// CountSort 计数排序，限定取值范围是0-200
func CountSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	max := math.MinInt
	// 遍历一遍找最大值
	for _, num := range arr {
		max = utils.Max(num, max)
	}
	// 准备一个0-max的计数数组
	count := make([]int, max+1)
	// 遍历一遍统计每个数的数量
	for _, num := range arr {
		count[num]++
	}
	// 遍历count数组并回填到原始数组
	i := 0 // i 记录当前回填的值在原始数组中的位置
	for j, c := range count {
		for c != 0 {
			arr[i] = j
			c--
			i++
		}
	}
}
