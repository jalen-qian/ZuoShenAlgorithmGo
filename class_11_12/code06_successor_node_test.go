package class_11_12

import (
	"ZuoShenAlgorithmGo/class_03"
	"math/rand"
	"testing"
	"time"
)

// 测试找后继节点

func TestGetSuccessorNode(t *testing.T) {
	t.Log("测试开始...")
	for i := 0; i < 500000; i++ {
		// 生成一个随机的带父指针的二叉树
		root := generateRandomBTWithParent(0, 100, 20)
		// 获取这颗二叉树的任意一个节点
		randomNode := getRandomTreeNodeP(root)
		// 分别用两种方式获取后继节点，看是否相同
		successor1 := GetSuccessorNode(randomNode)
		successor2 := GetSuccessorNodeNormal(randomNode)
		if successor1 != successor2 {
			t.Errorf("测试失败:\n successor1:%v\n successor2:%v\n", successor1, successor2)
			return
		}
	}
	t.Log("测试成功！！！")
}

// 生成一个随机的有父指针的二叉树
// minValue 最小值 maxValue 最大值 maxLevel 最大的层
func generateRandomBTWithParent(minValue int, maxValue int, maxLevel int) *TreeNodeP {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 从第1层开始递归创建，直到层数 >= maxLevel
	return generateTreeNodeP(r, nil, 1, minValue, maxValue, maxLevel)
}

// 从一颗二叉树中随机挑选节点一个返回
func getRandomTreeNodeP(root *TreeNodeP) *TreeNodeP {
	// 中序遍历
	inQueue := class_03.NewMyQueue[*TreeNodeP]()
	inTreeNodeP(root, inQueue)
	if inQueue.IsEmpty() {
		return nil
	}
	// 随机从队列中弹出一个
	var arr []*TreeNodeP
	for !inQueue.IsEmpty() {
		arr = append(arr, inQueue.Poll())
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(arr))
	return arr[index]
}

func generateTreeNodeP(r *rand.Rand, parent *TreeNodeP, curLevel int, minValue int, maxValue int, maxLevel int) *TreeNodeP {
	// 40%概率会生成空节点
	percent := r.Float64()
	if curLevel > maxLevel || percent < 0.4 {
		return nil
	}
	// 否则，生成一个随机的节点
	node := &TreeNodeP{
		Val:    r.Intn(maxValue-minValue+1) + minValue, // 5,10  [0,6)即[0,5]+5 => [5,10]
		Parent: parent,
	}
	// 下一层随机生成
	node.Left = generateTreeNodeP(r, node, curLevel+1, minValue, maxValue, maxLevel)
	node.Right = generateTreeNodeP(r, node, curLevel+1, minValue, maxValue, maxLevel)
	return node
}
