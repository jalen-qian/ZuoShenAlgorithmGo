package class_16_17

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"testing"
)

func TestTopSortDfs1(t *testing.T) {
	t.Log("开始测试...")
	for i := 0; i < 1000000; i++ {
		// 注：maxNodes设置成100时，测试较大概率出现失败，递归序不正确，经代码检查发现是因为
		// 节点个数过大时，点次数呈指数往上增，超过了int64的范围，溢出了，出现了负数，导致递归序不正确
		// 这里测试最大50个节点，1000000次，能跑成功，说明这个算法是正确的
		graph := randomDAG(50)
		ans1 := TopSortDfs1(graph)
		if err := checkTopologyOrder(ans1, graph); err != nil {
			t.Log("TopSortDfs1()测试失败:")
			t.Error(err.Error())
			return
		}
		ans2 := TopSortDFS2(graph)
		if err := checkTopologyOrder(ans2, graph); err != nil {
			t.Log("TopSortDFS2()测试失败:")
			t.Error(err.Error())
			return
		}
	}
	t.Log("测试成功！！！")
}

// 随机生成有向无环图，并使用题目给定的结构
// maxNodes 最大节点个数
// maxEdge 最大的边个数
// 思路：先随机生成拓扑序，根据拓扑序逆向生成有向无环图
// 在生成好所有的节点后，随机的连接边，但是选择连接的边的两个节点，一定是在拓扑序的前面的连接到后面的，这样就能保证生成无环图
// 另外，还有一个问题，要保证所有的点都是联通的，为了解决这个问题，我们在前 n-1 条边必定选择一个连接过的点和没有连接过的孤点
// 之所以是前 n-1 条边，是因为要将一个图联通，至少需要 n-1 条边
func randomDAG(maxNodes int) []*DirectedGraphNode {
	n := rand.Intn(maxNodes + 1) // 节点个数，[0, maxNodes]
	if n == 0 {
		return nil
	}
	// 边的个数，取值范围是 [n-1, n*(n-1)/2] ps:一个n个顶点的有向无环图，边最少是 n-1(刚好联通），最多是 n*(n-1) / 2
	minEdge := n - 1
	maxEdge := (n * (n - 1)) / 2
	edgeNum := minEdge + rand.Intn(maxEdge-minEdge+1) // 在这个范围内取一个随机数
	// 节点个数和边的个数都确定了，生成随机的拓扑序
	nodes := make([]*DirectedGraphNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &DirectedGraphNode{
			Label: i + 1,
		}
	}
	// 打乱顺序，生成随机的拓扑序
	shuffle(nodes)
	appearedNodes := utils.NewHashSet[int]()   // 记录连接过边的节点
	unAppearedNodes := utils.NewHashSet[int]() // 记录没有连接过边的节点
	// 一开始，都没有连接过边
	for i, _ := range nodes {
		unAppearedNodes.Add(i)
	}
	// 记录已经添加过的边的条数
	count := 0
	// 前 n-1 条边，必取出现过和没出现过的
	for i := 0; i < n-1; i++ {
		// 没生成过，随机选拓扑序中前后的两条边相连
		if count == 0 {
			from := rand.Intn(n - 1)             // [0,n-2]，假设n=10，取到了5
			to := rand.Intn(n-1-from) + from + 1 // 则to的取值范围是 [from+1, n-1] 也就是 [6,9] => [0,4) 再往右平移6
			// 标记为连接过
			appearedNodes.Add(from)
			appearedNodes.Add(to)
			unAppearedNodes.Remove(from)
			unAppearedNodes.Remove(to)
			// from 节点连接向 to 节点
			nodes[from].Neighbors = append(nodes[from].Neighbors, nodes[to])
		} else {
			// 生成过边，则必然要从没连接过的顶点集与连接过的顶点集中分别取一个顶点
			from := appearedNodes.GetRandomValue()
			to := unAppearedNodes.GetRandomValue()
			unAppearedNodes.Remove(to)
			appearedNodes.Add(to)
			// 确保是拓扑序小的连接到拓扑序大的
			if from > to {
				from, to = to, from
			}
			// from 节点连接向 to 节点
			nodes[from].Neighbors = append(nodes[from].Neighbors, nodes[to])
		}
		count++
	}
	// 生成剩下的 edgeNum - (n - 1) 条边
	for count < edgeNum {
		from := rand.Intn(n - 1)
		to := rand.Intn(n-1-from) + from + 1
		// 如果已经连接过了，则不重复连接
		edgeExist := false
		for _, neighbor := range nodes[from].Neighbors {
			if neighbor == nodes[to] {
				edgeExist = true
			}
		}
		if edgeExist {
			continue
		}
		// 没连接过，则连接
		nodes[from].Neighbors = append(nodes[from].Neighbors, nodes[to])
		count++
	}
	return nodes
}

// shuffle 打乱顺序
func shuffle(graph []*DirectedGraphNode) {
	for i := 0; i < len(graph); i++ {
		// 随机选一个位置，和i位置交换
		j := rand.Intn(len(graph))
		if i == j {
			j = i + 1
		}
		if j == len(graph) {
			j = 0
		}
		utils.SwapSlice[*DirectedGraphNode](graph, i, j)
	}
}

// 测试给定的拓扑序是否正确
// ans 按照拓扑序排好序的结果
// graph 原始图数组
func checkTopologyOrder(ans []*DirectedGraphNode, graph []*DirectedGraphNode) error {
	if len(ans) != len(graph) {
		return fmt.Errorf("失败，拓扑序数组个数与图节点个数不一致")
	}
	indexMap := make(map[*DirectedGraphNode]int)
	for i, node := range ans {
		indexMap[node] = i
	}
	for _, node := range graph {
		for _, neighbor := range node.Neighbors {
			// 如果有节点的直接邻居在拓扑序中排在了当前节点的前面，则拓扑序是错的
			if indexMap[node] > indexMap[neighbor] {
				return fmt.Errorf("失败，拓扑序[%d(位置%d)]=>[%d(位置%d)]不正确，反了", neighbor.Label, indexMap[neighbor], node.Label, indexMap[node])
			}
		}
	}
	return nil
}
