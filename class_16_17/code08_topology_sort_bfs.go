package class_16_17

// 图的拓扑排序真题：https://www.lintcode.com/problem/127/
// 使用宽度优先遍历解决这个问题，方法与 code05_topology_sort.go 的一致，入度为0的先遍历，
// 遍历完消除入度为0的节点的影响，然后继续。使用一个队列来实现。
// 由于题目给定的接口没有入度和出度，所以我们使用一张哈希表来标记每个节点的入度（这个算法只需要入度）

func TopSortBFS(graph []*DirectedGraphNode) []*DirectedGraphNode {
	// 1.使用一张哈希表来记录每个节点的入度
	inMap := make(map[*DirectedGraphNode]int)
	// 初始时入度都标记为0
	for _, node := range graph {
		inMap[node] = 0
	}
	// 统计入度并保存
	for _, node := range graph {
		for _, neighbor := range node.Neighbors {
			inMap[neighbor]++
		}
	}
	// 2.使用一个队列，入度为0的入队
	queue := make([]*DirectedGraphNode, 0)
	for _, node := range graph {
		if inMap[node] == 0 {
			queue = append(queue, node)
		}
	}
	result := make([]*DirectedGraphNode, 0, len(graph))
	// 3.当队列不为空时，一直循环
	for len(queue) > 0 {
		// 出队加入到结果
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)
		// 消除当前节点的影响
		for _, neighbor := range cur.Neighbors {
			inMap[neighbor]--
			// 入度为0的重新入队
			if inMap[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}
