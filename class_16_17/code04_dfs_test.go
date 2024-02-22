package class_16_17

import "testing"

func TestGraphDfs(t *testing.T) {
	/*
	        1  ——————→ 8
	     ↙︎  ↓  ↘︎   ↙︎
	   2 ——→ 3  ——→ 4
	   ↘︎  ↗︎
	     5
	*/
	// 生成上面的图
	graph := CreateGraph([][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 4},
		{0, 1, 8},
		{0, 2, 3},
		{0, 2, 5},
		{0, 3, 4},
		{0, 5, 3},
		{0, 8, 4},
	})
	// 从1开始，深度优先遍历
	// 打印结果：1 2 3 4 5 8
	GraphDfs(graph.Nodes[1])

}
