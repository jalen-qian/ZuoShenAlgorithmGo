package class_06

import "ZuoShenAlgorithmGo/utils"

/**
荷兰国旗问题：
给定一个数组nums和一个数x，将这个数组划分为 <x 的在左边 =x 的在中间 >x 的在右边
要求：1.时间复杂度O(N) 2.除了少数几个变量，不分配额外空间
*/

// FlagOfNetherlands 荷兰国旗问题，简易版
func FlagOfNetherlands(nums []int, x int) {
	if len(nums) <= 1 {
		return
	}
	small := -1 // 小于区域
	i := 0
	for i <= len(nums)-1 {
		// 如果 nums[i] <= x，则将i位置与左侧区域下一个位置交换，并将小于区域往右扩，i跳下一位
		if nums[i] <= x {
			utils.Swap(nums, small+1, i)
			small++
		}
		// 如果 num[i] > x，则i跳下一位
		i++
	}
}

// FlagOfNetherlands1 荷兰国旗问题，复杂版
func FlagOfNetherlands1(nums []int, x int) {
	if len(nums) <= 1 {
		return
	}
	small := -1        // 小于区域初始位置
	large := len(nums) // 大于区域初始位置
	i := 0
	// 当i没有与大于区域撞上
	for i < large {
		// 如果 nums[i] < x，则将i位置与左侧区域下一个位置交换，并将小于区域往右扩，i跳下一位
		if nums[i] < x {
			utils.Swap(nums, small+1, i)
			small++
			i++
		} else if nums[i] == x {
			// 如果 nums[i] == x ，则i直接跳下一位
			i++
		} else {
			// 如果 nums[i] > x，则与大于区域左侧位置交换，大于区域往左扩，i留在原地
			utils.Swap(nums, large-1, i)
			large--
		}
	}
}

// FlagOfNetherlands2 荷兰国旗问题，复杂版2
func FlagOfNetherlands2(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	small := -1         // 小于区域初始位置
	large := length - 1 // 大于区域初始位置
	i := 0
	// 当i没有与大于区域撞上
	for i < large {
		// 如果 nums[i] < nums[length-1]，则将i位置与左侧区域下一个位置交换，并将小于区域往右扩，i跳下一位
		if nums[i] < nums[length-1] {
			utils.Swap(nums, small+1, i)
			small++
			i++
		} else if nums[i] == nums[length-1] {
			// 如果 nums[i] == nums[length-1] ，则i直接跳下一位
			i++
		} else {
			// 如果 nums[i] > nums[length-1]，则与大于区域左侧位置交换，大于区域往左扩，i留在原地
			utils.Swap(nums, large-1, i)
			large--
		}
	}
	// 最后将n-1位置，与large位置交换
	utils.Swap(nums, length-1, large)
}
