package main

/*
给定一个无序的数组，使用递归返回数组中的最大值，要求时间复杂度O(logN)

思路：使用二分的思想递归，假设我设计了一个process函数，给定数组和一个区间，process会返回这个区间内的最大值

我们假设给定的array的区间分别是 L 和 R (假设数组长度是n，则L>=0，R<=n-1, L <= R)，我找到L和R的中点 mid
则必然满足  maxValue = max(process(arr, L, mid), process(arr, mid+1, R))
也就是将[L,mid]区间和[mid+1,R]区间分别用process求最大值，最终的最大值肯定是这两个值中较大的那个

*/

// GetMax 已知我们实现了process函数，则最终的结果是[0, n-1]上的最大值
func GetMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

// process 就是要实现的函数，会返回arr在L~R上的最大值
func process(arr []int, L, R int) int {
	// 如果左右区间都相等了，只有一个数了，则返回这个数
	if L == R {
		return arr[L]
	}
	if L > R {
		return 0
	}
	// 否则，找到中点
	mid := L + (R-L)>>1
	// 取这两个区间的最大值，返回
	return max(process(arr, L, mid), process(arr, mid+1, R))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// GetMaxForTest 对数器，直接遍历找最大值
func GetMaxForTest(arr []int) int {
	var maxValue int
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
