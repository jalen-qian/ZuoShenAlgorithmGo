package class_16_17

import (
	"ZuoShenAlgorithmGo/utils"
	"container/heap"
)

// 最小生成树之 Prim 算法，也叫P算法

// 也是先将所有的边隐藏，并选择任意一个节点作为出发节点，每选择一个点，就解锁其对应的边，再从已解锁边中找个最小的，解锁新的点。
// 过程就是点->边->点... 一直到所有的点都选中，停止。
//
//	1）可以从任意节点出发来寻找最小生成树
//	2）某个点加入到被选取的点中后，解锁这个点出发的所有新的边
//	3）在所有解锁的边中选最小的边，然后看看这个边会不会形成环
//	4）如果会，不要当前边，继续考察剩下解锁的边中最小的边，重复3）
//	5）如果不会，要当前边，将该边的指向点加入到被选取的点中，重复2）
//	6）当所有点都被选取，最小生成树就得到了

// 这里我们使用一个小根堆来存储被解锁的边，使用一个集合来存储被选中的节点
//

func PrimMST(graph *Graph) *utils.HashSet[*Edge] {
	// 初始化一个小根堆
	smallHeap := NewEdgeHeap()
	// 存储哪些点被解锁出来了
	nodeSet := utils.NewHashSet[*Node]()
	result := utils.NewHashSet[*Edge]()
	for _, node := range graph.Nodes {
		nodeSet.Add(node)
		// 解锁这个点对应的所有边
		for _, edge := range node.Edges {
			heap.Push(smallHeap, edge)
		}
		// 从解锁的边中弹出最小的边，看是否会形成环
		for smallHeap.Len() > 0 {
			edge := heap.Pop(smallHeap).(*Edge)
			// 如果不包含，就是新的点，要这个点
			if !nodeSet.Contains(edge.To) {
				nodeSet.Add(edge.To)
				result.Add(edge)
				// 继续解锁新的点对应的所有边
				for _, e := range edge.To.Edges {
					heap.Push(smallHeap, e)
				}
			}
		}
		break
	}
	return result
}
