package class_13

import "testing"

func TestIsFull(t *testing.T) {
	t.Log("Test starting...")
	for i := 0; i < 500000; i++ {
		root := generateRandomBT(0, 1000, 20)
		isFull := IsFull(root)
		isFull2 := IsFull2(root)
		if isFull != isFull2 {
			t.Errorf("Testing failed \n isFull:%v \n isFull2:%v", isFull, isFull2)
			return
		}
	}
	t.Log("Test succeed !!!")
}
