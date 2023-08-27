package class_11_12

import "fmt"

// 二叉树折纸问题

// PrintAllFolds 折N次依次打印折痕顺序
func PrintAllFolds(n int) {
	// 从第1层开始打印（从根节点开始），而且总共有n，且第1层的折痕是凹折痕
	process(1, n, true)
	fmt.Println()
}

// 递归过程：想象一个如题的二叉树，当前是在这个二叉树的第i层的某个节点，总共有N层。
// 并且当前节点的折痕是 down == true 则是凹折痕，false 则为凸折痕
// 打印这个子树的中序遍历
func process(i int, n int, down bool) {
	// 如果当前层已经超过了总层数，则返回
	if i > n {
		return
	}
	// 先左再头再右，中序遍历。
	// 先打印左子树，我的左子树一定是凹折痕，且层数比我大1
	process(i+1, n, true)
	// 再打印我自己
	if down {
		fmt.Print("凹 ")
	} else {
		fmt.Print("凸 ")
	}
	// 最后打印右子树，我的右子树一定是凸折痕，且层数比我大1
	process(i+1, n, false)
}
