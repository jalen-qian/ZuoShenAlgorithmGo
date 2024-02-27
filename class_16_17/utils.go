package class_16_17

import (
	"math/rand"
)

// 工具
// 生成随机的图

// GenerateGraph 生成一个随机图
func GenerateGraph(maxNodes int, maxEdges int, maxWeight int) *Graph {
	n := rand.Intn(maxNodes + 1) // 节点个数，[0, maxNodes]
	g := NewGraph()
	if n == 0 {
		return g
	}
	// 边的个数也随机
	edgeN := rand.Intn(maxEdges + 1)
	// 边的个数最大值是n * (n - 1)
	if edgeN > n*(n-1) {
		edgeN = n * (n - 1)
	}
	// 创建n个节点
	for i := 0; i < n; i++ {
		g.Nodes[i+1] = NewNode(i + 1) // value:[1,n]
	}
	var edges []edgeRecord
	// 一共n个顶点，则会产生n*(n-1)条边的可能性
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j {
				edges = append(edges, edgeRecord{
					from: i,
					to:   j,
				})
			}
		}
	}
	leftEdgeN := len(edges) // 还剩下的边的条数
	for i := 0; i < edgeN; i++ {
		// 生成一条没生成过的边，如果生成过或者自己指向自己，则重新生成，直到生成
		edgeIndex := rand.Intn(leftEdgeN)
		weight := rand.Intn(maxWeight) + 1 // 权重取值范围 [0,maxWeight) + 1 => [1,maxWeight]
		from := edges[edgeIndex].from
		to := edges[edgeIndex].to
		edge := NewEdge(weight, g.Nodes[from], g.Nodes[to])
		// 加入到边集合中
		g.Edges[edge] = struct{}{}
		// 将边指向的节点，加入到from节点的直接邻居
		g.Nodes[from].Nexts = append(g.Nodes[from].Nexts, g.Nodes[to])
		// 拼接上直接边
		g.Nodes[from].Edges = append(g.Nodes[from].Edges, edge)
		// 更新from节点的出度和to节点的入度
		g.Nodes[from].Out++
		g.Nodes[to].In++
		// 将生成过的边剔除掉
		edges = append(edges[:edgeIndex], edges[edgeIndex+1:]...)
		leftEdgeN--
	}
	return g
}

type edgeRecord struct {
	from int
	to   int
}
