package class_13

import "testing"

func TestIsBST(t *testing.T) {
	t.Log("Test starting...")
	for i := 0; i < 500000; i++ {
		root := generateRandomBT(0, 1000, 20)
		isBst := IsBST(root)
		isBst2 := IsBST2(root)
		if isBst != isBst2 {
			t.Errorf("Testing failed \n isBst:%v \n isBst2:%v", isBst, isBst2)
			return
		}
	}
	t.Log("Test succeed !!!")
}
