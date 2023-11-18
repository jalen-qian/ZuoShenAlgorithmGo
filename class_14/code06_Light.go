package class_14

import (
	"ZuoShenAlgorithmGo/utils"
	"math"
)

// 贪心：最少需要几盏灯？

// 给定一个字符串str，只由‘X’和‘.’两种字符构成。
// ‘X’表示墙，不能放灯，也不需要点亮
// ‘.’表示居民点，可以放灯，需要点亮
// 如果灯放在i位置，可以让 i-1，i 和i+1 三个位置被点亮
// 返回如果点亮str中所有需要点亮的位置，至少需要几盏灯

// MinLight1 贪心策略：不断遍历字符串的每个字符，并决定这个位置是否要安装一个灯，安装完灯后，跳到第一个没被当前灯照亮的位置
// 对于当前遍历的 i 位置：
//  1. 如果遇到'X'，直接跳过，因为'X'位置不能放灯
//  2. 如果遇到'.'，可以在i或者i+1位置放灯，取决于i+1位置是什么：
//     2.1 如果 i + 1 是'x'，则给i位置放灯，并往前跳两步，因为没有照亮 i + 2的位置
//     2.2 如果 i + 1 是'.'，则给i+1位置放灯，并往前跳三步，因为i位置，i+1位置，i+2位置都被照亮了，跳到没被照亮的i+3位置
//     2.3 如果 i + 1 位置已经越界了，则i位置亮灯
//
// 综上，如果是'.'则一定会放一盏灯，
func MinLight1(road string) int {
	lightNumber := 0
	i := 0
	// 保证i位置不越界
	for i < len(road) {
		// 遇到了墙，则跳过
		if road[i] == 'X' {
			i++
			continue
		}
		// 遇到了'.'，一定会放灯
		lightNumber++
		if i+1 < len(road) && road[i+1] == '.' {
			i += 3
		} else {
			i += 2
		}
	}
	return lightNumber
}

// MinLight2 方法二：暴力方法，遍历每个.的位置，自由选择亮灯或者不亮灯，并返回最少能点多少灯
func MinLight2(road string) int {
	spots := []rune(road)
	// 递归
	return minLightProcess(spots, 0, make(map[int]bool))
}

// spots 所有的地点（可能是墙，也可能是居民点）
// index 当前来到i位置，[index....N-1]位置随意放灯，并返回在都点亮的情况下，最少的灯数
// lights 已经放灯的位置集合，存放的是[0....index-1]所有已经放灯的位置 light[i] == true 表示这个位置已经放灯了
// 选择出所有能照亮.的方案中，放置灯最少的方案。
// 返回这种情况下，最少放多少灯
func minLightProcess(spots []rune, index int, lights map[int]bool) int {
	// 结束的时候
	if index == len(spots) {
		for i := 0; i < len(spots); i++ {
			// 当前位置是.的话，判断当前位置是否被照亮了，如果没照亮，说明这个方案是失败的，返回最大的INT
			if spots[i] != 'X' {
				if !lights[i-1] && !lights[i] && !lights[i+1] {
					return math.MaxInt
				}
			}
		}
		total := 0
		for _, hasLight := range lights {
			if hasLight {
				total++
			}
		}
		return total
	} else {
		no := minLightProcess(spots, index+1, lights)
		yes := math.MaxInt
		if spots[index] == '.' {
			lights[index] = true
			yes = minLightProcess(spots, index+1, lights)
			lights[index] = false
		}
		return utils.Min(yes, no)
	}
}
