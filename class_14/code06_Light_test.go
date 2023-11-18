package class_14

import (
	"math/rand"
	"testing"
)

// 贪心：最少需要几盏灯？

func TestMinLight(t *testing.T) {
	t.Log("Test starts...")
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		test := generateRandomStr(20)
		ans1 := MinLight1(test)
		ans2 := MinLight2(test)
		if ans1 != ans2 {
			t.Errorf("Test failed, testcase:%s\nans1:%d\nans2:%d\n", test, ans1, ans2)
			return
		}
	}
	t.Log("Test succeed!")
}

func generateRandomStr(length int) string {
	res := make([]rune, rand.Intn(length+1))
	for i := 0; i < len(res); i++ {
		res[i] = '.'
		if rand.Float32() < 0.5 {
			res[i] = 'X'
		}
	}
	return string(res)
}
