package class_06

import (
	"ZuoShenAlgorithmGo/utils"
	"math/rand"
)

/**
快速排序:
前面掌握了复杂版本荷兰国旗问题的过程。快速排序就是利用这个流程和递归实现的排序。
*/

// QuickSort1 快排1.0版本
func QuickSort1(nums []int) {
	if len(nums) < 2 {
		return
	}
	process1(nums, 0, len(nums)-1)
}

// process1 让nums在L到R上排好序
func process1(nums []int, L, R int) {
	// 打到的x位置如果正好是上一轮L的位置（比如[3,2]，会变成[2,3]，2换到了L位置）
	// 或者L+1的位置（比如[3,5,4]，会变成[3,4,5]，4来到了L+1位置）
	// 这次递归的R就等于L或者小于L
	// 说明要处理的区域没有数或者只有一个数，直接返回
	if L >= R {
		return
	}
	// 至少有两个数，先执行一个荷兰国旗
	i := NetherlandsFlag3(nums, L, R)
	// 重复在小于区域递归
	process1(nums, L, i-1)
	// 重复在大于区域递归
	process1(nums, i+1, R)
}

// NetherlandsFlag3 从L到R位置执行荷兰国旗，返回R位置的数最终的位置
func NetherlandsFlag3(nums []int, L, R int) int {
	small := L - 1 // 小于区域初始位置，不将L扩进去
	large := R     // 大于区域初始位置，R，先将R包含进去
	i := L         // i从L位置开始
	// 当i没有与大于区域撞上
	for i < large {
		// 如果 nums[i] < nums[R]，则将i位置与左侧区域下一个位置交换，并将小于区域往右扩，i跳下一位
		if nums[i] < nums[R] {
			utils.Swap(nums, small+1, i)
			small++
			i++
		} else if nums[i] == nums[R] {
			// 如果 nums[i] == nums[R] ，则i直接跳下一位
			i++
		} else {
			// 如果 nums[i] > nums[R]，则与大于区域左侧位置交换，大于区域往左扩，i留在原地
			utils.Swap(nums, large-1, i)
			large--
		}
	}
	// 最后将R位置，与large位置交换
	utils.Swap(nums, R, large)
	// 最终x与谁交换，就说明落到了哪个位置
	return large
}

// ================快排2.0版本=============

// QuickSort2 快排2.0版本，每次递归定位一个x的区间
func QuickSort2(nums []int) {
	if len(nums) < 2 {
		return
	}
	process2(nums, 0, len(nums)-1)
}

// process2 让nums在L到R上排好序
func process2(nums []int, L, R int) {
	// 打到的x位置如果正好是上一轮L的位置（比如[3,2]，会变成[2,3]，2换到了L位置）
	// 或者L+1的位置（比如[3,5,4]，会变成[3,4,5]，4来到了L+1位置）
	// 这次递归的R就等于L或者小于L
	// 说明要处理的区域没有数或者只有一个数，直接返回
	if L >= R {
		return
	}
	// 至少有两个数，先执行一个荷兰国旗
	left, right := NetherlandsFlag4(nums, L, R)
	// 重复在小于区域递归
	process2(nums, L, left-1)
	// 重复在大于区域递归
	process2(nums, right+1, R)
}

// NetherlandsFlag4 从L到R位置执行荷兰国旗，返回R位置的数的区间坐标
func NetherlandsFlag4(nums []int, L, R int) (int, int) {
	small := L - 1 // 小于区域初始位置，不将L扩进去
	large := R     // 大于区域初始位置，R，先将R包含进去
	i := L         // i从L位置开始
	// 当i没有与大于区域撞上
	for i < large {
		// 如果 nums[i] < nums[R]，则将i位置与左侧区域下一个位置交换，并将小于区域往右扩，i跳下一位
		if nums[i] < nums[R] {
			utils.Swap(nums, small+1, i)
			small++
			i++
		} else if nums[i] == nums[R] {
			// 如果 nums[i] == nums[R] ，则i直接跳下一位
			i++
		} else {
			// 如果 nums[i] > nums[R]，则与大于区域左侧位置交换，大于区域往左扩，i留在原地
			utils.Swap(nums, large-1, i)
			large--
		}
	}
	// 最后将R位置，与large位置交换
	utils.Swap(nums, R, large)
	// 最终相同的区域，是在small+1，large的区间
	return small + 1, large
}

// ============== 快排3.0 随机快排 ================

// QuickSort3 快排3.0版本，每次递归随机定位一个x
func QuickSort3(nums []int) {
	if len(nums) < 2 {
		return
	}
	process3(nums, 0, len(nums)-1)
}

// process3 让nums在L到R上排好序
func process3(nums []int, L, R int) {
	// L>=R说明要处理的区域没有数或者只有一个数，直接返回
	if L >= R {
		return
	}
	// 至少有两个数，先执行一个荷兰国旗
	left, right := NetherlandsFlag5(nums, L, R)
	// 重复在小于区域递归
	process3(nums, L, left-1)
	// 重复在大于区域递归
	process3(nums, right+1, R)
}

// NetherlandsFlag5 从L到R位置执行荷兰国旗，返回R位置的数的区间坐标
func NetherlandsFlag5(nums []int, L, R int) (int, int) {
	// 每次随机定位一个位置，L到R上的位置，并与R位置交换
	utils.Swap(nums, R, rand.Intn(R-L+1)+L) // [0,R-L] + L =>   [L,R]
	small := L - 1                          // 小于区域初始位置，不将L扩进去
	large := R                              // 大于区域初始位置，R，先将R包含进去
	i := L                                  // i从L位置开始
	// 当i没有与大于区域撞上
	for i < large {
		// 如果 nums[i] < nums[R]，则将i位置与左侧区域下一个位置交换，并将小于区域往右扩，i跳下一位
		if nums[i] < nums[R] {
			utils.Swap(nums, small+1, i)
			small++
			i++
		} else if nums[i] == nums[R] {
			// 如果 nums[i] == nums[R] ，则i直接跳下一位
			i++
		} else {
			// 如果 nums[i] > nums[R]，则与大于区域左侧位置交换，大于区域往左扩，i留在原地
			utils.Swap(nums, large-1, i)
			large--
		}
	}
	// 最后将R位置，与large位置交换
	utils.Swap(nums, R, large)
	// 最终相同的区域，是在small+1，large的区间
	return small + 1, large
}

// =============== 使用非递归的方式，实现随机快排 ===========================
