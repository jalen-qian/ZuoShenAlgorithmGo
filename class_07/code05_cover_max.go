package class_07

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
)

/**
最大线段重合问题：
给定很多个线段，每个线段都有两个数[start, end]，表示线段的开始位置和结束位置，左右都是闭区间，规定：
1. 线段的开始时间和结束时间一定都是整数值。
2. 线段的重合区域长度必须 >= 1。

请返回线段的最多重合区域中，包含了几条线段。
*/

// MaxCover 最大线段重合问题 m是线段的切片，[2]int代表一个线段
func MaxCover(m [][2]int) int {
	// 先按照开始位置从小到大排好序
	sort.SliceStable(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})
	// 准备一个小根堆
	minHeap := NewMyHeap[int](func(a int, b int) bool {
		return a < b
	})
	// 准备一个变量，保存最大值
	max := 0
	// 遍历这些线段
	for _, line := range m {
		// 先将堆中所有 <= 线段左边界的值弹出
		for !minHeap.IsEmpty() && minHeap.Peek() <= line[0] {
			minHeap.Pop()
		}
		// 将线段的右边界Push到堆中
		minHeap.Push(line[1])
		// 此时堆的大小，就是以line[0]开头的重合区域的线段数，求max与它之间的最大值，并赋值给max
		max = utils.Max(max, minHeap.Size())
	}
	return max
}

// MaxCover1 使用笨办法实现
func MaxCover1(m [][2]int) int {
	// 先找到全局最小左边界和最大右边界
	left, right := 0, 0
	for _, line := range m {
		left = utils.Min(left, line[0])
		right = utils.Max(right, line[1])
	}
	max := 0
	// 从min遍历到max，每次都取 0.5 的中间值
	for p := float64(left) + 0.5; p < float64(right); p++ {
		// 遍历所有的线段，判断是否 cover 了当前值，cover的就+1
		cover := 0
		for _, line := range m {
			if float64(line[0]) < p && p < float64(line[1]) {
				cover++
			}
		}
		max = utils.Max(cover, max)
	}
	return max
}
