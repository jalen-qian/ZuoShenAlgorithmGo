package class_16_17

import (
	"container/heap"
	"math"

	"ZuoShenAlgorithmGo/utils"
)

// 迪瑞克斯拉算法
// 给你一个有向有权重的图，且权重都不是负数，给定你一个初始节点a，请你返回一个哈希表，这个表中记录了a到所有其他节点的最短距离
//      b ———5——→ e
//    3↗︎       ↗︎1
//   a ——6——→ c ←—2— f
//    1↘︎   ↗︎2
//        d
// 那么需要返回这样一张表
// a:0  a到自己的距离是0
// b:3
// c:3  a->d->c的距离是3，比直接到c的距离6要短
// d:1
// e:4  a->d->c->e的距离和是4，比a->c->e和a->b->e两条路线的距离和都短
//
// 如果某个点从a无法到达，那么认为a到这个点的距离无穷大，不记录到表中，比如f点是a无法到达的，则不记录到哈希表
//
// 注意：迪瑞克斯拉算法，不一定要求权重都不能为负数，比如上面那张图，将d->c的权重改成-1也可以，那么a到c的距离就变成了0
// a到e的距离变成了1
// 迪瑞克斯拉算法真正规定的是：如果有环时，环路的权重和不能是负数，比如：
//     c
//  1↙︎  ↖︎-3
// a ————→ b
//     1
// 请问a到c的距离最小是多少，我们发现 a->b->c的距离是-2，但是我们绕一圈 a->b->c->a->b->c 距离变成了-3
// 也就是每绕一圈，距离都减少1，可以无限循环，无限减少。
// 为了简单起见，一般考察迪瑞克斯拉算法的题目，都直接规定所有边的权重都是非负数。
// 如果与现实结合，一般也不会出现负权重。比如权重表示过收费站要收的费用不可能是负的。
//
// 具体步骤如下：
//	1）Dijkstra算法必须指定一个源点
//	2）生成一个源点到各个点的最小距离表，一开始只有一条记录，即原点到自己的最小距离为0，源点到其他所有点的最小距离都为正无穷大
//	3）从距离表中拿出没拿过记录里的最小记录，通过这个点发出的边，更新源点到各个点的最小距离表，不断重复这一步
//	4）源点到所有的点记录如果都被拿过一遍，过程停止，最小距离表得到了

// Dijkstra1 给定一个初始点，返回这个点所有能到达的点的最小距离
// 这个算法有个优化点，getMinDistanceAndUnselectedNode函数每次都将已经出现过的点都遍历一遍，导致时间复杂度较高
// 当前算法的时间复杂度是O(N^2)，要优化这个问题，可以使用手写堆，将时间复杂度降低到O(N*logN)
func Dijkstra1(from *Node) map[*Node]int {
	if from == nil {
		return nil
	}
	distanceMap := make(map[*Node]int)        // 存储from节点到其他节点的最小距离表
	distanceMap[from] = 0                     // 自己到自己的距离是0
	selectedNode := utils.NewHashSet[*Node]() // 已经固定的点
	// 这个函数从已知的距离的点中，取一个距离最小的点
	minNode := getMinDistanceAndUnselectedNode(distanceMap, selectedNode)
	count := 0
	for minNode != nil {
		count++
		// 依次考察每条边
		for _, edge := range minNode.Edges {
			// 如果这条边通向的点没有出现过，则记录距离是最短距离+边权重
			toNode := edge.To
			if _, ok := distanceMap[toNode]; !ok {
				distanceMap[toNode] = distanceMap[minNode] + edge.Weight
			} else {
				// 如果这条边通向的点出现过，则比较新的距离是否比旧的距离小，如果小就更新
				distanceMap[toNode] = utils.Min(distanceMap[toNode], distanceMap[minNode]+edge.Weight)
			}
		}
		// 所有边都考察完了，当前节点就固定下来
		selectedNode.Add(minNode)
		// 继续取一个距离最短的点来考察
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selectedNode)
	}
	return distanceMap
}

// 从距离列表中取一个距离最短的点，并且这个点不能在已经考查完的点中
func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, selectedNodes *utils.HashSet[*Node]) *Node {
	var minNode *Node
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		if !selectedNodes.Contains(node) && distance < minDistance {
			minNode, minDistance = node, distance
		}
	}
	return minNode
}

