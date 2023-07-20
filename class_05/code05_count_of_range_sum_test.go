package class_05

import (
	"ZuoShenAlgorithmGo/utils"
	"math/rand"
	"testing"
	"time"
)

// countRangeSumForTest 使用暴力的O(N^3)的方式实现区间和问题
func countRangeSumForTest(nums []int, lower int, upper int) int {
	Length := len(nums)
	if Length == 0 {
		return 0
	}
	var total int
	for i := 0; i < Length; i++ {
		for j := 0; j <= i; j++ {
			// 统计 nums j到i的区间和
			var sum int
			for k := j; k <= i; k++ {
				sum += nums[k]
			}
			if sum >= lower && sum <= upper {
				total++
			}
		}
	}
	return total
}

// TestCountOfRangeSum 测试区间和问题
func TestCountOfRangeSum(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100000; i++ {
		// 随机生成一个数组
		arr := utils.GenerateRandomSlice(300, -500, 500)
		// 随机生成一个区间，区间的数在[-500, 500]范围内
		// [0, 500] - [0, 500] = [-500, 500]
		lower := r.Intn(501) - r.Intn(501)
		upper := r.Intn(501) - r.Intn(501)
		// 如果 lower 比 upper 大，则交换
		if lower > upper {
			tmp := lower
			lower = upper
			upper = tmp
		}
		// 拷贝数组，用于对数器函数测试
		arrCopy := utils.Copy(arr)
		total := countRangeSum(arr, lower, upper)
		forTest := countRangeSumForTest(arrCopy, lower, upper)
		if total != forTest {
			t.Errorf("测试失败：获取到满足条件的区间和个数：%d, 测试函数获取的个数：%d", total, forTest)
			return
		}
	}
	t.Log("测试成功！！！")
}
