package class_13

import "ZuoShenAlgorithmGo/utils"

// 判断一棵二叉树是否是平衡二叉树

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 主函数只需要获取平衡，不获取高度
	isBalanced, _ := processIsBalanced(root)
	return isBalanced
}

// 递归函数，返回一个节点是否是平衡二叉树，以及高度
func processIsBalanced(x *TreeNode) (isBalanced bool, height int) {
	// 如果是空树，则高度是0，且是平衡二叉树
	if x == nil {
		isBalanced, height = true, 0
		return
	}
	// 不是空树，分别递归获取左右的高度以及是否平衡
	leftIsB, leftH := processIsBalanced(x.Left)
	rightIsB, rightH := processIsBalanced(x.Right)

	// 平衡条件：左右子树都平衡，且左右子树高度差绝对值小于2
	isBalanced = leftIsB && rightIsB && getDelAbs(leftH, rightH) < 2
	// 当前树的高度，左右子树中高度最大的+1
	height = utils.Max(leftH, rightH) + 1
	return
}

// 求相减后的绝对值
func getDelAbs(a, b int) int {
	del := a - b
	if del < 0 {
		del = -del
	}
	return del
}

//
