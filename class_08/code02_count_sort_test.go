package class_08

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
	"testing"
)

func TestCountSort(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(1000, 0, 200)
		arr1 := utils.Copy(arr)
		CountSort(arr)
		sort.SliceStable(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("Fucking fucked!!\n arr:%v\n arr1:%v\n", arr, arr1)
			return
		}
	}
	t.Log("Great!!!")
}
