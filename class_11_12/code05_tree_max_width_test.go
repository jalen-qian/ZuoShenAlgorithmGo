package class_11_12

import "testing"

// 给定一颗二叉树，返回最大的宽度，测试

func TestMaxWidth(t *testing.T) {
	t.Log("测试开始...")
	for i := 0; i < 500000; i++ {
		root := generateRandomBT(0, 100, 20)
		maxWidth := MaxWidth(root)
		maxWidth1 := MaxWidthWithMap(root)
		if maxWidth != maxWidth1 {
			t.Errorf("测试失败:\n maxWidth:%d\n maxWidth1:%d\n", maxWidth, maxWidth1)
			return
		}
	}
	t.Log("测试成功！！！")
}
