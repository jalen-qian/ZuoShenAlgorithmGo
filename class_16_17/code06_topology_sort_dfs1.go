package class_16_17

import "sort"

/**
图的拓扑排序真题:
https://www.lintcode.com/problem/127/

描述
给定一个有向图，图节点的拓扑排序定义如下:

对于图中的每一条有向边 A -> B , 在拓扑排序中A一定在B之前.
拓扑排序中的第一个节点可以是图中的任何一个没有其他节点指向它的节点.
针对给定的有向图找到任意一种拓扑排序的顺序.

这个题目的思路，可以使用深度优先遍历。先说一个概念：点次！！
过图中的某个点，将这个点能走到的路都走一遍，统计所有经过点的个数，就是这个节点的点次。
比如：
     a ———→ b
  ↗︎    ↗︎  ↓
x —→ e  —→  c

c只有自己，点次是1
b能找到c，点次是2
a(b->c) : 3
e(b->c & c ) ：5 [注意：这里有两条路，c重复计算，b(2)+c(1)+c(1)+(e自己)1 = 5]
x(a & e) : 5+3+1 = 9

那么有一个结论：如果点 x 的点次 >= 点y的点次，则x的拓扑序一定 <= y的拓扑序
这个很好理解：如果x是y的下级，则x走到的点，y也一定能走到，则x的点次一定 < y的点次
反过来，如果x是y的上级，则y经过的所有路径，x也一定能经过，所以x的点次会 > y的点次

有了上面的结论，在这个题目中，我们的思路是：统计所有节点的点次，然后按照点次大到小排序，就是正确的拓扑序。
但是这里有个问题，我统计x的点次时，已经将点a的所有路径都走了一遍，统计点a时，又要重复都走一遍，会造成大量的重复步骤。
这里说一个思想，也是动态规划的一个常用的思想：缓存！！如果某个点的点次统计完了，就缓存下来，如果下次需要，就直接获取。

*/

// DirectedGraphNode 每个点有对应的值，和有哪些直接邻居，这其实就是邻接表法的表达
type DirectedGraphNode struct {
	Label     int
	Neighbors []*DirectedGraphNode
}

func NewDirectedGraphNode(x int) *DirectedGraphNode {
	return &DirectedGraphNode{
		Label:     x,
		Neighbors: make([]*DirectedGraphNode, 0),
	}
}

// TopSortDfs1 传入一个图，返回拓扑序排序的节点列表
func TopSortDfs1(graph []*DirectedGraphNode) []*DirectedGraphNode {
	order := make(map[*DirectedGraphNode]*Record)
	// 获取全部节点的点次缓存
	for _, node := range graph {
		f(node, order)
	}
	var recordArr []*Record
	for _, record := range order {
		recordArr = append(recordArr, record)
	}
	// 排序，点次大的排前面
	sort.Slice(recordArr, func(i, j int) bool {
		return recordArr[i].Nodes >= recordArr[j].Nodes
	})
	// 返回结果
	ans := make([]*DirectedGraphNode, 0, len(recordArr))
	for _, record := range recordArr {
		ans = append(ans, record.Node)
	}
	return ans
}

// Record 保存每个节点的点次个数
type Record struct {
	Node  *DirectedGraphNode
	Nodes int64
}

// 给定一个图的节点，返回这个节点的点次个数，使用递归
// cur 给定的节点
// cache 缓存
func f(cur *DirectedGraphNode, cache map[*DirectedGraphNode]*Record) *Record {
	if _, ok := cache[cur]; ok {
		return cache[cur]
	}
	var nodes int64 = 1
	for _, neighbor := range cur.Neighbors {
		nodes += f(neighbor, cache).Nodes
	}
	record := &Record{Nodes: nodes, Node: cur}
	cache[cur] = record
	return record
}
