package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"sort"
)

// bubbleSort 冒泡排序
// 1. 0~N-1上最大的数移到最右边，比较0~1谁大谁到右边， 1~2 ... n-2~n-1的数，把大的交换到右边，最后
// 2. 0~N-2上最大的数移到最右边
// ...
// 0~1上最大的数移到最右边
// 时间复杂度 O(N^2)
func bubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 0 ~ n-1 上最大的挪到最右边
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				utils.Swap(arr, j, j+1)
			}
		}
	}
}

func main() {
	// 对数器测试500000次
	times := 500000
	for i := 0; i < times; i++ {
		// 最大长度1000，[-100,100]
		arr := utils.GenerateRandomSlice(1000, -100, 100)
		arr1 := utils.Copy(arr)
		bubbleSort(arr)
		// go自带的排序算法
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
		// 判断结果是否相等
		if !utils.IsEqual(arr, arr1) {
			fmt.Println("Fucking Fucked!!!")
			return
		}
	}
	fmt.Println("Great!!!")
}
