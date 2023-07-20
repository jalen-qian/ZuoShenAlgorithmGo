package class_05

/**
区间和的个数问题：
给定一个整数数组 nums 和两个整数 lower 和 upper。求数组中，值位于范围[lower,upper]（包含lower和upper）之内的区间和的个数。

区间和的概念：s(i,j) 表示在 nums 中，位置从i到j的元素之和，包含 i和j,(i <= j)

leetcode 原题：https://leetcode.cn/problems/count-of-range-sum/
*/

// countRangeSum 区间和个数问题
func countRangeSum(nums []int, lower int, upper int) int {
	// 如果是空数组，则区间都没有，所以区间和个数一定是0
	if len(nums) == 0 {
		return 0
	}
	// 生成一个前缀和数组
	sum := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			sum[i] = num
		} else {
			sum[i] = sum[i-1] + num
		}
	}
	// 将前缀和数组用来计算
	return count(sum, 0, len(nums)-1, lower, upper)
}

// count 干这样一件事：传入一个前缀和数组，L和R区间，和取值范围[lower与upper]
// 返回[L~R]范围内的区间和满足处于[lower,upper]的个数。
func count(sum []int, L int, R int, lower, upper int) int {
	// 1. 如果 L == R，说明求的是sum[R~R]范围
	// sum[R] 表示 nums 在[0~R]范围的区间和，如果这个区间和满足要求，则说明[0~R]这个区间是达标的，返回1
	// 不达标则返回0
	if L == R {
		if sum[R] >= lower && sum[R] <= upper {
			return 1
		}
		return 0
	}
	// 2. 如果 L < R 则至少可以分为左右两个部分，递归
	M := L + (R-L)>>1
	// 2.1 先求左边部分的个数，同时将左边排好序
	leftCount := count(sum, L, M, lower, upper)
	// 2.2 再求右边部分的个数，同时将右边排好序
	rightCount := count(sum, M+1, R, lower, upper)
	// 2.3 再将左右两个数组merge，并求merge过程中的个数
	mergeCount := mergeForCountOfRangeSum(sum, L, M, R, lower, upper)
	return leftCount + rightCount + mergeCount
}

// mergeForCountOfRangeSum 区间和个数问题的merge过程
func mergeForCountOfRangeSum(sum []int, L int, M int, R int, lower int, upper int) int {
	// 1. 先求区间和满足要求的个数
	windowL, windowR := L, L
	index := M + 1
	var total int
	for index <= R {
		// 计算出当前右组数的限定范围
		min := sum[index] - upper
		max := sum[index] - lower
		// 将windowL移动到一个>= min的位置
		for windowL <= M && sum[windowL] < min {
			windowL++
		}
		// 将windowR移动到第一个 > max的位置
		for windowR <= M && sum[windowR] <= max {
			windowR++
		}
		total += windowR - windowL
		index++
	}
	// 2. 进行常规的merge过程
	help := make([]int, R-L+1)
	indexL, indexR := L, M+1
	helpI := 0
	for indexL <= M && indexR <= R {
		if sum[indexL] <= sum[indexR] {
			help[helpI] = sum[indexL]
			indexL++
		} else {
			help[helpI] = sum[indexR]
			indexR++
		}
		helpI++
	}
	// 不管哪个数组没倒完，都拷贝到help中
	for indexL <= M {
		help[helpI] = sum[indexL]
		indexL++
		helpI++
	}
	for indexR <= R {
		help[helpI] = sum[indexR]
		indexR++
		helpI++
	}
	// 将help数组拷贝回sum
	// 最后将help中的数倒回arr
	helpI = 0
	for i := L; i <= R; i++ {
		sum[i] = help[helpI]
		helpI++
	}
	return total
}
