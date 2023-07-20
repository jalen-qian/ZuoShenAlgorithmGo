package class_06

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
	"testing"
	"time"
)

func TestQuickSort1(t *testing.T) {
	// 测试快排1
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(1000, -200, 200)
		//arr := []int{2, 0, 9, 7, 6}
		arr1 := utils.Copy(arr)
		QuickSort1(arr)
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] <= arr1[j]
		})
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("测试失败，排序后不相等")
			return
		}
	}
	t.Logf("测试通过！！！")
}

func TestQuickSort2(t *testing.T) {
	// 测试快排1
	ts := time.Now()
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSortedSlice(1000, -200, 200)
		//arr := []int{2, 0, 9, 7, 6}
		arr1 := utils.Copy(arr)
		QuickSort2(arr)
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] <= arr1[j]
		})
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("测试失败，排序后不相等")
			return
		}
	}
	// 都是升序数组时，耗时4.4秒
	t.Logf("测试通过，耗时：%s\n", time.Since(ts).String())
}

func TestQuickSort3(t *testing.T) {
	// 测试快排1
	ts := time.Now()
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSortedSlice(1000, -200, 200)
		arr1 := utils.Copy(arr)
		QuickSort3(arr)
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] <= arr1[j]
		})
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("测试失败，排序后不相等")
			return
		}
	}
	// 都是升序数组时，耗时1.97秒，而快排2.0耗时4.4秒
	t.Logf("测试通过，耗时：%s\n", time.Since(ts).String())
}

func TestQuickSort4(t *testing.T) {
	// 测试快排1
	ts := time.Now()
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(1000, -200, 200)
		arr1 := utils.Copy(arr)
		// 使用栈实现的非递归随机快排
		QuickSort4(arr)
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] <= arr1[j]
		})
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("测试失败，排序后不相等")
			return
		}
	}
	// 都是升序数组时，耗时1.97秒，而快排2.0耗时4.4秒
	t.Logf("测试通过，耗时：%s\n", time.Since(ts).String())
}
