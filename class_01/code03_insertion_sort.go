package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"sort"
)

// insertionSort 插入排序
// 1.  保证 0~1 有序，大的排到右边
// 2.  保证 0~2 有序，2往1看，1往0看只要左边的大，就和左边的交换
// 3.  保证 0~3 有序，3往2看，2往1看，1往0看
// ...
// n-1 保证 0~n-1有序
// 时间复杂度 O(N^2)
func insertionSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 保证 0-i 有序，i从 1 ~ n-1递增
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				utils.Swap(arr, j, j-1)
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
		insertionSort(arr)
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
