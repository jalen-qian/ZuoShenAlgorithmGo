package main

import "fmt"

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
func onlyKTimes(arr []int32, k int32, m int32) int {
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
	result := 0
	for i := 0; i < 32; i++ {
		// 如果取模不是0，说明这个数在第i位是1
		if t[i]%m != 0 {
			result = (result << 1) | 1
		}
	}
	// 上面的流程忽略了出现k次的数本身就是0的情况，因为如果这个数是0，则t中每个数必能被m整除

	return result
}

func main() {
	// 1. 进行简单测试，假设 3 出现 两次，5 8 7都出现3次
	arr := []int32{5, 0, 8, 7, 8, 0, 5, 7, 8, 9, 5, 9, 7, 9}
	rst := onlyKTimes(arr, 2, 3)
	fmt.Printf("结果为：%d\n", rst)
}
