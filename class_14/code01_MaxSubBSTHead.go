package class_14

import (
	"ZuoShenAlgorithmGo/utils"
)

// 题目：给定一个二叉树的头结点，返回这颗二叉树最大二叉搜索树子树的头结点，如果不存在，则返回null
// 二叉搜索树（BST Binary Search Tree)的概念：
// 1. 所有节点，其左子树最大值 < 当前根节点
// 2. 所有节点，其右子树最小值 > 当前根节点
// 3. 左右子树都是二叉搜索树

// MaxSubBSTHead1 思路1：暴力方法，给定一个X节点，判断当前X节点是否是BST，如果是，则返回X
// 如果不是，则递归X的左右子树，返回左右子树中最大的
func MaxSubBSTHead1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 判断当前x是否是二叉搜索树，判断方法是中序遍历。
	// 之前的笔记中有总结：中序遍历后的顺序是升序且没重复值，是这棵树是BST的充分必要条件
	// 我们只需要设计一个函数，传入node，如果node是BST，则返回node的大小，如果不是或者node为空，都返回0
	// 1. 如果当前root已经是BST了，则返回root
	if getBSTSize(root) != 0 {
		return root
	}
	// 当前root不是BST，则递归左右子树，判断哪个子树的最大BST大，则返回哪个，一样大则选左子树
	leftNode := MaxSubBSTHead1(root.Left)
	rightNode := MaxSubBSTHead1(root.Right)
	if getBSTSize(leftNode) >= getBSTSize(rightNode) {
		return leftNode
	} else {
		return rightNode
	}
}

// 返回二叉搜索树的大小，如果是BST，则返回大小，如果不是或者为空，都返回0
func getBSTSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 获取中序遍历
	var list []*TreeNode
	in(root, &list)
	// 判断是否是严格升序
	for i := 0; i < len(list)-1; i++ {
		// 不是升序，则当前root一定不是BST，直接返回0
		if list[i].Val >= list[i+1].Val {
			return 0
		}
	}
	// 走到这里，是BST，返回list大小
	return len(list)
}

func in(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	}
	in(root.Left, list)
	*list = append(*list, root)
	in(root.Right, list)
}

// =====================

// 解法二：使用二叉树的递归套路
// 我们假定得到了一个二叉树的节点X，讨论需要左右子树提供什么信息，才能获取X的最大BST子树节点？
// 1. 需要知道X本身是否是二叉搜索树BST，如果是，则直接返回X即可，要知道这个，需要知道：
//    a. 左右子树是不是BST
//    b. 左子树的最大值，右子树的最小值
//    ps: 虽然左子树只需要最大值，右子树只需要最小值，但是我们都要就行了，为了实现逻辑的统一。
// 2. 如果判断了X本身不是BST，则需要知道左右子树各自的最大BST子树节点，以及大小，返回较大的那个。
//
// 综上，我们需要在递归中，左右子树提供给当前X节点的信息有：
// 1. 最大BST子树的节点
// 2. 最大BST子树的size（节点个数）
// 3. 左右子树自己的最大值
// 4. 左右子树自己的最小值

func MaxSubBSTHead(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 主函数，前面已经判空了，所以这里获取到的info一定不是空指针
	return maxSubBSTHeadProcess(root).maxSubBST
}

type info struct {
	maxSubBST     *TreeNode // 最大BST子树头节点
	maxSubBSTSize int       // 最大BST子树的大小
	maxVal        int       // 整棵树最大值
	minVal        int       // 整棵树最小值
}

// 递归x节点拿信息
func maxSubBSTHeadProcess(x *TreeNode) *info {
	// 如果x节点本身是空的，则返回nil
	if x == nil {
		return nil
	}
	leftInfo := maxSubBSTHeadProcess(x.Left)
	rightInfo := maxSubBSTHeadProcess(x.Right)
	// 最大最小值，初始都赋值为当前节点的值，再根据左右子树的值判断
	maxVal, minVal := x.Val, x.Val
	var maxSubBSTRoot *TreeNode
	maxSubBSTSize := 0
	if leftInfo != nil {
		maxVal = utils.Max(maxVal, leftInfo.maxVal)
		minVal = utils.Min(minVal, leftInfo.minVal)
		maxSubBSTRoot = leftInfo.maxSubBST
		maxSubBSTSize = leftInfo.maxSubBSTSize
	}
	if rightInfo != nil {
		maxVal = utils.Max(maxVal, rightInfo.maxVal)
		minVal = utils.Min(minVal, rightInfo.minVal)
		// 如果右子树的 maxSubBSTSize 比左子树大，则赋值为右子树的
		if rightInfo.maxSubBSTSize > maxSubBSTSize {
			maxSubBSTRoot = rightInfo.maxSubBST
			maxSubBSTSize = rightInfo.maxSubBSTSize
		}
	}
	// 当前x是二叉搜索树条件：
	// 左子树是BST 且 右子树是BST 且左子树的最大值 < x.Val < 右子树的最小值
	// ps:左子树是空，也算BST（空树是算BST）
	if (leftInfo == nil || (leftInfo.maxSubBST == x.Left && leftInfo.maxVal < x.Val)) &&
		(rightInfo == nil || (rightInfo.maxSubBST == x.Right && rightInfo.minVal > x.Val)) {
		// 当前x是BST，则最大的BST子树是x本身
		maxSubBSTRoot = x
		maxSubBSTSize = 1 // 当前x占1个节点
		// 左右子树不是空树的情况下，累加上左右子树的size
		if leftInfo != nil {
			maxSubBSTSize += leftInfo.maxSubBSTSize
		}
		if rightInfo != nil {
			maxSubBSTSize += rightInfo.maxSubBSTSize
		}
	}
	return &info{
		maxSubBST:     maxSubBSTRoot,
		maxSubBSTSize: maxSubBSTSize,
		maxVal:        maxVal,
		minVal:        minVal,
	}
}
