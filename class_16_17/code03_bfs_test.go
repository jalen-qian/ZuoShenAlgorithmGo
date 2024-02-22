package class_16_17

import "testing"

func TestGraphBfs(t *testing.T) {
	graph := CreateGraph([][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{0, 2, 4},
		{0, 3, 5},
		{0, 3, 6},
		{0, 4, 5},
		{0, 4, 1},
		{0, 5, 6},
	})
	from := graph.Nodes[1]
	GraphBfs(from)
}
