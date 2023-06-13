package class_05

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
	"testing"
)

func TestMergeSortByIter(t *testing.T) {
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(999, -500, 500)
		arr1 := utils.Copy(arr)
		MergeSortByIter(arr)
		// 用系统的排序做对数器
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("测试出错了！")
			return
		}
	}
	t.Log("测试成功！！！")
}
