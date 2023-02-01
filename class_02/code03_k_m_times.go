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
//  4. 遍历32位数组，每一位与m做取模运算，如果不是0，则将 rst |= 1 << i (i是当前循环的位)
// 分析复杂度：
// 1.虽然是双层遍历，但是内部遍历是固定长度32，所以时间复杂度是O(N)
// 2.只分配了有限几个变量，还有固定32长度的数组，额外空间复杂度O(1)

// 扩展：改一下题目，如果题目是其他数都出现m次，剩下一种数不一定出现k次，但是出现的次数肯定小于m
// 如果正好出现k次，则返回这个数，如果没有出现k次，则返回-1，题目要怎么改？
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
		// 如果是m的整数倍，说明第i位上是0，跳过
		if t[i]%m == 0 {
			continue
		}
		// 如果取模之后是k，说明正好出现k次，则或进去
		if t[i]%m == k {
			result |= 1 << i
		} else {
			// 否则，说明一定没有出现k次，直接返回-1
			return -1
		}
	}
	// 假如输入的数组中，0出现了3次，其他数出现了5次，但是k = 2, m=5
	// 按照题目，应该返回-1，但是由于0的二进制没有1，所以t数组每一位都是必然被m整除的
	// 导致一定会一直进入上面的 continue，最终返回0，实际应该返回-1
	// 如果结果是0，判断0是否真的出现了k次，如果不是，则返回-1
	if result == 0 {
		var count0 int32
		for _, num := range arr {
			if num == 0 {
				count0++
			}
		}
		if count0 != k {
			return -1
		}
	}

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
	// 填数字，个数为k的数字
	kNum := utils.GenerateRandInt32(maxNum) // [-maxNum, maxNum]
	var i int32                             // i标记当前正在填入的索引
	for ; i < k; i++ {
		arr[i] = kNum
	}
	// 定义一个map，用来统计一个数是否被加入过
	numMap := make(map[int32]int)
	numMap[kNum] = 0
	// 循环产生个数为m的数字，并填入
	for j := int32(0); j < mSize; j++ {
		var mNum int32
		// 不断生成数，如果加入过，就重新生成
		for {
			mNum = utils.GenerateRandInt32(maxNum)
			if _, ok := numMap[mNum]; !ok {
				break
			}
		}
		numMap[mNum] = -1
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
	// 数组中实际有一个数出现了k次，根据扩展思考的题目要求，保证50%的概率返回k，50%的概率不返回k
	if rand.Float64() < 0.5 {
		k = m - (rand.Int31n(m-1) + 1)
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
	// 对数器测试
	rand.Seed(time.Now().UnixNano())
	testTimes := 100000
	fmt.Println("开始测试！！！")
	for i := 0; i < testTimes; i++ {
		arr, k, m := generateArray(100, 20, 1000)
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
