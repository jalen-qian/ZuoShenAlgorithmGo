package class_14

import (
	"ZuoShenAlgorithmGo/utils"
	"math/rand"
	"testing"
	"time"
)

func TestBestArrange(t *testing.T) {
	testTimes := 100000 // 测试次数
	maxSize := 12
	timeMax := 12
	t.Log("测试开始...")
	for i := 0; i < testTimes; i++ {
		programs := generateRandomPrograms(maxSize, timeMax)
		ans1 := BestArrange1(programs)
		ans2 := BestArrange2(programs)
		if ans1 != ans2 {
			t.Errorf("测试失败 \n 样本：%v \n ans1:%d \n ans2:%d", programs, ans1, ans2)
			return
		}
	}
	t.Log("测试成功")
}

// 返回随机的项目样本
// maxSize 最大项目个数
// timeMax 每个项目，最大的时间
func generateRandomPrograms(maxSize int, timeMax int) []Program {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	ans := make([]Program, myRand.Intn(maxSize+1))
	for i := 0; i < len(ans); i++ {
		t1 := myRand.Intn(timeMax + 1)
		t2 := myRand.Intn(timeMax + 1)
		if t1 == t2 {
			ans[i] = Program{Start: t1, End: t1 + 1}
		} else {
			ans[i] = Program{Start: utils.Min(t1, t2), End: utils.Max(t1, t2)}
		}
	}
	return ans
}
