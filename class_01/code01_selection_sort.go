package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"sort"
)

/*
选择排序
从 0 ~ n-1上，找到最小的数，与位置0的数交换
从 1 ~ n-1上，找到最小的数，与位置1的数交换
...
从 n-2 ~ n-1上，找最小的数，与位置n-2的数交换
时间复杂度：O(N^2)
*/
func selectionSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 从i到n-1位置上，找到最小的数，与i位置交换
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				utils.Swap(arr, i, j)
			}
		}
	}
	return
}

func main() {
	// 对数器测试500000次
	times := 500000
	for i := 0; i < times; i++ {
		// 最大长度1000，[-100,100]
		arr := utils.GenerateRandomSlice(1000, -100, 100)
		arr1 := utils.Copy(arr)
		selectionSort(arr)
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
