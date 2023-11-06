package class_14

import (
	"math/rand"
	"testing"
	"time"
)

func TestMinConcatenation(t *testing.T) {
	testTimes := 100000 // 测试次数
	maxStrLength := 5   // 样本中最大字符串长度
	maxArrayLength := 6 // 样本中数组最大大小
	t.Log("测试开始...")
	for i := 0; i < testTimes; i++ {
		arr := generateRandomStrArr(maxStrLength, maxArrayLength)
		arr1 := copyStringArray(arr)
		arr2 := copyStringArray(arr)
		ans1 := minConcatenation1(arr)
		ans2 := minConcatenation2(arr1)
		if ans1 != ans2 {
			t.Errorf("测试失败\n 数组为：%v\n 暴力方法结果：%s\n 贪心方法结果: %s", arr2, ans1, ans2)
			return
		}
	}
	t.Log("测试成功")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func generateRandomStrArr(maxStrLength, maxLength int) []string {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ans := make([]string, 0)
	n := myRand.Intn(maxLength + 1)
	for i := 0; i < n; i++ {
		ans = append(ans, generateRandomString(maxStrLength))
	}
	return ans
}

func generateRandomString(maxLength int) string {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := myRand.Intn(maxLength+1) + 1
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[myRand.Intn(len(letterBytes))]
	}
	return string(b)
}

func copyStringArray(arr []string) []string {
	// 创建一个新的字符串数组，长度与输入数组相同
	copiedArray := make([]string, len(arr))

	// 使用循环将每个元素复制到新数组中
	for i, value := range arr {
		copiedArray[i] = value
	}

	return copiedArray
}
