package class_07

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"ZuoShenAlgorithmGo/utils"
)

// 测试

// 生成一个几乎有序的无序数组，原理：先生成一个有序数组，再随机交换，每次交换的两个数相隔距离不超过k
func GenerateArrayDistanceLessK(maxLen int, minNum int, maxStep int, k int) []int {
	// 先生成一个有序的数组
	var arr []int
	// 保证 arr 的长度必须大于k
	for len(arr) < k {
		arr = utils.GenerateRandomSortedSlice(maxLen, minNum, maxStep)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//对数组中的元素随机往后交换位置，但是交换位置的最远距离不超过k
	//记录这个位置是否已经移动过（交换和被交换的位置都算）
	movedMap := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if movedMap[i] {
			continue
		}
		// 如果这个位置没移动过，则随机移动 [0,k]的距离
		distance := r.Intn(k + 1) // [0,k]
		// 如果两个位置都存在，且都没移动过，则交换，并记录移动的位置
		if !movedMap[i+distance] && (i+distance) < len(arr) {
			utils.Swap(arr, i, i+distance)
			movedMap[i] = true
			movedMap[i+distance] = true
		}
	}
	return arr
}

func TestSortedArrDistanceLessK(t *testing.T) {
	for i := 0; i < 100000; i++ {
		randomArr := GenerateArrayDistanceLessK(100, 0, 100, 10)
		copyArr := utils.Copy(randomArr)
		SortedArrDistanceLessK(randomArr, 10)
		sort.SliceStable(copyArr, func(i, j int) bool {
			return copyArr[i] < copyArr[j]
		})
		if !utils.IsEqual(randomArr, copyArr) {
			t.Errorf("出错了！！！")
			return
		}
	}
	t.Log("测试成功！！")
}
