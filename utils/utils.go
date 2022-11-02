package utils

import (
	"math/rand"
	"time"
)

// Swap 交换一个切片中的两个位置的数
func Swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// GenerateRandomSlice 初始化一个随机切片
// 	@param maxLen 最大长度
//  @param minNum 数组成员最小值
//  @param maxNum 数组成员最大值
func GenerateRandomSlice(maxLen int, minNum int, maxNum int) []int {
	if maxLen < 0 || minNum > maxNum {
		panic("最大长度不可小于0，且最小值不可大于最大值！")
	}
	rand.Seed(time.Now().UnixNano())
	// [0,maxLen]
	length := rand.Intn(maxLen + 1)
	result := make([]int, length)
	// 填充数
	for i := 0; i < length; i++ {
		// 49 100
		// [0, 51] + 49
		// [49, 100]
		num := rand.Intn(maxNum-minNum+1) + minNum
		result[i] = num
	}
	return result
}

// Copy 拷贝一个相同的切片
func Copy(arr []int) []int {
	if arr == nil {
		return nil
	}
	result := make([]int, len(arr))
	if len(arr) == 0 {
		return result
	}
	for i, num := range arr {
		result[i] = num
	}
	return result
}

// IsEqual 比较一个切片是否全等
func IsEqual(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 != nil {
		return false
	}
	if arr2 == nil && arr1 != nil {
		return false
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr2[i] != arr1[i] {
			return false
		}
	}
	return true
}

// GenerateRandomSortedSlice 初始化一个随机有序切片，从小到大排序
// 	@param maxLen 最大长度
//  @param minNum 数组成员最小值
//  @param maxStep 数组成员递增幅度
func GenerateRandomSortedSlice(maxLen int, minNum int, maxStep int) []int {
	if maxLen < 0 || maxStep < 0 {
		panic("最大长度不可小于0，且最大递增幅度不可小于0！")
	}
	rand.Seed(time.Now().UnixNano())
	// [0, maxLen]
	length := rand.Intn(maxLen + 1)
	result := make([]int, length)
	if length == 0 {
		return result
	}
	// 起始数，保证 >= minNum
	result[0] = minNum + rand.Int()
	for i := 1; i < length; i++ {
		step := rand.Intn(maxStep + 1)
		result[i] = result[i-1] + step
	}
	return result
}
