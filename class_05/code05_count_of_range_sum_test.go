package class_05

import (
	"fmt"
	"testing"
)

// TestCountOfRangeSum 测试区间和问题
func TestCountOfRangeSum(t *testing.T) {
	arr := []int{-2, 5, -1}
	result := countRangeSum(arr, -2, 2)
	fmt.Printf("结果为:%d\n", result)
	// 测试10万次
	//for i := 0; i < 100000; i++ {
	//	// 随机生成一个数组
	//	arr := utils.GenerateRandomSlice(1000, -500, 500)
	//	//arr = []int{0, 2, 7, 8, 3, 4, 6, 8, 7}
	//	// 拷贝数组，用于对数器函数测试
	//	arrCopy := utils.Copy(arr)
	//	reversePairs := ReversePairs(arr)
	//	forTest := ReversePairsForTest(arrCopy)
	//	if reversePairs != forTest {
	//		t.Errorf("测试失败：获取到的逆序对个数：%d, 测试函数获取的个数：%d", reversePairs, forTest)
	//		return
	//	}
	//}
	//t.Log("测试成功！！！")
}
