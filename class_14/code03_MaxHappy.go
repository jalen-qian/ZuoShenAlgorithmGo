package class_14

import "ZuoShenAlgorithmGo/utils"

// 二叉树的递归套路深度实践：派对的最大快乐值
// 某公司的组织架构可以看成一个标准的，没有环的多叉树，员工信息定义如下：
//
// type Employee struct {
// 	 Happy int // 这名员工如果参加派对，可以带来的快乐值
// 	 Subordinates []*Employee // 这名员工的直接下级
// }
// 公司的每个员工都符合上述Employee的描述。树的头结点是公司的唯一老板。除了老板外的每个员工都有唯一的直接上级。
// 叶子节点是没有任何下级的基层员工（Subordinates为空），除了基层员工以外，每个员工都有一个或者多个直接下级。
//
// 派对的最大快乐值
// 这个公司要办party,你可以决定哪些员工来，哪些员工不来，规则：
// 1. 如果某个员工来了，那么这个员工的所有直接下级都不能来
// 2. 派对的整体快乐值是所有到场员工快乐值的累加
// 3. 你的目标是让派对的整体快乐值尽量大
// 给定一颗多叉树的头结点boss，请返回派对的最大快乐值。

type Employee struct {
	Happy        int         // 这名员工如果参加派对，可以带来的快乐值
	Subordinates []*Employee // 这名员工的直接下级
}

// 思路1:暴力递归，分情况讨论：
// 分析：对于任意一个员工x，分析这个员工贡献的快乐总值（这个员工就是这颗子树，下属所有员工贡献的快乐值累加起来，就是这颗子树的快乐总值）
// 可以分为两种情况，x的上级来和不来
// 1. x的上级来:
//    此时x不能来，x贡献的快乐值是所有x的下级在x不来的情况下贡献的快乐值
// 2. x的上级不来，此时x可来可不来
//    2.1 x自己来，贡献值是 x.Happy + x的所有下级在x来的情况下的快乐值，我们记作P1
//    2.2 x自己不来，贡献值是 x的所有下级在x不来的情况下的快乐值，我们记作P2
//    则：x这颗树贡献的快乐总值是p1 p2中的较大者。
//
// 综上，我们设计一个递归函数, 函数签名为：
//  func maxHappyProcess1(cur *Employee, superiorIsComing bool) int
// 其中 cur 是当前员工，superiorIsComing 是cur的上级是否来，这个递归函数将会返回这种情况下最大的快乐值。

// MaxHappy1 解法1
func MaxHappy1(boss *Employee) int {
	// boss 都是空的，一个人都没有，快乐值为0
	if boss == nil {
		return 0
	}
	// boss没有上级，所以“上级是否来”一定是false
	return maxHappyProcess1(boss, false)
}

func maxHappyProcess1(cur *Employee, superiorIsComing bool) int {
	// 1. 如果cur的上级来，则cur一定不能来，则cur的快乐值是在cur不来的情况下所有下级的累加
	if superiorIsComing {
		ans := 0 // 不能来，所以ans是0
		for _, subordinate := range cur.Subordinates {
			ans += maxHappyProcess1(subordinate, false)
		}
		return ans
	} else {
		// 2.如果cur的上级不来，则cur可以来可以不来
		// 2.1 cur来的情况，标记为ans1，因为来，所以ans1的初始值是cur的快乐值
		ans1 := cur.Happy
		// 2.2 cur不来，标记为ans2，因为不来，所以ans2的初始值是0
		ans2 := 0
		for _, subordinate := range cur.Subordinates {
			ans1 += maxHappyProcess1(subordinate, true)
			ans2 += maxHappyProcess1(subordinate, false)
		}
		// 最后的答案是 ans1 和 ans2 中的较大者
		return utils.Max(ans1, ans2)
	}
}

// ================ 使用二叉树递归套路 ====================
// 解法二：使用我们递归套路的思想，虽然这不是一个二叉树，而是多叉树，但是思想同样可以用
// 我们同样分类讨论，对于任意一个员工x：
// 1. 如果x来，则x的所有下级都不能来，此时x的快乐总值是 x.Happy + x的所有下级在不来情况下的快乐值
// 2. 如果x不来，则x的所有下级可以来也可以不来，x贡献的快乐值是 x下级在来和不来情况下贡献快乐值的较大者
// 所以综上：我们可以归纳出，需要从下级拿哪些信息：
// 1. 下级在自己来时的最大快乐值
// 2. 下级在自己不来时的最大快乐值

type happyInfo struct {
	yes int // 来的情况下的快乐值
	no  int // 不来的情况下的快乐值
}

// MaxHappy2 解法2，使用二叉树递归套路
func MaxHappy2(boss *Employee) int {
	// boss 都是空的，一个人都没有，快乐值为0
	if boss == nil {
		return 0
	}
	// boss没有上级，所以“上级是否来”一定是false
	bossInfo := maxHappyProcess2(boss)
	return utils.Max(bossInfo.yes, bossInfo.no)
}

func maxHappyProcess2(x *Employee) happyInfo {
	yes := x.Happy // 如果来，至少有x本身的快乐值
	no := 0        // 如果不来，则初始快乐值是0，要累加下属员工的快乐值
	// 遍历员工，并计算yes和no
	for _, subordinate := range x.Subordinates {
		subInfo := maxHappyProcess2(subordinate)
		// x来的情况下，下属都不能来，所以累加上下属不来的快乐值
		yes += subInfo.no
		// x不来的情况下，下属可以来可以不来，所以累加上下属这两种情况的较大值
		no += utils.Max(subInfo.yes, subInfo.no)
	}
	return happyInfo{
		yes: yes,
		no:  no,
	}
}
