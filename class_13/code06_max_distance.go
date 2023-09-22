package class_13

import "ZuoShenAlgorithmGo/utils"

// 求二叉树的最大距离

func MaxDistance(root *TreeNode) int {
	maxDistance, _ := distanceProcess(root)
	return maxDistance
}

// 递归，返回子树的最大距离和高度
func distanceProcess(x *TreeNode) (maxDistance int, height int) {
	// 如果是空树，则最大距离和高度都是0
	if x == nil {
		return
	}
	// 递归获取左子树最大距离与高度
	leftMaxDistance, leftHeight := distanceProcess(x.Left)
	// 递归获取右子树最大距离与高度
	rightMaxDistance, rightHeight := distanceProcess(x.Right)
	// 如果经过x,最大距离是左右高度相加再加1
	maxDistance = leftHeight + rightHeight + 1
	// 如果不经过x,最大距离是左右子树最大距离的较大者
	maxDistance = utils.Max(leftMaxDistance, maxDistance)
	maxDistance = utils.Max(rightMaxDistance, maxDistance)
	// 当前树最大高度 = 左右子树高度最大值 + 1
	height = utils.Max(leftHeight, rightHeight) + 1
	// 返回当前x的最大距离和高度
	return
}
