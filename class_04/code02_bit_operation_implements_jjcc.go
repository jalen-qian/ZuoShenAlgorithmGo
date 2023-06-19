package class_04

// Add 使用位运算，实现一个加法
func Add(a int, b int) int {
	// 如果b = 0，那么a + b = a 直接返回
	for b != 0 {
		tmp := a ^ b   // tmp 是 a' 也就是 无进位相加， 当b' = 0时，a' + b' = a' ^ b' = a'
		b = a & b << 1 // b 是 b' 也就是 进位信息
		a = tmp        // 将a赋值成 a'
	}
	// 最后如果进位信息变成0了，就算到了最后的结果，则返回
	return a
}

// oppositeNum 相反数
func oppositeNum(num int) int {
	return Add(^num, 1) // 取反+1
}

// Sub 减法
func Sub(a int, b int) int {
	return Add(a, oppositeNum(b)) // a - b = a + (-b)
}

func Mul(a int, b int) int {
	// 找到更小的数，则另一个数是被叠加的，更小的是叠加次数，优化性能，进行最少次数的加法运算
	//smaller := 0
	//bigger := 0
	//if a < b {
	//	smaller = a
	//	bigger = b
	//} else {
	//	smaller = b
	//	bigger = a
	//}
	//for i := 1; i < smaller; i++ {
	//	bigger = Add(bigger, bigger)
	//}
	//return bigger

	// 更好的解法
	ans := 0
	for b != 0 {
		// b 的最后1位是1
		if b&1 != 0 {
			// 将移动位数之后的a叠加到ans上
			ans = Add(ans, a)
		}
		// a往左移动1位
		a <<= 1
		// b往右移动1位
		b >>= 1
	}
	return ans
}

func Div(a int, b int) int {
	if b == 0 {
		panic("除数不能为0")
	}
	return 0
}
