package class_15

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNumberOfIsland(t *testing.T) {
	testTimes := 1000
	t.Log("Test start...")
	for i := 0; i < testTimes; i++ {
		grids := generateRandomGrids(300, 300)
		gridsCopy := copyGrids(grids)
		gridsCopy2 := copyGrids(grids)
		ans1 := numIslands1(grids)      // 使用数组实现的并查集
		ans2 := numIslands2(gridsCopy)  // 使用map实现的并查集
		ans3 := numIslands3(gridsCopy2) // 感染方式
		if ans1 != ans2 {
			t.Errorf("Test failed!\n ans1:%d\n ans2:%d\n and3:%d\n ", ans1, ans2, ans3)
			return
		}
	}
	t.Log("Test succeed!")
}

// 测试3种方式，执行的时间对比
// 结论：1000 * 1000 的矩阵
// 感染方法：8ms 并查集(数组) 15ms 并查集(map实现) 437ms
// 10000 * 10000 的矩阵
// 感染方法：801ms 并查集(数组) 1551ms
// 结论：时间复杂度都是O(N)，但是并查集map实现，map底层的常数时间大，所以时间远高于感染方式和并查集数组实现的方式
// 感染方式最快，并查集数组实现的方式消耗的时间略小于感染方式时间的两倍，但是远低于map实现方式。
func TestNumIslandTime(t *testing.T) {
	var grids1 [][]byte
	var grids2 [][]byte
	var grids3 [][]byte
	row, col := 1000, 1000
	grids1 = generateGrids(row, col)
	grids2 = copyGrids(grids1)
	grids3 = copyGrids(grids1)
	var startTime, endTime time.Time
	fmt.Println("感染方法、并查集(map实现)、并查集(数组实现)的运行结果和运行时间")
	fmt.Printf("随机生成的二维矩阵规模：%d * %d \n", row, col)

	startTime = time.Now()
	fmt.Printf("感染方法的运行结果：%d\n", numIslands3(grids1))
	endTime = time.Now()
	fmt.Printf("感染方法的运行时间：%d ms\n", endTime.Sub(startTime).Milliseconds())
	fmt.Println()

	startTime = time.Now()
	fmt.Printf("并查集(map实现)的运行结果：%d\n", numIslands2(grids2))
	endTime = time.Now()
	fmt.Printf("并查集(map实现)的运行时间：%d ms\n", endTime.Sub(startTime).Milliseconds())
	fmt.Println()

	startTime = time.Now()
	fmt.Printf("并查集(数组实现)的运行结果：%d\n", numIslands1(grids3))
	endTime = time.Now()
	fmt.Printf("并查集(数组实现)的运行时间：%d ms\n", endTime.Sub(startTime).Milliseconds())
	fmt.Println()
	fmt.Println()

	row = 10000
	col = 10000
	grids1 = generateGrids(row, col)
	grids2 = copyGrids(grids1)
	fmt.Println("感染方法、并查集(数组实现)的运行结果和运行时间")
	fmt.Printf("随机生成的二维矩阵规模：%d * %d \n", row, col)

	startTime = time.Now()
	fmt.Printf("感染方法的运行结果：%d\n", numIslands3(grids1))
	endTime = time.Now()
	fmt.Printf("感染方法的运行时间：%d ms\n", endTime.Sub(startTime).Milliseconds())
	fmt.Println()

	startTime = time.Now()
	fmt.Printf("并查集(数组实现)的运行结果：%d\n", numIslands1(grids2))
	endTime = time.Now()
	fmt.Printf("并查集(数组实现)的运行时间：%d ms\n", endTime.Sub(startTime).Milliseconds())
	fmt.Println()
}

func generateGrids(row, col int) [][]byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ans := make([][]byte, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if r.Float64() < 0.5 {
				ans[i][j] = '1'
			} else {
				ans[i][j] = '0'
			}
		}
	}
	return ans
}

func generateRandomGrids(maxRow, maxCol int) [][]byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	row := r.Intn(maxRow) + 1 // [0,maxRow) [1, maxRow]
	col := r.Intn(maxCol) + 1 // [0,maxCol) [1, maxCol]
	ans := make([][]byte, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if r.Float64() < 0.5 {
				ans[i][j] = '1'
			} else {
				ans[i][j] = '0'
			}
		}
	}
	return ans
}

func copyGrids(grids [][]byte) [][]byte {
	row, col := len(grids), len(grids[0])
	ans := make([][]byte, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]byte, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ans[i][j] = grids[i][j]
		}
	}
	return ans
}
