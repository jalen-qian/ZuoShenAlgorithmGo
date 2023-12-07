package class_15

import (
	"math/rand"
	"testing"
	"time"
)

// 测试岛问题2 使用两种方式实现的并查集
func TestNumIslands21(t *testing.T) {
	t.Log("Test start...")
	testTimes := 10000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < testTimes; i++ {
		// 初始化一个m * n 的矩阵，并随机生成岛屿
		m, n := r.Intn(100)+1, r.Intn(100)+1
		positionLen := r.Intn(m*n) + 1
		positions := make([][]int, positionLen)
		for j := 0; j < positionLen; j++ {
			positions[j] = []int{r.Intn(m), r.Intn(n)}
		}
		ans := numIslands21(m, n, positions)
		ans2 := numIslands22(m, n, positions)
		if !isEqual(ans, ans2) {
			t.Errorf("出错了！\nans1:%v\n ans2:%v\n", ans, ans2)
			return
		}
	}
	t.Log("Test succeed!")
}

func isEqual(ans1 []int, ans2 []int) bool {
	if len(ans1) != len(ans2) {
		return false
	}
	for i, v := range ans1 {
		if ans2[i] != v {
			return false
		}
	}
	return true
}
