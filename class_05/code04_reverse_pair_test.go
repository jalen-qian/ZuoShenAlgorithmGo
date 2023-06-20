package class_05

import (
	"testing"

	"ZuoShenAlgorithmGo/utils"
)

// 逆序对问题对数器测试

func TestReversePair(t *testing.T) {
	// 测试10万次
	for i := 0; i < 1; i++ {
		// 随机生成一个数组
		arr := utils.GenerateRandomSlice(1000, -500, 500)
		arr = []int{0, 2, 7, 8, 3, 4, 6, 8, 7}
		// 拷贝数组，用于对数器函数测试
		arrCopy := utils.Copy(arr)
		reversePairs := ReversePair(arr)
		forTest := ReversePairForTest(arrCopy)
		if reversePairs != forTest {
			t.Errorf("测试失败：获取到的逆序对个数：%d, 测试函数获取的个数：%d", reversePairs, forTest)
			return
		}
	}
	t.Log("测试成功！！！")
}
