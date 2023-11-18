package class_14

import (
	"ZuoShenAlgorithmGo/class_07"
	"ZuoShenAlgorithmGo/utils"
	"math"
)

/*
贪心题：用最小的金额分割金条
一块金条切成两半，是需要花费和长度数值一样的铜板的。
比如长度为20的金条，不管怎么切，都要花费20个铜板。 一群人想整分整块金条，怎么分最省铜板?

例如,给定数组{10,20,30}，代表一共三个人，整块金条长度为60，金条要分成10，20，30三个部分。

如果先把长度60的金条分成10和50，花费60; 再把长度50的金条分成20和30，花费50;一共花费110铜板。
但如果先把长度60的金条分成30和30，花费60;再把长度30金条分成10和20， 花费30;一共花费90铜板。
输入一个数组，返回分割的最小代价。

其实这就是著名的哈弗曼编码的问题

贪心策略：
1. 准备一个小根堆，开始时将所有金条长度加入小根堆
2. 当堆不为空时，一直进行下去：
   2.1 每次弹出两个，合并成一个数
   2.2 合并后的数重新加入小根堆
   2.3 每次合并时，记录合并的值，并累加到最终的花费
   2.4 直到堆全部弹出，则退出循环

比如，给定一个金条分割方案是 `[1,1,5,2,3,9]`
1. 先加入小根堆，则堆为：{1,1,2,3,5,9}，记录ans=0
2. 弹出1 1，合并成2，并将2重新加入堆，堆为：{2,2,3,5,9} ans = 2
3. 弹出2 2, 合并成4，并将4重新加入堆，堆为：{3,4,5,9} ans = 2 + 4 = 6
4. 弹出3 4，合并成7，并将7重新加入堆，堆为：{5,7,9} ans = 6 + 7 = 13
5. 弹出5 7，合并成12，并将12重新加入堆，堆为：{9,12} ans = 13 + 12 = 25
6. 弹出9 12，合并成21，此时堆已经空了，结束循环，ans = 25 + 21 = 46
所以：最终最优的方案下，花费最少需要 46 个铜板

*/

// LessMoney1 使用贪心
func LessMoney1(arr []int) int {
	// 1. 新建一个小根堆
	minHeap := class_07.NewMyHeap[int]()
	// 2. 将所有值直接压入堆中
	for _, num := range arr {
		minHeap.Push(num)
	}
	ans := 0
	cur := 0
	for minHeap.Size() > 1 {
		// 先弹出两个值合并
		cur = minHeap.Pop() + minHeap.Pop()
		ans += cur
		// 将弹出的值相加后再次加入堆中
		minHeap.Push(cur)
	}
	return ans
}

// LessMoney2 暴力方法，穷举出所有切割方案，并找出使用铜板最少的
func LessMoney2(arr []int) int {
	// 主函数，如果 processLessMoney 是对的，则传入arr表示之前还没分割过，所以cur为0
	return processLessMoney(arr, 0)
}

// 递归函数，arr表示所有等待合并的数，cur表示之前的和并已经花费的铜板
// 合并：每次合并都表示一次分割的逆过程，比如上面例子中，[1,1,5,2,3,9]， 如果1和9合并，表示将长度为10的金条切割成1和9，花费是10
// 合并后，数组变成[1,5,2,3,10]，下次假如合并 3,10 表示将一个13长度的金条分割成3和10，花费13
// processLessMoney 函数会返回在当前条件下，花费铜板的最小值
func processLessMoney(arr []int, cur int) int {
	// 如果少于2个（0个或者1个实际上不会变成0个，变成1个递归就结束了），则不需要分割了，直接返回之前分割的
	if len(arr) <= 1 {
		return cur
	}
	ans := math.MaxInt
	// 0 和 1 位置的数合并，然后看后序是啥
	// 0 和 1 位置的数合并，然后看后序是啥
	// ...
	// n-2 和 n-1 位置的数合并，看看后序是啥
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			ans = utils.Min(ans, processLessMoney(copyAndMergeTwo(arr, i, j), cur+arr[i]+arr[j]))
		}
	}
	return ans
}

// 将i和j位置的数合并后，拷贝到arr里
// 比如 [1,4,5,2,3] i = 1 j = 3
// 将会返回 [1,5,3,6]
func copyAndMergeTwo(arr []int, i int, j int) []int {
	// ans数组大小肯定不arr小1
	ans := make([]int, len(arr)-1)
	// 当前填入到了ans数组的第几个
	ansi := 0
	for arri, num := range arr {
		// 不在要拷贝的数里面，则依次填入
		if arri != i && arri != j {
			ans[ansi] = num
			ansi++
		}
	}
	// 最后将合并后的填入ans
	ans[ansi] = arr[i] + arr[j]
	return ans
}
