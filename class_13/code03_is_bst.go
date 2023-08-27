package class_13

import "ZuoShenAlgorithmGo/utils"

// IsBST 判断是否是搜索二叉树
func IsBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return processIsBST(root).isBST
}

type info struct {
	isBST bool
	min   int
	max   int
}

func processIsBST(x *TreeNode) *info {
	if x == nil {
		// 涉及值的比较，不能返回0，因为无法与值就是0，但是非空的节点区分，所以返回info结构，x==nil时，返回空
		return nil
	}
	leftInfo := processIsBST(x.Left)
	rightInfo := processIsBST(x.Right)

	// 最小值

	return isBst, max, min
}

func getMin(a, b, c int) int {
	return utils.Min(utils.Min(a, b), c)
}

func getMax(a, b, c int) int {
	return utils.Max(utils.Max(a, b), c)
}
