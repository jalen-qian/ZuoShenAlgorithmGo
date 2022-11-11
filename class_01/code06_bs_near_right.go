package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"time"
)

// 给定一个有序数组，和数num，找到<=num最右侧的位置
// 如果没有找到，则返回 -1
// 要求时间复杂度 O(logN)
func findNearRight(sortedArr []int, num int) int {
	if len(sortedArr) == 0 {
		return -1
	}
	L, R := 0, len(sortedArr)-1
	// 二分法
	for L <= R {
		mid := L + (R-L)>>2
		// 往右扩充
		if sortedArr[mid] <= num {
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	if sortedArr[L] <= num {
		return L
	}
	return -1
}

// 使用O(N)流程实现的对数器
func findNearRightComparator(sortedArr []int, num int) int {
	if len(sortedArr) == 0 {
		return -1
	}
	// 从最右侧开始找，直到找到第一个 <= num的位置
	for i := len(sortedArr) - 1; i >= 0; i-- {
		if sortedArr[i] <= num {
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
		if findNearRight(sortedArr, num) != findNearRightComparator(sortedArr, num) {
			fmt.Println("Fucking fucked!!!")
			return
		}
	}
	fmt.Println("Great!!!")
}
