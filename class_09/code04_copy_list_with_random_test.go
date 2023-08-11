package class_09

import (
	"ZuoShenAlgorithmGo/utils"
	"math/rand"
	"testing"
	"time"
)

// 测试拷贝带有随机指针的单链表
func TestCopyRandomList(t *testing.T) {
	t.Log("Start testing...")
	for i := 0; i < 100000; i++ {
		head := generateListWithRand(1000, 0, 1000)
		newHead := CopyRandomList1(head)
		newHead2 := CopyRandomList2(newHead)
		if !checkCopyRight(head, newHead) || !checkCopyRight(newHead2, newHead2) {
			t.Errorf("Fucking fucked!! 拷贝的与原始链表不一致！！")
			return
		}
	}
	t.Log("Great!!!")
}

// 随机生成一个带Rand指针的单链表
func generateListWithRand(maxLen int, minNum, maxNum int) *NodeWithRandom {
	// 先生成一个随机数组
	randomArr := utils.GenerateRandomSlice(maxLen, minNum, maxNum)
	if len(randomArr) == 0 {
		return nil
	}
	// 根据这个随机数组，生成对应的链表节点数组
	nodeArr := make([]*NodeWithRandom, len(randomArr))
	// 创建这个随机数组
	for i, v := range randomArr {
		nodeArr[i] = &NodeWithRandom{Value: v}
	}
	// 将数组中的节点的Next节点一个个串起来
	for i := 0; i < len(nodeArr)-1; i++ {
		// 从0串到n-2
		nodeArr[i].Next = nodeArr[i+1]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 遍历，随机指rand指针
	for i := 0; i < len(nodeArr); i++ {
		percent := r.Float64()
		// 40% 的概率指向null
		if percent < 0.4 {
			nodeArr[i].Rand = nil
		} else {
			// 60%的概率，随机指向任意一个节点
			loc := rand.Intn(len(nodeArr)) // [0,N-1]
			nodeArr[i].Rand = nodeArr[loc]
		}
	}
	// 第1个节点就是头
	return nodeArr[0]
}
