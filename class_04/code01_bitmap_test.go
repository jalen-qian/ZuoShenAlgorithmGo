package class_04

import (
	"math/rand"
	"testing"
	"time"
)

// TestBitmapSimple 先进行简单测试
func TestBitmapSimple(t *testing.T) {
	b := NewBitmap(2048)
	b.Add(1023)
	t.Logf("1023是否存在：%v", b.Exist(1023))
	b.Add(10)
	t.Logf("10是否存在：%v", b.Exist(10))
	b.Add(20)
	t.Logf("30是否存在：%v", b.Exist(30))
	b.Remove(10)
	t.Logf("10是否存在：%v", b.Exist(10))
	t.Logf("20是否存在：%v", b.Exist(20))
	b.Add(999)
	t.Logf("999是否存在：%v", b.Exist(999))
}

// 对数器测试
func TestBitmap(t *testing.T) {
	// 测试次数 10000 * 100000
	for i := 0; i < 10000; i++ {
		testTimes := 100000
		maxNum := rand.Int63n(2049) // 0 ~ 2049
		rand.Seed(time.Now().UnixNano())
		b := NewBitmap(maxNum) // 可加入数的范围 [0, maxNum-1]
		m := make(map[int64]int)
		for j := 0; j < testTimes; j++ {
			percent := rand.Float64()
			// 33%概率同步加一个数
			var num int64
			if maxNum != 0 {
				num = rand.Int63n(maxNum)
			}
			if percent < 0.3 {
				b.Add(num)
				//t.Logf("添加数字：%v", num)
				m[num] = 0
			} else if percent < 0.6 {
				// 33%概率同步移除一个数
				b.Remove(num)
				//t.Logf("移除数字：%v", num)
				delete(m, num)
			} else {
				// 剩下33%概率判断这个数是否存在
				exist := b.Exist(num)
				_, exist1 := m[num]
				if exist != exist1 {
					t.Fatalf("测试失败，数字%d在位图中是否存在：%v，在map中是否存在：%v", num, exist, exist1)
				}
			}
		}
	}
	t.Logf("测试通过！！！\n")
}
