package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Swap 交换一个切片中的两个位置的数
func Swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// GenerateRandomSlice 初始化一个随机切片
//
//		@param maxLen 最大长度
//	 @param minNum 数组成员最小值
//	 @param maxNum 数组成员最大值
func GenerateRandomSlice(maxLen int, minNum int, maxNum int, customizedValues ...int) []int {
	// 如果传入了自定义的值，则使用自定义的值生成数组直接返回
	if len(customizedValues) > 0 {
		customizedAns := make([]int, 0)
		for _, value := range customizedValues {
			customizedAns = append(customizedAns, value)
		}
		return customizedAns
	}
	if maxLen < 0 || minNum > maxNum {
		panic("最大长度不可小于0，且最小值不可大于最大值！")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// [0,maxLen]
	length := r.Intn(maxLen + 1)
	result := make([]int, length)
	// 填充数
	for i := 0; i < length; i++ {
		// 49 100
		// [0, 51] + 49
		// [49, 100]
		num := r.Intn(maxNum-minNum+1) + minNum
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
//
//		@param maxLen 最大长度
//	 @param minNum 数组成员最小值
//	 @param maxStep 数组成员递增幅度
func GenerateRandomSortedSlice(maxLen int, minNum int, maxStep int) []int {
	if maxLen < 0 || maxStep < 0 {
		panic("最大长度不可小于0，且最大递增幅度不可小于0！")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// [0, maxLen]
	length := r.Intn(maxLen + 1)
	result := make([]int, length)
	if length == 0 {
		return result
	}
	// 起始数，保证 >= minNum
	result[0] = minNum + r.Intn(10000)
	for i := 1; i < length; i++ {
		step := r.Intn(maxStep + 1)
		result[i] = result[i-1] + step
	}
	return result
}

// GenerateRandomSliceWithoutEqualNeighbor 初始化一个随机切片，并保证相邻的数不相等
//
//		@param maxLen 最大长度
//	 @param minNum 数组成员最小值
//	 @param maxNum 数组成员最大值
func GenerateRandomSliceWithoutEqualNeighbor(maxLen int, minNum int, maxNum int) []int {
	if maxLen < 0 || minNum > maxNum {
		panic("最大长度不可小于0，且最小值不可大于最大值！")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r.Intn(maxLen + 1)
	result := make([]int, length)
	// 填充数
	for i := 0; i < length; i++ {
		// 假如 minNum和maxNum分别是 -49 100
		// [0, 149] - 49
		// [-49, 100]
		num := r.Intn(maxNum-minNum+1) + minNum
		if i > 0 && num == result[i-1] {
			// 如何和前一个数相等，则随机增加1~100的任意一个值，保证不相等
			tmp := r.Intn(101) + 1
			num += tmp
		}
		result[i] = num
	}
	return result
}

// GenerateRandInt32 生成随机的int32整数，取值范围 [-maxNum, maxNum]
func GenerateRandInt32(maxNum int32) int32 {
	return rand.Int31n(maxNum+1) - rand.Int31n(maxNum+1)
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SprintList(arr []int) string {
	builder := strings.Builder{}
	builder.WriteString("[]int{")
	for i, num := range arr {
		if i != len(arr)-1 {
			builder.WriteString(fmt.Sprintf("%d, ", num))
		} else {
			builder.WriteString(fmt.Sprintf("%d", num))
		}
	}
	builder.WriteString("}")
	return builder.String()
}
