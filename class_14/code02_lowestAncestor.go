package class_14

import "ZuoShenAlgorithmGo/utils"

// 二叉树递归套路题：给你一个二叉树的根节点root和 a，b 两个节点，请你返回a b 的最低公共祖先。
// 最低公共祖先表示 从 a 和 b 不断往父节点走，汇聚到的第一个节点。
// 比如:
//        d                                  a
//      /   \                              /   \
//     a     c                            f     b
//          /  \
//         b    m
// 这棵树的最低公共祖先是d               这颗树的最低公共祖先是a
// 可以知道：1.最低公共祖先可能是a或者b本身；2.如果a和b都在树中出现，则一定有最低公共祖先

// 暴力解法思路：
// 1. 使用一个哈希表，遍历整棵树，并存储每个节点的父亲节点
// 2. 使用一个集合，将a节点的所有祖先存储在集合中
// 3. 从哈希表不断获取b节点的祖先，一旦发现在a祖先集合中，则返回，没发现则返回空

// 暴力方法
func lowestAncestor1(root *TreeNode, a, b *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 父亲节点map， key 的 父亲是 value
	parentMap := make(map[*TreeNode]*TreeNode)
	// 根节点的父亲是空
	parentMap[root] = nil
	// 遍历并得到所有节点的父亲
	fillParentMap(root, parentMap)
	// 获取a节点所有的祖先，并存储到set中
	aParentSet := utils.NewHashSet[*TreeNode]()
	cur := a
	for cur != nil {
		aParentSet.Add(cur)
		cur = parentMap[cur]
	}
	// 获取b的所有祖先（包括b)，如果在a祖先集合中，则返回
	cur = b
	for cur != nil {
		if aParentSet.Contains(cur) {
			return cur
		}
		cur = parentMap[cur]
	}
	return nil
}

func fillParentMap(root *TreeNode, parentMap map[*TreeNode]*TreeNode) {
	if root.Left != nil {
		parentMap[root.Left] = root
		fillParentMap(root.Left, parentMap)
	}
	if root.Right != nil {
		parentMap[root.Right] = root
		fillParentMap(root.Right, parentMap)
	}
}

// =========== 二叉树递归套路方法 ===============
// 假设我们遍历到了某个x节点，现在要知道x这颗树上a、b的最低公共祖先，可以分情况讨论，然后分析需要在递归时向左右子树要什么信息：
// 1. 如果a和b都在左子树，则a、b的最低公共祖先一定在左子树，此时x子树中a、b的最低公共祖先，就是左子树中的这个（a、b都在右子树同理）
// 2. 如果a在左子树，b在右子树，则a、b的最低公共祖先一定是当前的x节点（a在右子树，b在左子树同理）
// 3. 如果a和b有1个既不在左子树，也不在右子树，则当前x树一定不存在最低公共祖先。
// 综上，我们只需要左右子树返回以下3个信息：
// 1. 是否发现了a节点
// 2. 是否发现了b节点
// 3. 是否发现了a、b的最低公共祖先
// 代码实现如下：

func lowestAncestor2(root *TreeNode, a, b *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	_, _, lowestAncestor := processLowestAncestor2(root, a, b)
	return lowestAncestor
}

// 递归过程中，返回3个信息
// foundA: 是否找到了A
// foundB: 是否找到了B
// lowestAncestor 最低公共祖先
func processLowestAncestor2(x *TreeNode, a, b *TreeNode) (foundA, foundB bool, lowestAncestor *TreeNode) {
	if x == nil {
		return
	}
	// 递归左子树和右子树，获取信息
	leftFoundA, leftFoundB, leftLowestAncestor := processLowestAncestor2(x.Left, a, b)
	rightFoundA, rightFoundB, rightLowestAncestor := processLowestAncestor2(x.Right, a, b)
	// 当前树是否有a节点：左右子树有a,或者当前就是a
	foundA = leftFoundA || rightFoundA || x == a
	// 当前树是否有b节点：左右子树有b,或者当前就是b
	foundB = leftFoundB || rightFoundB || x == b
	// 当前树中 a b 的最低公共祖先
	// 1. 如果左右子树已经找找到了，则返回左右子树的
	if leftLowestAncestor != nil {
		lowestAncestor = leftLowestAncestor
	} else if rightLowestAncestor != nil {
		lowestAncestor = rightLowestAncestor
	} else {
		// 2. 左右子树没找到，则判断是否当前x节点是最低公共祖先
		if foundA && foundB {
			lowestAncestor = x
		}
	}
	return
}
