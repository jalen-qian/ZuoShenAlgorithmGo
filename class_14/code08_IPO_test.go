package class_14

import (
	"ZuoShenAlgorithmGo/utils"
	"math/rand"
	"testing"
	"time"
)

func TestFindMaximizedCapital(t *testing.T) {
	t.Log("Test start...")
	testTimes := 100000
	maxCost := 100
	maxProfits := 200
	maxSize := 10
	maxK := 100
	maxW := 10
	for i := 0; i < testTimes; i++ {
		costs, profits, k, w := generatePrograms(maxCost, maxProfits, maxSize, maxK, maxW)
		ans1 := FindMaximizedCapital(k, w, costs, profits)
		ans2 := FindMaximizedCapital2(k, w, costs, profits)
		if ans1 != ans2 {
			t.Errorf("Test failed!\n costs:%s\n profits:%s\n k:%d\n w:%d\n ans1:%d\n ans2:%d\n ",
				utils.SprintList(costs), utils.SprintList(profits), k, w, ans1, ans2)
			return
		}
	}
	t.Log("Test succeed!")
}

func generatePrograms(maxCost, maxProfit, maxSize, maxK, maxW int) ([]int, []int, int, int) {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := myRand.Intn(maxSize + 1)
	costs := make([]int, size)
	profits := make([]int, size)
	for i := 0; i < size; i++ {
		costs[i] = myRand.Intn(maxCost + 1)
		profits[i] = myRand.Intn(maxProfit + 1)
	}
	return costs, profits, rand.Intn(maxK + 1), rand.Intn(maxW + 1)
}
