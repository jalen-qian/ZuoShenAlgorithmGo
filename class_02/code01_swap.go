package main

import "fmt"

// 练习1：不使用第额外的变量，交换数组中的两个数
// 注意：i != j，如果i==j，则
func swap(arr []int, i, j int) {
	if i == j {
		return
	}
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

func main() {
	arr := []int{5, 7, 6}
	swap(arr, 1, 2)
	fmt.Println(arr)
}
