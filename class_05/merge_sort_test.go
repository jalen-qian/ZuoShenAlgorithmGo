package class_05

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
	"testing"
)

/*
*
归并排序单测，用对数器测试
*/
func TestMergeSort(t *testing.T) {
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(1000, -500, 500)
		arr1 := utils.Copy(arr)
		MergeSort(arr)
		// 用系统的排序做对比
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
