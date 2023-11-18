package class_14

import (
	"math/rand"
	"testing"
)

func TestLessMoney(t *testing.T) {
	t.Log("Test start...")
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		arr := generateRandomArray(6, 1000)
		ans1 := LessMoney1(arr)
		ans2 := LessMoney2(arr)
		if ans1 != ans2 {
			t.Errorf("Test failed! arr:%v\n ans1:%d\n ans2:%d\n", arr, ans1, ans2)
			return
		}
	}
	t.Log("Test succeed!")
	return
}

func generateRandomArray(maxLength, maxNum int) []int {
	length := rand.Intn(maxLength + 1)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = rand.Intn(maxNum) + 1 // [1, maxNum]
	}
	return ans
}
