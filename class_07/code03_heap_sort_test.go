package class_07

import (
	"ZuoShenAlgorithmGo/utils"
	"testing"
)

func TestHeapSort(t *testing.T) {
	for i := 0; i < 500000; i++ {
		arr := utils.GenerateRandomSlice(1000, -500, 500)
		arr1 := utils.Copy(arr)
		HeapSort(arr)
		Comparator(arr1)
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("测试失败！")
			return
		}
	}
	t.Log("测试成功！！")
}
