package class_05

import (
	"ZuoShenAlgorithmGo/utils"
	"testing"
)

// SmallSumForTest 对数器，使用O(N^2)的方式实现
func SmallSumForTest(arr []int) int {
	var ans int
	for i := 0; i < len(arr); i++ {
		// 从i+1开始找，找到比i位置大的数，则将i位置的数累加
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				ans += arr[i]
			}
		}
	}
	return ans
}

func TestSmallSum(t *testing.T) {
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(999, -500, 500)
		//arr := []int{2, 6, 1, 3, 5, 0, 4, 8}
		arr1 := utils.Copy(arr)
		s1 := SmallSum(arr)
		s2 := SmallSumForTest(arr1)
		if s1 != s2 {
			t.Errorf("测试出错了！")
			return
		}
	}
	t.Log("测试成功！！！")
}
