package class_16_17

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

// TestDijkstra 测试迪瑞克斯拉算法
func TestDijkstra(t *testing.T) {
	t.Log("测试开始...")
	for i := 0; i < 100000; i++ {
		graph := GenerateGraph(100, 9900, 1000)
		// 随机取一个节点作为初始节点
		var startIndex int
		if len(graph.Nodes) > 0 {
			startIndex = rand.Intn(len(graph.Nodes))
		}
		startIndex = 2
		fromNode := graph.Nodes[startIndex]
		ans1 := Dijkstra1(fromNode)
		ans2 := Dijkstra2(fromNode)
		if err := checkAnswerIsEqual(ans1, ans2); err != nil {
			t.Error(err)
			return
		}
	}
	t.Log("测试成功！！！")
}

// checkAnswerIsEqual 检查两个迪瑞克斯拉算法的结果是否相同，不相同就返回error
func checkAnswerIsEqual(ans1, ans2 map[*Node]int) error {
	if len(ans1) != len(ans2) {
		return errors.New("答案不一致，答案的数量不同")
	}
	for node, distance := range ans1 {
		distance2, ok := ans2[node]
		if !ok {
			return fmt.Errorf("答案不一致，答案1的节点[%d]在答案二中不存在", node.Value)
		} else if distance2 != distance {
			return fmt.Errorf("答案不一致，节点[%d]在答案1最近距离是[%d],在答案2最近距离是[%d]", node.Value, distance, distance2)
		}
	}
	return nil
}

func generateGraph() *Graph {
	// {weight, from, to}
	return CreateGraph([][]int{
		{283, 1, 2},
		{196, 1, 3},
		{493, 1, 4},
		{504, 2, 1},
		{765, 2, 3},
		{177, 2, 4},
		{105, 3, 1},
		{774, 3, 2},
		{801, 3, 4},
		{903, 4, 1},
		{226, 4, 2},
		{210, 4, 3},
	})
}
