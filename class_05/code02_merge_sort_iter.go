package class_05

import "ZuoShenAlgorithmGo/utils"

// MergeSortByIter 迭代的方式实现mergeSort
func MergeSortByIter(arr []int) {
	length := len(arr)
	// 数组长度<=1，直接返回
	if length <= 1 {
		return
	}
	// 初始步长是1
	step := 1
	// 步长小于数组长度，则继续
	for step < length {
		// L是每次左组的第1个位置，初始是为0位置
		L := 0
		// 一直执行直到L超过数据长度
		for L < length {
			// M为每次merge左组的结束位置
			M := L + step - 1
			// 如果左组右侧位置都 > length-1了，说明凑不齐两个组了，则不需要往下走了
			if M >= length {
				break
			}
			// R为每次merge右组的结束位置，是左组的结束位置+step
			// 如果R超过了数组边界，则设置为数组边界
			R := utils.Min(M+step, length-1)
			merge(arr, L, M, R)
			// L 往下跳两个step的距离（每次merge了两个step长度的区间）
			L += step << 1
		}
		// 步长每次*2
		step <<= 1
	}
}
