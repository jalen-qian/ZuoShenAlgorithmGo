package class_14

import (
	"sort"
	"strings"
)

// 贪心算法
// 从头到尾讲一道利用贪心求解的题目
// 给定一个由字符串组成的数组strings，必须把所有的字符串拼接起来，返回所有可能的拼接结果中，字典序最小的结果

// 思路：这本质是一个贪心排序的问题，因为我们只需要规定一个确定的贪心策略，将整个数组排好序，并且经过这个策略排好序后，
// 整个数组依次拼接就是字典序最小的结果。然后再从0到n-1位置 依次拼接起来返回。
//
// 排序的本质是比较大小，当遇到i位置和j位置的字符串时，谁排前面谁排后面，这就需要规定贪心策略。
// 贪心策略1：i位置和j位置谁字典序小，谁排前面，这个策略显然是不正确的，因为很容易举出反例：
//

// 方法一：暴力方法，获取所有可能的排列组合，并得到字典序最小的
func minConcatenation1(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	// 得到所有可能的排列组合拼接后的结果
	ans := process1(arr)
	// 将这些结果按照字典序从小到大排序
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	// 第1个就是要返回的值
	return ans[0]
}

// 传入一个字符串数组，返回所有排列组合后拼接的字符串结果
// 比如传入 ["ab", "cd"]
// 会返回 ["abcd", "cdab"]
func process1(arr []string) []string {
	ans := make([]string, 0)
	if len(arr) == 0 {
		ans = append(ans, "")
		return ans
	}
	// 以i位置的字符串开头，然后将删除i位置的字符串后，剩下的字符串列表继续排列组合，并在头部拼接上i位置的字符串
	for i := 0; i < len(arr); i++ {
		first := arr[i]
		removedList := removeIndexString(arr, i)
		next := process1(removedList)
		for _, cur := range next {
			ans = append(ans, first+cur)
		}
	}
	return ans
}

// 删除i位置的字符串，返回新的数组
func removeIndexString(arr []string, index int) []string {
	N := len(arr)
	if N == 0 {
		return nil
	}
	ans := make([]string, N-1)
	ansIndex := 0
	for i := 0; i < N; i++ {
		if i != index {
			ans[ansIndex] = arr[i]
			ansIndex++
		}
	}
	return ans
}

func minConcatenation2(arr []string) string {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] < arr[j]+arr[i]
	})
	return strings.Join(arr, "")
}
