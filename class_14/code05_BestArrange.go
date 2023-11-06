package class_14

import (
	"sort"

	"ZuoShenAlgorithmGo/utils"
)

// 贪心：返回最多的会议室宣讲场次

// Program 项目会议宣讲
type Program struct {
	Start int // 开始时间
	End   int // 结束时间
}

// BestArrange1 暴力方法：穷举所有情况，并返回选择的会议场数最多情况下的场数
// 使用递归实现
func BestArrange1(programs []Program) int {
	// 主函数传入所有项目，开始时间是0，之前已经安排的场数也是0
	return bestArrange1Process(programs, 0, 0)
}

// @param leftPrograms表示选择了某个项目后，还剩下的所有项目
// @param timeLine 选择完某个项目开完会议后的当前时间线
// @param done 之前已经安排了多少场会议
// 返回这种情况下，最大的场数
func bestArrange1Process(leftPrograms []Program, timeLine int, done int) int {
	// 如果剩下的项目为空，则直接返回之前已经安排的场数
	if len(leftPrograms) == 0 {
		return done
	}
	max := done
	// 穷举选择剩下的项目中，所有的项目，从当前时间开始，看哪一场最好
	for i := 0; i < len(leftPrograms); i++ {
		// 当前已经选了，将剩下的项目，除去当前已经选的，拷贝所有的，并递归
		if leftPrograms[i].Start >= timeLine {
			next := copyButExcept(leftPrograms, i)
			// 递归，时间线是当前选择的项目结束时间，已经完成的场次是当前done+1
			max = utils.Max(max, bestArrange1Process(next, leftPrograms[i].End, done+1))
		}
	}
	return max
}

func copyButExcept(leftPrograms []Program, index int) []Program {
	var ans []Program
	for i, program := range leftPrograms {
		if i != index {
			ans = append(ans, program)
		}
	}
	return ans
}

// BestArrange2 返回最多可以安排的场次，贪心策略：每次优先选择结束时间最早的。
func BestArrange2(programs []Program) int {
	// 先按照结束之间从早到晚排好序
	sort.Slice(programs, func(i, j int) bool {
		return programs[i].End < programs[j].End
	})
	curTime := 0 //当前时间点，从0开始
	count := 0   // 统计最后安排的次数
	for _, program := range programs {
		if program.Start >= curTime {
			count++
			curTime = program.End
		}
	}
	return count
}
