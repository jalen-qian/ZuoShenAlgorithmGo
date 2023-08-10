package class_03

import (
	"ZuoShenAlgorithmGo/utils"
	"testing"
)

func TestGetMax(t *testing.T) {
	t.Log("测试开始")
	testTimes := 200000
	for i := 0; i < testTimes; i++ {
		//t.Logf("测试第%d次", i+1)
		// 生成一个随机数组
		arr := utils.GenerateRandomSlice(1000, -1000, 1000)
		max1 := GetMax(arr)
		max2 := GetMaxForTest(arr)
		if max1 != max2 {
			t.Fatalf("获取到不同的最大值，max1:%d,max2:%d\n", max1, max2)
		}
	}
	t.Log("测试结束")
}
