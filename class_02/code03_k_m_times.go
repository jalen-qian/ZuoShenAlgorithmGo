package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"time"
)

// 在数组中，只有1种数出现了k次，其他数出现了m次，找到出现了k次的数，并返回，如果没有找到，返回-1
// m > 1, 且 k < m
// 要求：额外空间复杂度O(1)，时间复杂度O(N)
// ps:Go中的int类型在不同位数系统中占内存空间不同，为了避免歧义，这里指的是32位的int
//
// 分析：k是 >= 1的，如果k==0，说明所有数都出现了m次，与题意不符。
// 我们将这些数都看成二进制的形式，统计每一位上为1的数的个数。
// 以第0位为例，假如有15个数的第0位都是1，则记录第0位的值为15。
// 那么这个15可能有什么规律呢？有两种情况：
// a. 出现k次的数，在第0位是1，另外有其他n个数的第0位是1（注意这n个数必然出现m次，n>=0），那么 这里的 15 = n * m + k，
//    由于k < m，则 n*m + k 肯定不是m的倍数。所以15不是m的倍数
// b. 出现k次的数，在第0位是0，另外有其他n个数的第0位是1，那么 这里的 15 = n*m，15是m的倍数
// 这里分析了第0位的情况，梳理成章也能类推到全部的32位。
// 也就是说：“统计每个位上1的个数，如果这个数字是m的倍数，说明出现k次的数，在这位上是1，否则是0”，这样我们将这个数刷出来，就找到出现K次的数了。
//
// 具体步骤如下：
//  1. 定义一个32位的数组，初始值都是0
//  2. 循环给定数组中的数，并循环32次将这个数右移i位并与上1（i为当前循环到第i次），得到的1则累加到32位数组中（统计每位上1的个数）
//  3. 定义一个结果数 rst ，初始为0
//  4. 遍历32位数组，每一位与m做取模运算，如果是0，则将rst左移，并或上1
// 分析复杂度：
// 1.虽然是双层遍历，但是内部遍历是固定长度32，所以时间复杂度是O(N)
// 2.只分配了有限几个变量，还有固定32长度的数组，额外空间复杂度O(1)
func onlyKTimes(arr []int32, k int32, m int32) int32 {
	// 用于记录数组中的数字每位上1的个数
	t := make([]int32, 32)
	for _, num := range arr {
		//  00010101101 循环到第4次时，应该得到0
		//  右移4位，得到
		//  00000001010 与上1
		//& 00000000001 得到0
		for i := 0; i < 32; i++ {
			// 得到0，就加了个寂寞，得到1，就加1，统计个数
			t[i] += (num >> i) & 1
		}
	}
	var result int32
	for i := 0; i < 32; i++ {
		// 如果取模不是0，说明这个数在第i位是1
		result = (result << 1) | (t[i] % m)
	}
	// 上面的流程忽略了出现k次的数本身就是0的情况，因为如果这个数是0，则t中每个数必能被m整除

	return result
}

// 对数器对比函数，使用map来做词频统计
func test(arr []int32, k int32, m int32) int32 {
	countMap := make(map[int32]int)
	for _, num := range arr {
		if _, ok := countMap[num]; ok {
			countMap[num]++
		} else {
			countMap[num] = 1
		}
	}
	for num, count := range countMap {
		if count == int(k) {
			return num
		}
	}
	return -1
}

// 初始化一个符合要求的数组
// maxM m参数的最大值 [2, maxM]
// maxMSize 出现m次的数最多有多少种
// maxNum 数组中的数字的取值范围 [-maxNum, maxNum]
func generateArray(maxM int32, maxMSize int32, maxNum int32) ([]int32, int32, int32) {
	rand.Seed(time.Now().UnixNano())
	// 随机产生一个m
	m := rand.Int31n(maxM-1) + 2 // [2, maxM]
	// 随机产生一个k，保证k < m 且 k > 0
	k := m - (rand.Int31n(m-1) + 1)      // [1,m-1] , m - [1,m-1] 范围是[1, m-1]，一定比m小，且 >= 1
	mSize := rand.Int31n(maxMSize+1) + 1 // [1, maxMSize]
	// 数组长度
	arrLen := mSize*m + k
	arr := make([]int32, arrLen)
	// 填数字，先填个数为k的数字
	kNum := utils.GenerateRandInt32(maxNum) // [-maxNum, maxNum]
	var i int32                             // i标记当前正在填入的索引
	for ; i < k; i++ {
		arr[i] = kNum
	}
	// 循环产生个数为m的数字，并填入
	for j := int32(0); j < mSize; j++ {
		mNum := utils.GenerateRandInt32(maxNum)
		// 如果正好和K数字撞车了，就-1，保证不撞车
		if mNum == kNum {
			mNum--
		}
		// 填入数组中，填m次
		for n := int32(0); n < m; n++ {
			arr[i] = mNum
			i++
		}
	}
	// 最后，将数组中的数字顺序打乱
	for i = 0; i < arrLen; i++ {
		// [0, arrLen-1]中随机找一个位置交换
		swapArr(arr, i, rand.Int31n(arrLen))
	}
	return arr, k, m
}

func swapArr(arr []int32, i, j int32) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	// 1. 进行简单测试，假设 3 出现 两次，5 8 7都出现3次
	arr := []int32{52, 71, 71}
	rst := onlyKTimes(arr, 1, 2)
	fmt.Printf("简单测试结果为：%d\n", rst)
	// 2. 对数器测试
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		arr, k, m := generateArray(5, 10, 100)
		rst := onlyKTimes(arr, k, m)
		rstForTest := test(arr, k, m)
		if rst != rstForTest {
			fmt.Printf("出错了，输出结果：%d, 对数器结果：%d\n", rst, rstForTest)
			fmt.Printf("数组为：%v, k是%d, m是%d\n", arr, k, m)
			return
		}
	}
	fmt.Println("测试通过！！！")
}
