package class_13

import "testing"

// 判断一个二叉树是否是完全二叉树
func TestIsCBT(t *testing.T) {
	// 随机生成树
	testTimes := 1000000
	for i := 0; i < testTimes; i++ {
		// 生成一颗随机树
		root := generateRandomBT(0, 1000, 20)
		if IsCBT(root) != IsCBT1(root) || IsCBT(root) != IsCBT2(root) {
			t.Error("测试失败！")
			return
		}
	}
	t.Log("测试成功！")
}
