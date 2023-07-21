package class_07

import "sort"

type Student struct {
	ID   int
	Age  int
	Name string
}

func compare() {
	arr := []Student{
		{ID: 5, Age: 19, Name: "小红"},
		{ID: 9, Age: 18, Name: "小明"},
		{ID: 2, Age: 20, Name: "小马"},
	}
	sort.SliceStable(arr, func(i, j int) bool {
		// 按照年龄从小到大排序
		if arr[i].Age < arr[j].Age {
			return true
		}
		return false
	})
}
