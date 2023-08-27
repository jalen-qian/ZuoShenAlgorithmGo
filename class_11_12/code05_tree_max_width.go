package class_11_12

import "ZuoShenAlgorithmGo/utils"

// 给定一颗二叉树，返回最大的宽度

// MaxWidth 不使用容器实现
func MaxWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 生成一个TreeNode类型的队列
	queue := NewTreeNodeQueue()
	var curEnd = root     // 当前行的结束
	var nextEnd *TreeNode // 下一行的结束
	var count int         // 统计当前行的宽度
	var max int           // 整个树最大的宽度
	// 先将头节点入队
	queue.Push(root)
	// 队列不为空，则一直进行
	for !queue.IsEmpty() {
		// 弹出一个节点，就统计当前行宽度+1
		cur := queue.Poll()
		count++
		// 只要有左右子树，则一定是下一行的，下一行结束节点先记住
		if cur.Left != nil {
			queue.Push(cur.Left)
			nextEnd = cur.Left
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
			nextEnd = cur.Right
		}
		// 判断当前行是否现在结束了，如果当前行结束了，则统计max，并重置count,curEnd,nextEnd 3个变量
		if cur == curEnd {
			max = utils.Max(max, count)
			count = 0
			curEnd = nextEnd
			nextEnd = nil
		}
	}
	return max
}

// MaxWidthWithMap 使用容器实现
func MaxWidthWithMap(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 生成一个TreeNode类型的队列
	queue := NewTreeNodeQueue()
	// 生成一个map记录每个节点在第几层
	nodeLevelMap := make(map[*TreeNode]int)
	curLevel := 1 // 记录当前在第几层
	count := 0    // 统计当前层的个数
	max := 0      // 统计整棵树最大的宽度
	// 先将头节点入队，同时标记root在第1层
	queue.Push(root)
	nodeLevelMap[root] = 1

	// 队列不为空，则一直进行
	for !queue.IsEmpty() {
		// 弹出一个节点
		cur := queue.Poll()
		// 获取这个节点在第几层
		curNodeLevel := nodeLevelMap[cur]
		// 如果还是当前层，则当前层的数量++
		if curNodeLevel == curLevel {
			count++
		} else {
			// 如果不是当前层了，则curLevel跳到下一层
			curLevel = curNodeLevel
			count = 1
		}
		max = utils.Max(max, count)
		// 只要有左右子树，则一定是下一行的，入队同时将层数记录下来
		if cur.Left != nil {
			queue.Push(cur.Left)
			nodeLevelMap[cur.Left] = curLevel + 1
		}
		if cur.Right != nil {
			queue.Push(cur.Right)
			nodeLevelMap[cur.Right] = curLevel + 1
		}
	}
	return max
}
