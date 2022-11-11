package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
)

/*
局部最小值问题：
给定一个数组，这个数组相邻的数不相等，找出这个数组中的任意一个局部最小值，并返回这个局部最小值的位置，如果没有，则返回 -1

局部最小值定义：
1. 如果位置0 < 位置1，则位置0为局部最小值
2. 如果数组长度为1，则0位置算局部最小值
3. 如果最后位置 n-1 比 倒数第二位置 n-2 小，则 n-1 位置算局部最小值
4. 如果位置i处于数组中间某个位置，且 num[i-1] > num[i] < num[i+1]，则位置i为局部最小值
这个题目虽然不是有序数组，但是依然可以使用二分法求解
分析：首先，如果我找到一个区间 a ~ b [...x x x a x x ... b x x...]，且
arr[a]比右边的数大，arr[b]比左边的数大，也就是从a~b是先变小，再变大，则这种情况下，a~b中间一定存在局部最小值。
因为如果不存在，要么单调递增，要么单调递减，是不可能先变小后变大的。
假设我取 a ~ b的中点m [...x x x a x x...m...x x b x x...]
则只可能存在3种情况：
    1). m正好是局部最小值，则不用找了，返回m就行了
    2). m比左边大，则 a ~ m 就出现了 a ~ b一样的情形，可以继续在 a ~ m 二分
    3). m比右边大，则 m ~ b 也出现了 a ~ b一样的情形，可以继续在 m ~ b 二分
这样，利用二分法，也可以找出一个局部最小值，时间复杂度为O(logN)
总结：只要能找到一种规则，在砍掉一半后，剩余的样本中，这个规则仍然成立，则可以用二分法
*/
func findLessIndex(arr []int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return 0
	}
	// 长度至少是2，先判断第一个数是否是局部最小值
	if arr[0] < arr[1] {
		return 0
	}
	// 再判断最后一个数是否是局部最小值
	if arr[length-1] < arr[length-2] {
		return length - 1
	}
	// 走到这，说明至少有3个数，则在 [1~length-2]上二分
	L, R := 1, length-2
	for L <= R {
		mid := L + ((R - L) >> 1)
		// 如果中点正好是局部最小值，则直接返回
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[mid-1] {
			R = mid - 1
		} else if arr[mid] > arr[mid+1] {
			L = mid + 1
		}
	}
	return L
}

// checkLessIndex 检查局部最小值是否正确
func checkLessIndex(arr []int, lessIndex int) bool {
	if len(arr) == 0 {
		return lessIndex == -1
	} else if len(arr) == 1 {
		return lessIndex == 0
	} else {
		// 数组长度至少为2，此时必然会有局部最小值，lessIndex必须在arr的范围
		if lessIndex < 0 || lessIndex >= len(arr) {
			return false
		}
		// 如果是0位置，则判断是否比1位置小
		if lessIndex == 0 {
			return arr[0] < arr[1]
		}
		// 如果是最后位置，则判断是否比倒数第二位置小
		if lessIndex == len(arr)-1 {
			return arr[len(arr)-1] < arr[len(arr)-2]
		}
		// 如果是中间位置，则判断是否比相邻的两边都小
		return arr[lessIndex] < arr[lessIndex-1] && arr[lessIndex] < arr[lessIndex+1]
	}
}

func main() {
	testTimes := 500000
	for i := 0; i < testTimes; i++ {
		arr := utils.GenerateRandomSliceWithoutEqualNeighbor(1000, 10, 1000)
		lessIndex := findLessIndex(arr)
		if !checkLessIndex(arr, lessIndex) {
			fmt.Println("Fucking fucked!!!")
			return
		}
	}
	fmt.Println("Great!!!")
}
