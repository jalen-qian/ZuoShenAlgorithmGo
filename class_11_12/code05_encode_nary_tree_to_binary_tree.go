package class_11_12

// LeetCode测试链接: https://leetcode.cn/problems/encode-n-ary-tree-to-binary-tree/

// 这行不用提交

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}
	btRoot := &TreeNode{Val: root.Val}
	// 将root的所有子节点，挂在btRoot的左孩子右边界上（递归）
	btRoot.Left = c.ec(root.Children)
	return btRoot
}

// 将所有孩子挂在btRoot的左孩子右边界上
func (c *Codec) ec(children []*Node) *TreeNode {
	// 这些孩子一定有一个头，是要成为root的左孩子的，用head表示
	var head *TreeNode
	var cur *TreeNode
	for _, child := range children {
		// 遇到一个孩子，先建出来
		childBT := &TreeNode{Val: child.Val}
		if head == nil {
			head = childBT
		} else {
			cur.Right = childBT
		}
		cur = childBT
		// 深度优先遍历，当前循环结束之前，递归搞定当前child的孩子，挂在当前节点左树的右边界
		cur.Left = c.ec(child.Children)
	}
	return head
}

func (c *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}
	// 二叉树的根节点，也一定是多叉树的根节点，且多叉树的所有孩子，一定在二叉树的左孩子上挂着
	// dc就是将左孩子所有右边界遍历出来，放到一个切片中返回，并且在此过程中使用递归
	nRoot := &Node{Val: root.Val, Children: c.dc(root.Left)}
	return nRoot
}

func (c *Codec) dc(left *TreeNode) []*Node {
	var children []*Node
	cur := left
	for cur != nil {
		n := &Node{Val: cur.Val, Children: c.dc(cur.Left)}
		children = append(children, n)
		// 不断往右孩子走
		cur = cur.Right
	}
	return children
}
