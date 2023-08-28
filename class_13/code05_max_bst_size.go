package class_13

func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return largestBstProcess(root).largestBstN
}

func largestBstProcess(x *TreeNode) *largestBstInfo {
	if x == nil {
		return nil
	}
	leftInfo := largestBstProcess(x.Left)
	rightInfo := largestBstProcess(x.Right)
	max := x.Val
	if leftInfo != nil {
		max = getMax(leftInfo.max, max)
	}
	if rightInfo != nil {
		max = getMax(rightInfo.max, max)
	}
	min := x.Val
	if leftInfo != nil {
		min = getMin(leftInfo.min, min)
	}
	if rightInfo != nil {
		min = getMin(rightInfo.min, min)
	}
	isBst := true // 先默认是搜索二叉树
	// 左子树不为空且左子树不是搜索二叉树，则不是
	if leftInfo != nil && !leftInfo.isBst {
		isBst = false
	}
	// 右子树不为空且右子树不是搜索二叉树，则不是
	if rightInfo != nil && !rightInfo.isBst {
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
	n := 1
	if leftInfo != nil {
		n += leftInfo.n
	}
	if rightInfo != nil {
		n += rightInfo.n
	}
	var largestBstN int
	// 如果当前子树就是BST，则最大的BST节点个数就是当前子树的节点个数
	if isBst {
		largestBstN = n
	} else {
		// 如果当前子树不是BST，则最大的BST节点个数是左右子树中的较大者
		if leftInfo != nil {
			largestBstN = getMax(largestBstN, leftInfo.largestBstN)
		}
		if rightInfo != nil {
			largestBstN = getMax(largestBstN, rightInfo.largestBstN)
		}
	}

	return &largestBstInfo{
		isBst:       isBst,
		max:         max,
		min:         min,
		largestBstN: largestBstN,
		n:           n,
	}
}

type largestBstInfo struct {
	isBst       bool
	max         int
	min         int
	largestBstN int // 子树的最大BST节点个数
	n           int // 子树本身的节点个数
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
