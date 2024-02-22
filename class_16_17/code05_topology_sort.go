package class_16_17

/*
图的拓扑排序
在一个有向无环图中，对于图中的每一条有向边 A -> B , 在拓扑排序中A一定在B之前.
拓扑排序中的第一个节点可以是图中的任何一个没有其他节点指向它的节点.

在程序编译过程中，包与包之间的存在复杂的依赖关系，先编译哪个包后编译哪个包，这就是有一个图的拓扑序问题。
比如：
      B
  ↗︎
A —→  C —→ D —→ F
  ↘︎  ↓
      E
则编译顺序可以是：A B C E D F
也可以是 A C B D E F
拓扑序的排序可以有多个，但是都是正确的。

可以知道：谈到拓扑排序，一定是一个有序的无环图，否则没有意义，因为有环图互相依赖，没有办法确定谁先开始。
下面这个图，从哪个开始都不行，因为循环依赖了。这也是为何我们导包时如果出现循环引用，编译程序会报错。
        1 ——→ 2
        ↘︎   ↙︎
           3

*/

// SortedTopology 给定一个有序的无环图，返回按照拓扑序排序的数组
// 解决思路：如果是一个有序无环图，则一定有入度为0的节点，先遍历入度为0的节点，遍历之后，将这些入度为0节点邻居的入度抹去1
// 原因：入度为0表示没有被其他节点依赖，或者依赖已经解决了。
// 会再次产生入度为0的节点，重复上面的步骤。每找到一个入度为0的节点，则加入到结果集中。
// 为了避免破坏图的经典结构，我们使用一个哈希表存储每个节点的剩余入度。
func SortedTopology(graph *Graph) []*Node {
	inMap := make(map[*Node]int)    //存储节点的剩余入度
	zeroInQueue := make([]*Node, 0) // 队列，入度为0的才入队
	result := make([]*Node, 0, len(graph.Nodes))
	for _, node := range graph.Nodes {
		inMap[node] = node.In // 记录每个节点的入度
		if node.In == 0 {     // 入度为0，则加入到结果集中
			zeroInQueue = append(zeroInQueue, node)
		}
	}
	for len(zeroInQueue) > 0 {
		cur := zeroInQueue[0]
		zeroInQueue = zeroInQueue[1:]
		result = append(result, cur)
		// 将这个入度为0的节点，直接邻居的入度都减1
		for _, next := range cur.Nexts {
			inMap[next]--
			if inMap[next] == 0 {
				zeroInQueue = append(zeroInQueue, next)
			}
		}
	}
	return result
}
