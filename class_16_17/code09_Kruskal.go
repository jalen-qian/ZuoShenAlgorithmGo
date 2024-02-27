package class_16_17

import (
	"container/heap"

	"ZuoShenAlgorithmGo/utils"
)

// 最小生成树算法 Kruskal 也叫K算法
// 首先介绍几个概念：
//  - 连通图：在无向图中，若任意两个顶点x与y都有路径相通，则称该无向图为连通图。
//  - 强连通图：在有向图中，若任意两个顶点x与y都有路径相通，则称该有向图为强连通图。
//  - 连通网（边带权的连通图）：在连通图中，若图的边具有一定的意义，每一条边都对应着一个数，称为权；权代表着连接连个顶点的代价，称这种连通图叫做连通网。
//  - 生成树：一个连通图的生成树是指一个连通子图，它含有图中全部n个顶点，但只有足以构成一棵树的n-1条边。
//           一颗有n个顶点的生成树有且仅有n-1条边，如果生成树中再添加一条边，则必定成环。
//  - 最小生成树：在连通网的所有生成树中，所有边的权重和最小的生成树，称为最小生成树。
//
// Kruskal算法是给定一个连通图，生成这个图的最小生成树的一种算法，通俗也可以叫做“加边法”。(保留n-1条边，保证每个顶点连通，且这些边的权重和最小）
//
//
// Kruskal算法的流程：
// 1）总是从权值最小的边开始考虑，依次考察权值依次变大的边
// 2）当前的边要么进入最小生成树的集合，要么丢弃
// 3）如果当前的边进入最小生成树的集合中不会形成环，就要当前边
// 4）如果当前的边进入最小生成树的集合中会形成环，就不要当前边
// 5）考察完所有边之后，最小生成树的集合也得到了

// 这里用到并查集，因为加的边会将顶点连通，此时这两个顶点就加入到一个集合。如果后序加的边的两个顶点已经在一个集合了，说明这两个顶点已经连通了，
// 继续加这条边就会成环，此时要丢弃。这里使用并查集能很方便查询加入的边是否会导致成环。

// KruskalMST 给定一个连通图，使用 Kruskal 算法来获取最小生成树，并返回这颗最小生成树的边的集合
func KruskalMST(graph *Graph) *utils.HashSet[*Edge] {
	nodes := make([]*Node, 0, len(graph.Nodes))
	for _, node := range graph.Nodes {
		nodes = append(nodes, node)
	}
	// 初始化一个并查集
	unionFind := NewUnionFind(nodes...)
	// 实现一个小根堆，前面的代码有实现自己的堆，这里熟悉系统的heap接口的写法，自己实现
	smallHeap := NewEdgeHeap()
	// 将所有的边加入到小根堆
	for edge, _ := range graph.Edges {
		heap.Push(smallHeap, edge)
	}
	result := utils.NewHashSet[*Edge]()
	// 依次弹出，并考察是否两个顶点在一个集合，在就说明形成环了，舍弃掉，不在就加入
	for smallHeap.Len() > 0 {
		edge := heap.Pop(smallHeap).(*Edge)
		// 不在同一个集合，则加入到结果集，并将这两个顶点所在集合合并
		if !unionFind.IsSameSet(edge.From, edge.To) {
			result.Add(edge)
			unionFind.Union(edge.From, edge.To)
		}
	}
	return result
}

// UnionFind 首先实现一个并查集，集合元素是图的节点
type UnionFind struct {
	Parents map[*Node]*Node // key的父亲是value
	SizeMap map[*Node]int   // 记录某个代表节点所在集合的集合大小
}

// NewUnionFind 初始化一个并查集，并将给定的节点创建各自的集合
func NewUnionFind(nodes ...*Node) *UnionFind {
	parents := make(map[*Node]*Node)
	sizeMap := make(map[*Node]int)
	// 初始时各自是单独的集合
	for _, node := range nodes {
		parents[node] = node // 自己是自己的父亲
		sizeMap[node] = 1    // 集合大小是1
	}
	return &UnionFind{
		Parents: parents,
		SizeMap: sizeMap,
	}
}

// findFather 往上到不能再往上，找到代表节点，并将经过的链路打平
func (u *UnionFind) findFather(node *Node) *Node {
	path := make([]*Node, 0) // 将经过的节点都指向代表节点
	if u.Parents[node] != node {
		path = append(path, node)
		node = u.Parents[node]
	}
	for _, n := range path {
		u.Parents[n] = node
	}
	return node
}

// IsSameSet 判断是否在同一个集合，如果找到的代表节点是同一个，就在同一个集合
func (u *UnionFind) IsSameSet(node1, node2 *Node) bool {
	return u.findFather(node1) == u.findFather(node2)
}

// Union 合并两个集合
func (u *UnionFind) Union(node1, node2 *Node) {
	father1, father2 := u.findFather(node1), u.findFather(node2)
	// 已经是一个集合了，无需合并
	if father1 == father2 {
		return
	}
	// 不是一个集合，将较小集合的代表节点挂到较大集合
	big, small := father1, father2
	if u.SizeMap[big] < u.SizeMap[small] {
		big, small = small, big
	}
	u.Parents[small] = big
	u.SizeMap[big] += u.SizeMap[small]
	// 删除较小集合的集合大小记录
	delete(u.SizeMap, small)
}

type EdgeHeap []*Edge

func NewEdgeHeap() *EdgeHeap {
	return &EdgeHeap{}
}

func (e *EdgeHeap) Len() int {
	return len(*e)
}

func (e *EdgeHeap) Less(i, j int) bool {
	// 根据边的权重排序，小权重小的排前面，小根堆
	return (*e)[i].Weight < (*e)[j].Weight
}

func (e *EdgeHeap) Swap(i, j int) {
	tmp := (*e)[i]
	(*e)[i] = (*e)[j]
	(*e)[j] = tmp
}

func (e *EdgeHeap) Push(x any) {
	// 固定添加到末尾
	*e = append(*e, x.(*Edge))
}

func (e *EdgeHeap) Pop() any {
	// 弹出固定弹出 n-1 位置的值
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

var _ heap.Interface = (*EdgeHeap)(nil)