// Dijkstra2 迪瑞克斯拉算法优化版
// Dijkstra1 的 getMinDistanceAndUnselectedNode() 函数每次调用都会将已经出现过的点全部遍历一遍，才能找到距离最小的点。
// 如果将已经加入的点放入一个堆中，每次弹出距离最小的，就能将时间复杂度优化到O(logN)了，但是有个问题，每次执行，迪瑞克斯拉算法会
// 更新minNode所有的直接边相关的点到给定点的最近距离，这样堆中的数据就不满足堆结构了，需要重新调整成堆。如果是Java等语言，系统提供
// 的堆做不到，需要自己实现加强堆。但是在Go中，堆是实现系统的接口实现的，比较灵活（或者说Go中的堆都是手写加强堆，没有系统直接提供直接用的堆），
// 而且系统也提供了 heap.Fix() 函数来满足破坏堆结构后重新调整。可以通过实现系统的堆接口来实现。
func Dijkstra2(from *Node) map[*Node]int {
	if from == nil {
		return nil
	}
	// 1.初始化堆
	h := newNodeHeap()
	// 2.将from节点添加到堆中，from节点到自己的距离必定是0
	h.AddOrUpdateOrIgnore(from, 0)
	result := make(map[*Node]int)
	// 3.堆不为空，就一直循环
	for !h.isEmpty() {
		// 3.1 先弹出距离from最近的节点
		record := heap.Pop(h).(nodeRecord)
		// 3.2 遍历这个节点发出的边，添加或者更新这些边到达的节点的最近距离
		for _, edge := range record.node.Edges {
			h.AddOrUpdateOrIgnore(edge.To, record.distance+edge.Weight)
		}
		result[record.node] = record.distance
	}
	return result
}

type nodeHeap struct {
	nodes []*Node
	// 记录每个节点当前的距离，注意：如果弹出，不要删除对应的记录，这个map中的记录要作为添加过的证据
	// 如果节点被弹出了，则将distanceMap中的对应距离改成-1，-1作为特殊标识，表示这个节点加入过，但是已经被弹出了
	distanceMap map[*Node]int
	indexMap    map[*Node]int // 记录每个节点当前的位置
}

func newNodeHeap() *nodeHeap {
	return &nodeHeap{
		distanceMap: make(map[*Node]int),
		indexMap:    make(map[*Node]int),
	}
}

type nodeRecord struct {
	node     *Node
	distance int
}

// AddOrUpdateOrIgnore 将某个节点的最进距离保存在堆中。
// 如果不在堆中，则添加
// 如果在堆中，但是距离更近，则更新，并重新调整堆
// 如果在堆中，但是距离相等或者更远，则什么都不用做，忽略
func (n *nodeHeap) AddOrUpdateOrIgnore(node *Node, distance int) {
	// 没添加过到堆，则添加
	if !n.hasAdded(node) {
		heap.Push(n, nodeRecord{
			node:     node,
			distance: distance,
		})
	}
	// 添加过，且当前就在堆中，没有被弹出，且新加入的距离更短，则更新这个节点的距离，并且重新调整成堆结构
	if n.isEntered(node) && n.distanceMap[node] > distance {
		// 在堆中，但是现在要更新的距离更近，则重新调整堆
		n.distanceMap[node] = distance
		heap.Fix(n, n.indexMap[node])
	}
	// 否则，忽略
}

func (n *nodeHeap) Len() int {
	return len(n.nodes)
}

func (n *nodeHeap) isEmpty() bool {
	return n.Len() == 0
}

// 是否添加过到堆
func (n *nodeHeap) hasAdded(node *Node) bool {
	_, ok := n.distanceMap[node]
	return ok
}

// 是否当前就在堆中
func (n *nodeHeap) isEntered(node *Node) bool {
	return n.hasAdded(node) && n.distanceMap[node] != -1
}

// Less 定义谁排前面，这里是与给定点实时最短距离小的排前面
func (n *nodeHeap) Less(i, j int) bool {
	distanceI, distanceJ := n.distanceMap[n.nodes[i]], n.distanceMap[n.nodes[j]]
	return distanceI < distanceJ
}

func (n *nodeHeap) Swap(i, j int) {
	// 索引交换
	n.indexMap[n.nodes[i]], n.indexMap[n.nodes[j]] = n.indexMap[n.nodes[j]], n.indexMap[n.nodes[i]]
	// 数组中的位置交换
	n.nodes[i], n.nodes[j] = n.nodes[j], n.nodes[i]
}

// Push 确保 push的类型是 nodeRecord 类型
func (n *nodeHeap) Push(x any) {
	record := x.(nodeRecord)
	// 推入一个新的节点和距离
	// 按照Go堆的实现规则，直接加入到最后面
	n.nodes = append(n.nodes, record.node)
	// 记录坐标在最后面
	n.indexMap[record.node] = len(n.nodes) - 1
	n.distanceMap[record.node] = record.distance
}

// Pop 确保弹出的类型也是 nodeRecord 类型
func (n *nodeHeap) Pop() any {
	v := n.nodes[len(n.nodes)-1]
	n.nodes = n.nodes[:len(n.nodes)-1]
	// 弹出，删除indexMap
	delete(n.indexMap, v)
	// 不能删除distanceMap中的值，distanceMap同时作为是否添加过的证据，设置距离为-1
	distance := n.distanceMap[v]
	n.distanceMap[v] = -1
	return nodeRecord{
		node:     v,
		distance: distance,
	}
}

var _ heap.Interface = (*nodeHeap)(nil)
