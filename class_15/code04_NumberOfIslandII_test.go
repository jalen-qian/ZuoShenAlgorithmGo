package class_15

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试岛问题2
func TestNumIslands21(t *testing.T) {
	t.Log("Test start...")
	testTimes := 1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < testTimes; i++ {
		// 初始化一个m * n 的矩阵，并随机生成岛屿
		m, n := r.Intn(10000)+1, r.Intn(10000)+1
		positionLen := r.Intn(m*n) + 1
		positions := make([][]int, positionLen)
		for j := 0; j < positionLen; j++ {
			positions[j] = []int{r.Intn(m), r.Intn(n)}
		}
		ans := numIslands21(m, n, positions)
		fmt.Println(ans)
	}
	t.Log("Test succeed!")
}
