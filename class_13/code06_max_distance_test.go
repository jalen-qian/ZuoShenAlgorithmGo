package class_13

import "testing"

func TestMaxDistance(t *testing.T) {
	root := &TreeNode{Val: 10,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 8}},
		Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 7}},
	}
	PrintBT(root)
	t.Log(MaxDistance(root))
}
