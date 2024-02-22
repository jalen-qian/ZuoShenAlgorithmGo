package class_16_17

// 图的表达多种多样，这里将其中一种表达转换成我们的 Node Edge 和 Graph 的表达

// CreateGraph 将二维矩阵表达的图，转换成我们熟悉的结构
// matrix 的每一行数据表示图的一条边
// 这是一个N*3的矩阵，也就是每行固定3个整数
// [weight, from节点上面的值, to节点上面的值]
// [ 5, 0, 7]
// [ 3, 0, 1]
func CreateGraph(matrix [][]int) *Graph {
	// 先初始化一个图
	graph := NewGraph()
	// 遍历，每一行都是一条边
	for _, e := range matrix {
		weight := e[0]
		from := e[1]
		to := e[2]
		// 如果没有创建过from节点，则创建
		if _, ok := graph.Nodes[from]; !ok {
			graph.Nodes[from] = NewNode(from)
		}
		// 如果没有创建过to节点，则创建
		if _, ok := graph.Nodes[to]; !ok {
			graph.Nodes[to] = NewNode(to)
		}
		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		// 创建当前边
		edge := NewEdge(weight, fromNode, toNode)
		// 将toNode添加到fromNode的Next集合
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		// fromNode的出度++
		fromNode.Out++
		// toNode的入度++
		toNode.In++
		// fromNode的边++
		fromNode.Edges = append(fromNode.Edges, edge)
		// 将边添加到图的边集合中
		graph.Edges[edge] = struct{}{}
	}
	return graph
}
