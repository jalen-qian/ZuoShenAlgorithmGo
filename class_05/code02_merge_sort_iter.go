package class_05

import "ZuoShenAlgorithmGo/utils"

// MergeSortByIter 迭代的方式实现mergeSort
func MergeSortByIter(arr []int) {
	// N是数组长度
	N := len(arr)
	// 数组长度<=1，直接返回
	if N <= 1 {
		return
	}
	// 初始步长是1
	step := 1
	// 步长小于数组长度，则继续
	for step < N {
		// L是每次左组的第1个位置，初始是为0位置
		L := 0
		// 一直执行直到L超过数据长度
		for L < N {
			// M为每次merge左组的结束位置
			M := L + step - 1
			// 如果左组右侧位置都 > length-1了，说明凑不齐两个组了，则不需要往下走了
			if M >= N {
				break
			}
			// R为每次merge右组的结束位置，是左组的结束位置+step
			// 如果R超过了数组边界，则设置为数组边界
			R := utils.Min(M+step, N-1)
			merge(arr, L, M, R)
			// L 来到下一次的左组（当前R+1位置）
			L = R + 1
		}
		// 防止溢出，假设数组长度特别长，接近int的最大值，此时step * 2 可能会超过int最大值，变成负数
		// 如果step > N / 2，则step * 2 必定会 > N，我们也就没必要计算一次 step*2了（因为可能溢出）
		if step > N>>1 {
			break
		}
		// 步长每次*2
		step <<= 1
	}
}
