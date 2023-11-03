package class_14

import (
	"math/rand"
	"testing"
	"time"
)

func TestLowestAncestor(t *testing.T) {
	// 使用对数器测试
	t.Log("测试开始...")
	testTimes := 100000
	minValue := -1000
	maxValue := 1000
	maxLevel := 10
	for i := 0; i < testTimes; i++ {
		// 生成随机的树，a b 节点
		root, a, b := generateRandomBTAndAB(minValue, maxValue, maxLevel)
		if lowestAncestor1(root, a, b) != lowestAncestor2(root, a, b) {
			t.Errorf("出错了！\n lowestAncestor1:%v \n lowestAncestor2:%v", lowestAncestor1(root, a, b), lowestAncestor2(root, a, b))
			return
		}
	}
	t.Log("测试成功！")
}

// generateRandomBTAndAB 随机生成一个二叉树，返回根节点，并且随机返回二叉树中的两个节点a 和 b
// ps：a 和 b 可能不在树中，也可能是nil
func generateRandomBTAndAB(minVal, maxVal, maxLevel int) (root, a, b *TreeNode) {
	// 先生成一颗随机树
	root = generateRandomBT(minVal, maxVal, maxLevel)
	// 遍历这颗树，拿到所有节点的集合
	totalNodes := make([]*TreeNode, 0)
	getAllNodes(root, &totalNodes)
	// 如果是空树，则返回
	n := len(totalNodes)
	if n == 0 {
		return root, a, b
	}
	// 随机返回两个节点
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 随机生成数组中的一个位置，作为a节点
	indexA := myRand.Intn(n) // n=10, [0,9]
	// 随机生成数组中的一个位置，作为b节点，如果生成的位置和a相同，则返回a后面的一个
	indexB := myRand.Intn(n)
	if indexB == indexA {
		// 如果只有1个，则A是这个唯一的节点，B是空
		if n == 1 {
			indexB = -1
		} else {
			// 至少有两个时，如果B与A碰撞了，则B挪到A下一个，如果A是最后1个，则B挪到0
			indexB = indexA + 1
			if indexB > n {
				indexB = 0
			}
		}
	}
	if indexA >= 0 && indexA < n {
		a = totalNodes[indexA]
	}
	if indexB >= 0 && indexB < n {
		b = totalNodes[indexB]
	}
	// a b 一定在树中，现在随机将a和b替换为空或者一个在树中不存在的节点
	randomReplace(a, myRand)
	randomReplace(b, myRand)
	return
}

func randomReplace(x *TreeNode, myRand *rand.Rand) {
	// 30%的概率才替换
	if myRand.Float32() < 0.3 {
		// 替换的情况下，50%几率替换为空，50%几率替换为一个任意的节点
		if myRand.Float32() < 0.5 {
			x = nil
		} else {
			x = &TreeNode{Val: myRand.Intn(100)}
		}
	}
}

func getAllNodes(x *TreeNode, nodeList *[]*TreeNode) {
	if x == nil {
		return
	}
	*nodeList = append(*nodeList, x)
	getAllNodes(x.Left, nodeList)
	getAllNodes(x.Right, nodeList)
}
