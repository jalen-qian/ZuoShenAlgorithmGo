package leetcode

import "sort"

// 第39题，组合总和

func combinationSum(candidates []int, target int) [][]int {
	// 先按照从小到大排序
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var ans [][]int
	for i, num := range candidates {
		sub, isFound := x(candidates, i, target-num)
		if isFound {
			sub = append(sub, target-num)
			ans = append(ans, sub)
		}
	}
	return ans
}

func x(candidates []int, i int, target int) ([]int, bool) {
	if candidates[i] > target {
		return nil, false
	}
	sub, isFound := x(candidates, i, target-candidates[i])
	if isFound {
		sub = append(sub, target-candidates[i])
		return sub, isFound
	}
	return nil, false
}
