package class_11_12

import (
	"fmt"
	"testing"
)

// 二叉树按层遍历测试

func TestLevelTraversalBT_Level(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
	}
	PrintBT(root)
	fmt.Println()

	new(LevelTraversalBT).Level(root)
}
