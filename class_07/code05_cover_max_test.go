package class_07

import (
	"math/rand"
	"testing"
	"time"
)

// 测试

// GenerateLines 生成很多线段
//
//	maxLength 线段条数最大值
//	minLeft 单条线段左边界最小值
//	maxRight 单条线段右边界最大值
func GenerateLines(maxLength int, minLeft int, maxRight int) [][2]int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 数组长度
	length := random.Intn(maxLength + 1)
	var ans [][2]int
	for i := 0; i < length; i++ {
		left := minLeft + random.Intn(maxRight-minLeft+1)
		right := minLeft + random.Intn(maxRight-minLeft+1)
		// 保证left 一定小于 right
		if left >= right {
			right = left + 1
		}
		line := [2]int{left, right}
		ans = append(ans, line)
	}
	return ans
}

// TestMaxCover 测试最大覆盖线段条数问题
func TestMaxCover(t *testing.T) {
	t.Log("测试开始...")
	for i := 0; i < 500000; i++ {
		lines := GenerateLines(1000, 0, 500)
		maxCover1 := MaxCover1(lines)
		maxCover := MaxCover(lines)
		if maxCover != maxCover1 {
			t.Errorf("测试失败，maxCover1:%d, maxCover:%d", maxCover1, maxCover)
			return
		}
	}
	t.Log("测试成功！！")
}
