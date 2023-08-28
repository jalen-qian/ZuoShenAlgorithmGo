package class_13

import "ZuoShenAlgorithmGo/utils"

// 判断一颗二叉树是否是满二叉树

func IsFull(root *TreeNode) bool {
	if root == nil {
		return true
	}
	height, n := isFullProcess(root)
	// 如果n = 2 ^ height - 1，则是满二叉树
	return n == 1<<height-1
}

// 判断是否是满二叉树的递归过程，返回当前子树的高度和节点个数
func isFullProcess(x *TreeNode) (height int, n int) {
	// 是空树，则高度和节点个数都是0
	if x == nil {
		height, n = 0, 0
		return
	}
	// 递归获取左子树的高度和节点个数
	leftHeight, leftN := isFullProcess(x.Left)
	// 递归获取右子树的高度和节点个数
	rightHeight, rightN := isFullProcess(x.Right)
	// 计算当前树的高度和节点个数
	height = utils.Max(leftHeight, rightHeight) + 1
	n = leftN + rightN + 1
	return
}

// IsFull2 另一种方式，搜集左右子树是否是满二叉树，且高度是否相等
// 如果一棵树的左右子树都是满二叉树，且高度也相等，则这棵树也是满二叉树
func IsFull2(root *TreeNode) bool {
	isFull, _ := isFullProcess2(root)
	return isFull
}

// 判断是否是满二叉树的递归过程，返回当前子树的高度和节点个数
func isFullProcess2(x *TreeNode) (isFull bool, height int) {
	// 是空树，则算是满的，但是高度是0
	if x == nil {
		isFull, height = true, 0
		return
	}
	// 递归获取左子树是否满，以及高度
	isLeftFull, leftHeight := isFullProcess2(x.Left)
	// 递归获取右子树是否满，以及高度
	isRightFull, rightHeight := isFullProcess2(x.Right)
	isFull = true
	if !isLeftFull {
		isFull = false
	}
	if !isRightFull {
		isFull = false
	}
	if leftHeight != rightHeight {
		isFull = false
	}
	// 当前树的高度是左右子树最大高度+1
	height = utils.Max(leftHeight, rightHeight) + 1
	return
}
