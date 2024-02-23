package class_16_17

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
)

// 图的拓扑序，解法2，使用最大深度
// 解法1使用的是“点次”的概念，这里也可以用最大深度
// 最大深度指的是，对于任意的一个节点x，沿着最长的路径走到底经过的节点个数。
// 假设有两个节点 x 和 y，如果x的最大深度 > y的最大深度，则 x的拓扑序一定 <= y
// 因为和点次同理，如果y是x的某个子路径上的节点，则y的最大深度一定不可能小于x:
// 		1. 如果y的最大深度所在路径就是x最大深度所在路径的一部分，那么x的最大路径至少都会比y的最大深度大1
//      2. 如果y的最大深度所在路径不是x最大深度所在路径的一部分，则x必然有一条比y的最大深度更深的子路径，则x的最大深度至少比y的最大深度大2
// 综上，我们可以求出所有节点的最大深度，然后按照最大深度从大到小的顺序返回结果

// TopSortDFS2 传入一个图，返回拓扑序排序的节点列表
func TopSortDFS2(graph []*DirectedGraphNode) []*DirectedGraphNode {
	order := make(map[*DirectedGraphNode]*Record1)
	// 获取全部节点的点次缓存
	for _, node := range graph {
		f1(node, order)
	}
	var recordArr []*Record1
	for _, record := range order {
		recordArr = append(recordArr, record)
	}
	// 排序，点次大的排前面
	sort.Slice(recordArr, func(i, j int) bool {
		return recordArr[i].Deep >= recordArr[j].Deep
	})
	// 返回结果
	ans := make([]*DirectedGraphNode, 0, len(recordArr))
	for _, record := range recordArr {
		ans = append(ans, record.Node)
	}
	return ans
}

// Record1 保存每个节点以及其最大深度
type Record1 struct {
	Node *DirectedGraphNode
	Deep int
}

// 给定一个图的节点，返回这个节点的最大深度个数，使用递归
// cur 给定的节点
// cache 缓存
func f1(cur *DirectedGraphNode, cache map[*DirectedGraphNode]*Record1) *Record1 {
	if _, ok := cache[cur]; ok {
		return cache[cur]
	}
	var follow int
	for _, neighbor := range cur.Neighbors {
		follow = utils.Max(f1(neighbor, cache).Deep, follow)
	}
	record := &Record1{Deep: follow + 1, Node: cur}
	cache[cur] = record
	return record
}
