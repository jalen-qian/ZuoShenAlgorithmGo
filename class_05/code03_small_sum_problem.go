package class_05

/**
小和问题：
给定一个数组，每个数中，所有左边比它小的数加起来的结果，叫做这个数的“小和”，求一个数组中所有小和相加的结果，就是小和问题。
使用递归版本的MergeSort来解决这个问题。
*/

// SmallSum 使用MergeSort求小和
func SmallSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	// 求 0 ~ N-1 范围内的小和
	return processForSmallSum(arr, 0, len(arr)-1)
}

// process1 小和问题的process，除了在L到R排好序以外，还返回局部的小和累加值
func processForSmallSum(arr []int, L, R int) int {
	if R-L == 0 {
		return 0
	}
	// 找到中点
	M := L + (R-L)>>1
	// 先将左半部分变为有序，并求出左半部分的小和
	left := processForSmallSum(arr, L, M)
	// 再将右半部分变为有序，并求出右半部分的小和
	right := processForSmallSum(arr, M+1, R)
	// 合并，将整体变为有序，并求出整体的小和
	return left + right + mergeForSmallSum(arr, L, M, R)
}

func mergeForSmallSum(arr []int, L, M, R int) int {
	// 两个索引分别指向两个子数组的第0号位置
	index1 := L
	index2 := M + 1
	// ans记录当前merge过程累加的小和
	var ans int
	// 分配一个L-R长度的数组
	help := make([]int, R-L+1)
	var helpI int // help数组的索引，标示当前拷贝的位置
	// 当两个索引都不越界时，谁小拷贝谁
	for index1 <= M && index2 <= R {
		// 注意：只有左边小时，才拷贝左边的值，并累加，左右相等时，拷贝右边的值
		if arr[index1] < arr[index2] {
			help[helpI] = arr[index1]
			ans += arr[index1] * (R - index2 + 1)
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
	return ans
}
