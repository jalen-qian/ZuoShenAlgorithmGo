package class_16_17

// 图的表达

// Node 图节点
type Node struct {
	Value int     // 值
	In    int     // 入度，连接到当前节点的边条数（直接边）
	Out   int     // 出度，从这个节点出去的边条数（直接边）
	Nexts []*Node // 从当前节点出发的直接邻居
	Edges []*Edge // 从当前节点出发的边
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		In:    0,
		Out:   0,
		Nexts: make([]*Node, 0),
		Edges: make([]*Edge, 0),
	}
}

// Edge 图的边
type Edge struct {
	Weight int   // 边的权重
	From   *Node // 从哪个顶点来
	To     *Node // 连接到哪个顶点
}

func NewEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

// Graph 图，图是顶点和边的集合
type Graph struct {
	// 节点集合，之所以key是int类型，是因为题目会给定一个整数让你生成对应的节点，这里顺便记录下整数和节点的对应关系
	Nodes map[int]*Node
	Edges map[*Edge]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[*Edge]struct{}),
	}
}
