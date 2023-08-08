package class_07

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// 测试

// 简单测试暴力方式
func TestTopKCompare(t *testing.T) {
	arr := []int{3, 3, 1, 2, 1, 2, 5}
	//            3     3     1     2     1      2      5
	op := []bool{true, true, true, true, false, true, false}
	k := 2
	// 正确答案应该是：
	// [[3] [3] [1 3] [1 3] [3 2] [3 2] [3 2]]
	fmt.Println(TopKCompare(arr, op, k))

	// 第二组用例
	arr = []int{1, 5, 3, 0, 1, 0, 0, 0, 0, 2}
	//          1,     5,     3 ,   0,     1,    0,    0,     0,     0,     2
	op = []bool{true, false, true, true, false, true, false, false, false, true}
	k = 1
	// 正确答案应该是：
	//  [[1],[1],[1],[1],[3],[0],[0],[3],[3],[3]]
	fmt.Println(TopK(arr, op, k))
}

// 使用对数器的方式，测试使用手写堆的实现
func TestTopK(t *testing.T) {
	t.Log("测试开始...")
	for i := 0; i < 100000; i++ {
		// 生成一个随机的事件数组
		arr, op, k := generateEvents(1000, 0, 100)
		// 分别用两种方式计算，判断结果是否一致
		ans1 := TopK(arr, op, k)
		ans2 := TopKCompare(arr, op, k)
		if !isSameAnswer(ans1, ans2) {
			t.Errorf("Fucking Fucked! arr:%v\n op:%v\n k:%v\n ans1:%v \n ans2:%v", arr, op, k, ans1, ans2)
			return
		}
	}
	t.Log("Great!!!")
}

// 判断两组答案是否是相同的
func isSameAnswer(ans1 [][]int, ans2 [][]int) bool {
	// 如果答案个数不一致，则不相同
	if len(ans1) != len(ans2) {
		return false
	}
	// 长度相等，判断里面的每个子答案是否相等
	for i, subAns1 := range ans1 {
		subAns2 := ans2[i]
		// 子答案个数不相等，不一致
		if len(subAns1) != len(subAns2) {
			return false
		}
		// 每个子答案都按照从小到大排序
		sort.SliceStable(subAns1, func(i, j int) bool {
			return subAns1[i] < subAns1[j]
		})
		sort.SliceStable(subAns2, func(i, j int) bool {
			return subAns2[i] < subAns2[j]
		})
		// 遍历，判断是否每个值都相等
		for j, _ := range subAns1 {
			if subAns2[j] != subAns1[j] {
				return false
			}
		}
	}
	return true
}

// 随机生成一组事件（一个事件表示一个用户买或者退货商品）
//
//		maxLen 最大事件个数
//		minId 用户id的最小值  2
//	 maxId 用户id的最大值  5  5-2 = 3 => [0,3] +2 => [2,5]
func generateEvents(maxLen int, minId, maxId int) ([]int, []bool, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := r.Intn(maxLen + 1) // 数组长度 [0, maxLen]
	arr := make([]int, length)
	op := make([]bool, length)
	for i := 0; i < length; i++ {
		id := r.Intn(maxId-minId+1) + minId
		arr[i] = id
		// 一半概率是买，一半概率是退货
		if r.Float64() > 0.5 {
			op[i] = true
		} else {
			op[i] = false
		}
	}
	// k一般远小于数组长度，从数组长度的1/10范围内取一个随机数，但是要保证必须是>0的
	k := 0
	if length > 0 {
		kRange := length / 10
		if kRange == 0 {
			kRange = length
		}
		k = r.Intn(kRange + 1)
		if k == 0 {
			k = 1
		}
	}
	return arr, op, k
}
