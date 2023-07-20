package class_06

import (
	"fmt"
	"testing"
)

func TestFlagOfNetherlands(t *testing.T) {
	arr := []int{3, 6, 8, 1, 2, -5, 9, 10, 1}
	// 测试简易荷兰国旗
	FlagOfNetherlands(arr, 2)
	fmt.Println(arr)

	// 测试复杂荷兰国旗
	arr1 := []int{-2, 8, 5, 2, 6, 9, 2, 11, 7, 6, 3, 5, 6, 10, 6}
	FlagOfNetherlands1(arr1, 6)
	fmt.Println(arr1)

	// 测试复杂荷兰国旗2
	arr2 := []int{-2, 8, 5, 2, 6, 9, 2, 11, 7, 6, 3, 5, 6, 10, 2}
	FlagOfNetherlands2(arr2)
	fmt.Println(arr2)
}
