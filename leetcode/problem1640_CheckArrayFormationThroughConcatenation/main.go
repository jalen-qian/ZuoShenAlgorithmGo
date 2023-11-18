package main

// 1640题：能否连接形成数组
// url: https://leetcode.cn/problems/check-array-formation-through-concatenation/description/

func canFormArray(arr []int, pieces [][]int) bool {
	piecesMap := make(map[int]int)
	for index, p := range pieces {
		piecesMap[p[0]] = index
	}
	for i := 0; i < len(arr); {
		index, ok := piecesMap[arr[i]]
		if !ok {
			return false
		}
		for _, x := range pieces[index] {
			if arr[i] != x {
				return false
			}
			i++
		}
	}
	return true
}

func main() {

}
