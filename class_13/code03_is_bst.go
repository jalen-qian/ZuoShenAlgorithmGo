package class_13

import "ZuoShenAlgorithmGo/utils"

// IsBST 判断是否是搜索二叉树
// 搜索二叉树概念：一颗二叉树的任意子树，其左子树都比头节点小，右子树都比头结点大，则这颗树是搜索二叉树。
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
	max := x.Val
	if leftInfo != nil {
		max = utils.Max(leftInfo.max, max)
	}
	if rightInfo != nil {
		max = utils.Max(rightInfo.max, max)
	}
	min := x.Val
	if leftInfo != nil {
		min = utils.Min(leftInfo.min, min)
	}
	if rightInfo != nil {
		min = utils.Min(rightInfo.min, min)
	}
	isBst := true // 先默认是搜索二叉树
	// 左子树不为空且左子树不是搜索二叉树，则不是
	if leftInfo != nil && !leftInfo.isBST {
		isBst = false
	}
	// 右子树不为空且右子树不是搜索二叉树，则不是
	if rightInfo != nil && !rightInfo.isBST {
		isBst = false
	}
	// 左子树的最大值不比当前值小，则不是
	if leftInfo != nil && leftInfo.max >= x.Val {
		isBst = false
	}
	// 又子树的最小值不比当前值大，则不是
	if rightInfo != nil && rightInfo.min <= x.Val {
		isBst = false
	}
	return &info{isBST: isBst, max: max, min: min}
}

// 判断是否是搜索二叉树的方式2，中序遍历是否升序
// 这是充分必要条件：中序遍历升序，则一定是搜索二叉树（这里指没有重复值的升序）
// 如果是搜索二叉树，则中序遍历一定是升序

// IsBST2 通过中序遍历是否升序的方式判断是否是搜索二叉树
func IsBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	inList := make([]int, 0)
	// 中序遍历
	in(root, &inList)
	// 判断inList是否是升序，如果不是，则不是搜索二叉树
	for i := 1; i < len(inList); i++ {
		if inList[i] <= inList[i-1] {
			return false
		}
	}
	return true
}

// 递归实现中序遍历，并将顺序记录到list中
func in(x *TreeNode, list *[]int) {
	if x == nil {
		return
	}
	in(x.Left, list)
	*list = append(*list, x.Val)
	in(x.Right, list)
}
