package class_05

/**
归并排序，使用的是递归的思想，是第一个时间复杂度为 O(N * logN)的排序算法。
*/

// MergeSort 归并排序，给定一个数组
func MergeSort(arr []int) {
	// 数组长度<=1，直接返回
	if len(arr) <= 1 {
		return
	}
	// 将数组的 0 ~ n-1位置排好序，则整体排好序了
	process(arr, 0, len(arr)-1)
}

// process 会将arr的 L - R 区间排好序
func process(arr []int, L int, R int) {
	if R-L == 0 {
		return
	}
	// 找到中点
	M := L + (R-L)>>1
	// 先将左半部分变为有序
	process(arr, L, M)
	// 再将右半部分变为有序
	process(arr, M+1, R)
	// 合并，将整体变为有序
	merge(arr, L, M, R)
}

// merge 将arr的两个已经排好序的部分变成整体有序
func merge(arr []int, L, M, R int) {
	// 两个索引分别指向两个子数组的第0号位置
	index1 := L
	index2 := M + 1
	// 分配一个L-R长度的数组
	help := make([]int, R-L+1)
	var helpI int // help数组的索引，标示当前拷贝的位置
	// 当两个索引都不越界时，谁小拷贝谁
	for index1 <= M && index2 <= R {
		if arr[index1] <= arr[index2] {
			help[helpI] = arr[index1]
			index1++
		} else {
			help[helpI] = arr[index2]
			index2++
		}
		helpI++
	}
	// 至少有一个数组越界了，不管谁没越界，都将剩下的数拷贝（下面两个for只会执行1个）
	for index1 <= M {
		help[helpI] = arr[index1]
		helpI++
		index1++
	}
	// 至少有一个数组越界了，不管谁没越界，都将剩下的数拷贝
	for index2 <= R {
		help[helpI] = arr[index2]
		helpI++
		index2++
	}
	// 最后将help中的数倒回arr
	helpI = 0
	for i := L; i <= R; i++ {
		arr[i] = help[helpI]
		helpI++
	}
}
