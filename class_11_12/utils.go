package class_11_12

import (
	"ZuoShenAlgorithmGo/class_03"
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getTreeHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + utils.Max(getTreeHeight(root.Left), getTreeHeight(root.Left))
}

func writeArray(root *Node, row, column, treeHeight int, resArray [][]string) {
	if root == nil {
		return
	}
	resArray[row][column] = strconv.Itoa(root.Value)
	currentHeight := (row + 1) / 2 // 当前高度
	if currentHeight == treeHeight {
		return // 下面没有子节点了
	}
	gap := treeHeight - currentHeight - 1 // 到左/右儿子的距离
	// 填充左儿子
	if root.Left != nil {
		// 先写树结构符号
		resArray[row+1][column-gap] = "/"
		// 再写左儿子
		writeArray(root.Left, row+2, column-gap*2, treeHeight, resArray)
	}
	// 填充右儿子
	if root.Right != nil {
		resArray[row+1][column+gap] = "\\"
		writeArray(root.Right, row+2, column+gap*2, treeHeight, resArray)
	}
}

func PrintBT(root *Node) {
	if root == nil {
		fmt.Println("空树！")
		return
	}
	height := getTreeHeight(root)
	fmt.Printf("height: %v\n", height)
	// 总宽度为节点高度 * 2 - 1, 因为还要画树枝符号
	totalHeight := height*2 - 1
	// 最大宽度为3 * 2^(n - 1) + 1，公式如下：
	/**
	   父亲节点占1, 两个孩子空间各占1, 连接线各占1, 每个父子关系共占5, 每个关系之间空1, 最左最右各空1
	  第2行： 5 + 2 （1个父子结构占位+左右两个空格分割）
	  第3行：2 * 5 + (1 + 2) （2个父子结构占位+1个相邻父子结构间空格+左右两个空格分割）
	  第4行：4 * 5 + (3 + 2) （4个父子结构占位+3个相邻父子结构间空格+左右两个空格分割）
	  第5行：8 * 5 + (7 + 2)
	  第n行：5 * 2 ^ (n - 2) + (2 ^ (n - 2) - 1) + 2 = 6 * 2 ^ (n-2) + 1 = 3 * 2 ^ (n - 1) + 1
	*/
	var totalWidth int
	if height == 0 {
		totalWidth = 1
	} else {
		totalWidth = (2<<(height-2))*3 + 1
	}

	// 创建数组
	printArray := make([][]string, totalHeight)
	for i := range printArray {
		printArray[i] = make([]string, totalWidth)
		for j := range printArray[i] {
			printArray[i][j] = " "
		}
	}

	// 计算打印数组
	writeArray(root, 0, totalWidth/2, height, printArray)

	// 打印
	for i := range printArray {
		var res string
		for j := range printArray[i] {
			res = res + printArray[i][j]
		}
		fmt.Println(res)
	}
}

var NewNodeQueue = class_03.NewMyQueue[*Node]

// 生成一个随机的二叉树
// minValue 最小值 maxValue 最大值 maxLevel 最大的层
func generateRandomBT(minValue int, maxValue int, maxLevel int) *Node {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 从第1层开始递归创建，直到层数 >= maxLevel
	return generateBT(r, 1, minValue, maxValue, maxLevel)
}

func generateBT(r *rand.Rand, curLevel int, minValue int, maxValue int, maxLevel int) *Node {
	// 40%概率会生成空节点
	percent := r.Float64()
	if curLevel > maxLevel || percent < 0.4 {
		return nil
	}
	// 否则，生成一个随机的节点
	node := &Node{
		Value: r.Intn(maxValue-minValue+1) + minValue, // 5,10  [0,6)即[0,5]+5 => [5,10]
	}
	// 下一层随机生成
	node.Left = generateBT(r, curLevel+1, minValue, maxValue, maxLevel)
	node.Right = generateBT(r, curLevel+1, minValue, maxValue, maxLevel)
	return node
}

// isBTEqual 判断两颗二叉树是否完全相等
func isBTEqual(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil && root1 != nil {
		return false
	}
	if root1 == nil && root2 == nil {
		return true
	}
	// 走到这里，说明都不为空
	if root1.Value != root2.Value {
		return false
	}
	// 左右子树也要完全相等
	return isBTEqual(root1.Left, root2.Left) && isBTEqual(root1.Right, root2.Right)
}
