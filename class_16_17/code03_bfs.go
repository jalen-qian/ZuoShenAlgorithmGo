package class_16_17

import "fmt"

// 图的宽度优先遍历
// 在二叉树的宽度优先遍历时，用一个队列就能搞定，但是在图的宽度优先遍历中，需要加一个set，因为图是有回路的
// 如果不加set，会使得某一个点多次进入队列，可能出现死循环。

func GraphBfs(from *Node) {
	if from == nil {
		return
	}
	queue := make([]*Node, 0) // 用切片实现一个队列
	set := make(map[*Node]struct{})
	queue = append(queue, from)
	set[from] = struct{}{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print(cur.Value, " ")
		for _, next := range cur.Nexts {
			if _, ok := set[next]; !ok {
				queue = append(queue, next)
				set[next] = struct{}{}
			}
		}
	}
	fmt.Println()
}
