package class_14

import "testing"

func TestMaxSubBSTHead(t *testing.T) {
	//MaxSubBSTHead1(&TreeNode{
	//	Val:   654,
	//	Left:  &TreeNode{Val: 603, Right: &TreeNode{Val: 719}},
	//	Right: &TreeNode{Val: 249, Right: &TreeNode{Val: 656}},
	//})
	//return
	// 使用对数器测试
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		root := generateRandomBT(0, 1000, 10)
		if MaxSubBSTHead(root) != MaxSubBSTHead1(root) {
			ans1 := MaxSubBSTHead(root)
			ans2 := MaxSubBSTHead1(root)
			t.Errorf("Test failed! ans1:%p, ans2:%p ", ans1, ans2)
			return
		}
	}
	t.Log("Test successful!")
}
