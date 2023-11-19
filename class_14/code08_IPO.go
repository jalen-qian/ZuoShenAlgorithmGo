package class_14

import (
	"ZuoShenAlgorithmGo/class_07"
	"ZuoShenAlgorithmGo/utils"
)

/**
贪心
输入: 正数数组costs、正数数组profits、正数K、正数M
costs[i]表示i号项目的花费
profits[i]表示i号项目在扣除花费之后还能挣到的钱(利润)
K表示你只能串行的最多做k个项目
M表示你初始的资金
说明: 每做完一个项目，马上获得的收益，可以支持你去做下一个项目。不能并行的做项目。
输出：你最后获得的最大钱数。

比如 项目列表：[{1,2}, {3,1}, {5,3}, {3,2}] 初始资金M=1, k=3
一开始，只能选择 {1,2} 这个项目来做，由于2表示**利润（挣到的钱）**，做完这个项目后，不但1元的花费挣回来，还额外挣了2元，手上的钱一共有3元
然后，可以选{3,1}，也可以选{3,2}，肯定是选{3,2}划算，结束之后的金额是 5元
然后，可以选{3,1}，和{5,3}，肯定选{5,3}，此时最后的钱是 8元
由于已经达到了K个项目（3个），所以结束。最优的方案下答案是8元

贪心思路：每次选择当前资金下，收益最高的项目来做，用两个堆来实现
1. 准备两个堆，一个小根堆，一个大根堆，小根堆以cost来组织，大根堆以profit来组织
2. 先将所有项目都入下根堆，在上面的例子中，会变成 小根堆：[ {1,2}, {3,1}, {3,2}, {5,3} ], 大根堆：[] 为空， curMoney = 1
3. 将小根堆中，所有花费 <= 当前资金的，全部弹出到大根堆，小根堆：[{3,1}, {3,2}, {5,3}], 大根堆：[{1,2}]，curMoney = 1
4. 弹出大根堆堆顶的项目{1,2}来做，curMoney=3，然后重复2步骤，{3,1}，{3,2} 弹出到大根堆，此时小根堆：[{5,3}]，大根堆：[{3,2},{3,1}]
5. 弹出大根堆堆顶的项目{3,2}来做，curMoney=5, 然后重复2步骤，{5,3}弹出到大根堆，此时小根堆为空，大根堆为[{5,3}, {3,1}]
6. 弹出大根堆堆顶的项目{3,2}来做，curMoney=8, 由于此时达到了K(3)次，所以结束

可以看出，这个策略是很自然而然能想到的，每次在当前资金下，选择收益最高的，最后得到的钱一定是最多的。
*/

// FindMaximizedCapital 贪心策略：用堆来实现，每次选择当前资金下，收益最高的项目来做
// costs和profits的长度一定是相等的
// k表示最多做的项目个数
// m表示一开始的启动资金
func FindMaximizedCapital(k, w int, costs, profits []int) int {
	// 1. 准备一个小根堆和一个大根堆，泛型是program类型
	minCostHeap := class_07.NewMyHeapAny[program](func(a program, b program) bool {
		return a.cost < b.cost // 花费小的优先的小根堆
	})
	maxProfitHeap := class_07.NewMyHeapAny[program](func(a program, b program) bool {
		return a.profit > b.profit // 收益高的优先的大根堆
	})
	// 2. 将所有数据，变成项目，并加入小根堆
	for i := 0; i < len(costs); i++ {
		minCostHeap.Push(program{
			cost:   costs[i],
			profit: profits[i],
		})
	}
	// 3. 最多做K个项目，循环K次
	for i := 0; i < k; i++ {
		// 4. 每次先将所有小根堆中，花费小于等于当前资金的项目，弹入大根堆
		for !minCostHeap.IsEmpty() && minCostHeap.Peek().cost <= w {
			maxProfitHeap.Push(minCostHeap.Pop())
		}
		if maxProfitHeap.IsEmpty() {
			return w
		}
		// 5. 从大根堆中选择一个项目来做，并累计收益
		w += maxProfitHeap.Pop().profit

	}
	return w
}

// 项目
type program struct {
	cost   int
	profit int
}

// FindMaximizedCapital2 暴力方法，每次穷举所有能做的项目做，并返回最大收益
func FindMaximizedCapital2(k, w int, costs, profits []int) int {
	// 1. 先将数组加工成项目数组
	programs := make([]program, len(costs))
	for i := 0; i < len(costs); i++ {
		programs[i] = program{cost: costs[i], profit: profits[i]}
	}
	// 2. 递归获取最大的收益，如果递归函数是对的，那么下面这样传参一定能得到最优解
	// programs 剩余的项目是全部项目
	// 当前已经做了的项目数是0
	return findMaximizedCapitalProcess(programs, 0, k, w)
}

// 递归，当已经选择了一些项目做，并且知道当前的资金总额，返回剩下的项目全部做完或者做满K个后最终得到最多的钱
// programs 剩余的所有项目
// count 之前已经做了count个项目
// k 最多能做k个项目
// w 之前做了count个项目之后的总资金
func findMaximizedCapitalProcess(programs []program, count, k, w int) int {
	// 1. 如果之前已经做了k个项目了，则不能继续做了，资金直接返回之前的资金w
	if count >= k {
		return w
	}
	// 2. 如果没有达到K个，则穷举所有的项目，尝试看
	curProfit := w
	for i, p := range programs {
		if p.cost <= w {
			curProfit = utils.Max(curProfit, findMaximizedCapitalProcess(removeAndCopy(programs, i), count+1, k, w+p.profit))
		}
	}
	return curProfit
}

// removeAndCopy 从项目列表中移除一个项目，并返回
func removeAndCopy(programs []program, index int) []program {
	if len(programs) == 0 {
		return nil
	}
	ans := make([]program, len(programs)-1)
	ansi := 0
	for i, p := range programs {
		if i != index {
			ans[ansi] = p
			ansi++
		}
	}
	return ans
}
