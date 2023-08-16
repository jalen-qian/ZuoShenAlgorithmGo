package class_11_12

import (
	"fmt"
	"testing"
)

// 二叉树按层遍历测试

func TestLevelTraversalBT_Level(t *testing.T) {
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3, Left: &Node{Value: 6}, Right: &Node{Value: 7}},
	}
	PrintBT(root)
	fmt.Println()

	new(LevelTraversalBT).Level(root)
}
