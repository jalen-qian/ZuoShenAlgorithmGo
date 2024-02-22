package class_16_17

import (
	"fmt"
)

// ↑ ↓ ——→ ←—— ↘︎↙︎ ↗︎
// 图的深度优先遍历
// 深度优先遍历通俗做法：一条路没走完就一直走到底，走完再往回走一步再重新走
// 深度优先遍历需要用栈和一个set
/*
        1  ——————→ 8
     ↙︎  ↓  ↘︎   ↙︎
   2 ——→ 3  ——→ 4
   ↘︎  ↗︎
     5
*/
// 比如上面的图，从1开始，1->2->5->3->4 ，4没路了，退到3，3没路，退到2，2有到3的路，但是3已经遍历过了，退到1，1通往3、4的路都遍历过了，遍历8
// 最终深度优先遍历的结果是：1 2 5 3 4 8
// 当然也可以是 1 2 3 4 5 8

func GraphDfs(from *Node) {
	if from == nil {
		return
	}
	// 使用切片实现一个栈
	stack := make([]*Node, 0)
	set := make(map[*Node]struct{})
	stack = append(stack, from)
	// 入栈就打印
	fmt.Print(from.Value, " ")
	set[from] = struct{}{}
	for len(stack) > 0 {
		// 先弹出
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 如果cur有路可以走，就将cur重新入栈，并将cur的某个邻居入栈
		for _, n := range cur.Nexts {
			if _, ok := set[n]; !ok {
				stack = append(stack, cur, n)
				// 入栈就打印
				fmt.Print(n.Value, " ")
				set[n] = struct{}{}
				// 找到一条路，就退出
				break
			}
		}
	}
	fmt.Println()
}
