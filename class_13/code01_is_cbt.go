package class_13

import "ZuoShenAlgorithmGo/utils"

// 判断一个二叉树是否是完全二叉树
// 完全二叉树：一棵二叉树如果前面每层都是满的，最后一层就算不满，也是在从左往右依次变满的路上，则是完全二叉树

func IsCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 按层遍历
	queue := NewTreeNodeQueue()
	queue.Push(root)
	// 标记是否get到第一个叶子节点，初始是false
	var hasGotLeaf bool
	for !queue.IsEmpty() {
		cur := queue.Poll()
		// 如果有右孩子没左孩子，则一定不是完全二叉树，返回false
		if cur.Left == nil && cur.Right != nil {
			return false
		}
		// 已经遇到过不双全的节点了，当前节点却不是叶子节点，则返回false
		if hasGotLeaf && (cur.Left != nil || cur.Right != nil) {
			return false
		}
		// 遇到第一个不双全的节点，则标记“叶子节点”，表示“接下来的节点都只可能是叶子节点”
		if cur.Left == nil || cur.Right == nil {
			hasGotLeaf = true
		}
		if cur.Left != nil {
			queue.Push(cur.Left)
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
		}
	}
	// 完全遍历走完了，则返回true
	return true
}

func IsCBT1(root *TreeNode) bool {
	isCBT, _, _ := IsCBTProcess(root)
	return isCBT
}

// IsCBTProcess 递归过程，返回3个信息：
// 1：当前树是否是完全二叉树 2：当前树是否是满二叉树 3：当前树高度
func IsCBTProcess(x *TreeNode) (isCBT bool, isFull bool, height int) {
	// 如果是空树，则算完全二叉树，且高度为0
	if x == nil {
		return true, true, 0
	}
	// 先假定当前节点是完全二叉树
	isCBT = true
	// 递归获取左右孩子的数据
	leftIsCBT, leftIsFull, leftHeight := IsCBTProcess(x.Left)
	rightIsCBT, rightIsFull, rightHeight := IsCBTProcess(x.Right)
	// 满足以下5种情况，则当前树不是完全二叉树，否则就是
	// 1. 如果左右孩子有一个不是完全二叉树，则当前不是
	if !leftIsCBT || !rightIsCBT {
		isCBT = false
	}
	// 2. 如果左右孩子都不是满二叉树，则当前肯定不是
	if !leftIsFull && !rightIsFull {
		isCBT = false
	}
	// 3.如果左孩子是满二叉树，右孩子不是，但左右孩子高度不相同，则当前不是
	if leftIsFull && !rightIsFull && leftHeight != rightHeight {
		isCBT = false
	}
	// 4. 如果左孩子不是满二叉树，又孩子是满二叉树，但左孩子高度不比右孩子大1，则当前不是
	if !leftIsFull && rightIsFull && leftHeight != rightHeight+1 {
		isCBT = false
	}
	// 5. 如果左右孩子都是满二叉树，则要么左孩子高度等于右孩子高度，要么左孩子高度比右孩子高度大1
	if leftIsFull && rightIsFull && !(leftHeight == rightHeight || leftHeight == rightHeight+1) {
		isCBT = false
	}
	// 当前树是否是满二叉树：左右都是满二叉树，且左右高度相同，否则不是
	if leftIsFull && rightIsFull && leftHeight == rightHeight {
		isFull = true
	}
	// 当前树的高度：左右子树高度最大值+1
	height = utils.Max(leftHeight, rightHeight) + 1
	// 返回结果
	return
}

func IsCBT2(root *TreeNode) bool {
	isCBT, _, _ := IsCBTProcess2(root)
	return isCBT
}

// IsCBTProcess2 递归过程，返回3个信息：
// 1：当前树是否是完全二叉树 2：当前树是否是满二叉树 3：当前树高度
func IsCBTProcess2(x *TreeNode) (isCBT bool, isFull bool, height int) {
	// 如果是空树，则算完全二叉树，且高度为0
	if x == nil {
		return true, true, 0
	}
	// 递归获取左右孩子的数据
	leftIsCBT, leftIsFull, leftHeight := IsCBTProcess(x.Left)
	rightIsCBT, rightIsFull, rightHeight := IsCBTProcess(x.Right)
	// 当前树是否是满二叉树：左右子树都是满二叉树，且左右子树高度相等
	isFull = leftIsFull && rightIsFull && leftHeight == rightHeight
	// 默认isCBT为false，下面讨论哪些情况下isCBT可以是true
	// 如果当前树是满二叉树，则是完全二叉树
	if isFull {
		isCBT = true
	} else {
		// 如果当前不是满二叉树，要想当前树是完全二叉树，有两种情况
		// 1. 左子树是完全二叉树，右子树是满二叉树，且左子树高度比右子树高度大1
		if leftIsCBT && rightIsFull && leftHeight-rightHeight == 1 {
			isCBT = true
		}
		// 2. 左子树是满二叉树，右子树是完全二叉树，且左子树高度和右子树高度相等
		if leftIsFull && rightIsCBT && leftHeight == rightHeight {
			isCBT = true
		}
	}
	// 当前树的高度：左右子树高度最大值+1
	height = utils.Max(leftHeight, rightHeight) + 1
	// 返回结果
	return
}
