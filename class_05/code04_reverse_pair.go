package class_05

/**
逆序对问题：在一个数组中， 任何一个前面的数`a`，和任何一个后面的数`b`， 如果`(a,b)`是降序的，就称为逆序对。返回数组中所有的逆序对的个数。
*/

// ReversePair 逆序对问题
func ReversePair(arr []int) int {
	return processForReversePair(arr, 0, len(arr)-1)
}

// processForReversePair 逆序对问题的process递归过程，归并排序的同时返回逆序对个数
func processForReversePair(arr []int, L, R int) int {
	if L >= R {
		return 0
	}
	// 求中间位置
	M := L + (R-L)>>1
	// 总共的逆序对个数 = 左组递归逆序对个数 + 右组递归逆序对个数 + 左右组merge一次产生的逆序对个数
	left := processForReversePair(arr, L, M)
	right := processForReversePair(arr, M+1, R)
	merge := mergeForReversePair(arr, L, M, R)
	return left + right + merge
}

// mergeForReversePair 逆序对问题的merge过程
// 0 3 7
func mergeForReversePair(arr []int, L, M, R int) int {
	if L == R {
		return 0
	}
	count := 0 // 逆序对统计个数
	// 产生一个help数组，长度是当前区间的长度
	help := make([]int, R-L+1)
	// 左右两个数组的指针先指向最右边的位置，从右往左merge
	indexH := R - L // help数组也是从最右边开始填充
	indexL := M
	indexR := R
	// 当两个数组的下标都不越界时，执行merge的过程
	for indexL >= L && indexR > M {
		// 当左组数比右组数大时，拷贝左组，并产生逆序对
		if arr[indexL] > arr[indexR] {
			help[indexH] = arr[indexL]
			count += indexR - M // indexR - M 就是右组剩余数的个数，有多少个就有多少个逆序对
			indexL--
		} else {
			// 当左组数<=右组数时，都拷贝右组，并不产生逆序对
			help[indexH] = arr[indexR]
			indexR--
		}
		// 每拷贝一个数，help数组下标往左移动
		indexH--
	}
	// 不管谁越界了，都将剩余的数组拷贝到help中
	for indexL >= L {
		help[indexH] = arr[indexL]
		indexL--
		indexH--
	}
	for indexR > M {
		help[indexH] = arr[indexR]
		indexR--
		indexH--
	}
	// 最后将help数组拷贝回去
	indexL = L
	for _, num := range help {
		arr[indexL] = num
		indexL++
	}
	return count
}

// ReversePairForTest 逆序对问题对数器，使用暴力方法 O(N^2)
func ReversePairForTest(arr []int) int {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		// j 从 i的下一个开始找
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				count++
			}
		}
	}
	return count
}
