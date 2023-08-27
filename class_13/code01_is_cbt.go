package class_13

// 判断一个二叉树是否是完全二叉树

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
		// 如果有右孩子没左孩子，则一定不是完全二叉树
		if cur.Left == nil && cur.Right != nil {
			return false
		}
		// 是叶子节点，则标记遇到了叶子节点
		if cur.Left == nil && cur.Right == nil {
			hasGotLeaf = true
		} else if hasGotLeaf { // 如果不是叶子节点，但是之前已经遇到过叶子节点，则不是完全二叉树
			return false
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
