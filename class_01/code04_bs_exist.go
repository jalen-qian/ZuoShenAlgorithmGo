package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"time"
)

// exist 给定一个有序数组和数字，判断这个数字是否存在
// 要求：时间复杂度O(logN)
// 使用二分法
func exist(sortedArr []int, num int) bool {
	if len(sortedArr) == 0 {
		return false
	}
	l, r := 0, len(sortedArr)-1
	for l <= r {
		// 找到中点
		mid := l + (r-l)>>2
		if sortedArr[mid] == num {
			return true
		} else if sortedArr[mid] < num {
			// 下一次从右边部分找
			l = mid + 1
		} else {
			// 下一次从左边部分找
			r = mid - 1
		}
	}
	return false
}

// existComparator 比较器，使用遍历查找
func existComparator(sortedArr []int, num int) bool {
	if len(sortedArr) == 0 {
		return false
	}
	for _, n := range sortedArr {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	// 测试500000次
	testTimes := 500000
	for i := 0; i < testTimes; i++ {
		// 生成一个随机递增切片
		sortedSlice := utils.GenerateRandomSortedSlice(1000, -500, 20)
		rand.Seed(time.Now().Unix())
		// 在[-500, 9500]间生成一个随机数
		num := rand.Intn(10001) - rand.Intn(501)
		if exist(sortedSlice, num) != existComparator(sortedSlice, num) {
			fmt.Println("Fucking fucked!!!")
			return
		}
	}
	fmt.Println("Great!!!")
}
