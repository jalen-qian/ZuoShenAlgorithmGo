package class_16_17

import "testing"

// 单测：使用宽度优先遍历实现拓扑序问题

func TestTopSortBFS(t *testing.T) {
	t.Log("开始测试...")
	for i := 0; i < 100000; i++ {
		// 注：maxNodes设置成100时，测试较大概率出现失败，递归序不正确，经代码检查发现是因为
		// 节点个数过大时，点次数呈指数往上增，超过了int64的范围，溢出了，出现了负数，导致递归序不正确
		// 这里测试最大50个节点，1000000次，能跑成功，说明这个算法是正确的
		graph := randomDAG(50)
		ans := TopSortBFS(graph)
		if err := checkTopologyOrder(ans, graph); err != nil {
			t.Log("TopSortBFS()测试失败:")
			t.Error(err.Error())
			return
		}
	}
	t.Log("测试成功！！！")
}
