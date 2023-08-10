package class_08

import (
	"ZuoShenAlgorithmGo/utils"
	"sort"
	"testing"
)

func TestRadixSort1(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 1; i++ {
		arr := utils.GenerateRandomSlice(1000, 0, 500)
		arr = []int{5, 9, 7, 6, 8}
		arr1 := utils.Copy(arr)
		RadixSort1(arr)
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

func TestRadixSort2(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 100000; i++ {
		arr := utils.GenerateRandomSlice(1000, 0, 500)
		arr = []int{5, 9, 7, 6, 8}
		arr1 := utils.Copy(arr)
		RadixSort2(arr)
		RadixSort1(arr1)
		if !utils.IsEqual(arr, arr1) {
			t.Errorf("Fucking fucked!!\n arr:%v\n arr1:%v\n", arr, arr1)
			return
		}
	}
	t.Log("Great!!!")
}
