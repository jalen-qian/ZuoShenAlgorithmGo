package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"time"
)

// 给定一个有序数组和一个数num，求出>=num的最左侧的位置，没找到则返回-1
func findNearLeft(sortedArr []int, num int) int {
	if len(sortedArr) == 0 {
		return -1
	}
	l, r := 0, len(sortedArr)-1
	for l <= r {
		mid := l + (r-l)>>2
		if sortedArr[mid] >= num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 如果找到的l位置确实 >= num 则l位置即为所求位置
	if sortedArr[l] >= num {
		return l
	}
	// 否则没找到，返回-1
	return -1
}

// findNearLeftOperator 对数器，使用遍历法
func findNearLeftComparator(sortedArr []int, num int) int {
	if len(sortedArr) == 0 {
		return -1
	}
	for i := 0; i < len(sortedArr); i++ {
		if sortedArr[i] >= num {
			return i
		}
	}
	return -1
}

func main() {
	testTimes := 500000
	for i := 0; i < testTimes; i++ {
		sortedArr := utils.GenerateRandomSortedSlice(1000, -1000, 20)
		rand.Seed(time.Now().UnixNano())
		// [-20000, 20000]
		num := rand.Intn(40001) - 20000
		if findNearLeft(sortedArr, num) != findNearLeftComparator(sortedArr, num) {
			fmt.Println("Fucking fucked!!!")
			return
		}
	}
	fmt.Println("Great!!!")
}
