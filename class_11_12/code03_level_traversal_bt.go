package class_11_12

import (
	"ZuoShenAlgorithmGo/class_03"
	"fmt"
)

// 二叉树按层遍历

type LevelTraversalBT struct{}

// Level 实现二叉树按层遍历
func (r *LevelTraversalBT) Level(root *Node) {
	if root == nil {
		return
	}
	queue := class_03.NewMyQueue[*Node]()
	// 头先入队列
	queue.Push(root)
	for !queue.IsEmpty() {
		cur := queue.Poll()
		fmt.Printf("%d ", cur.Value)
		if cur.Left != nil {
			queue.Push(cur.Left)
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
		}
	}
	fmt.Println()
}
