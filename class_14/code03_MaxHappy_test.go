package class_14

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxHappy(t *testing.T) {
	testTimes := 100000  // 测试次数
	maxHappy := 10000    // 最大快乐值
	maxLevel := 5        // 最大层级数
	maxSubordinates := 5 // 每个员工最大下级数
	t.Log("Test starts...")
	for i := 0; i < testTimes; i++ {
		// 随机生成一个boss
		boss := generateBoss(maxHappy, maxLevel, maxSubordinates)
		maxHappy1, maxHappy2 := MaxHappy1(boss), MaxHappy2(boss)
		if maxHappy1 != maxHappy2 {
			t.Errorf("Test failed, maxHapp1:%d, maxHappy2:%d ", maxHappy1, maxHappy2)
			return
		}
	}
	t.Log("Test successful.")
}

// 初始化boss
func generateBoss(maxHappy, maxLevel, maxSubordinates int) *Employee {
	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 2%的概率返回空树
	if myRand.Float32() < 0.02 {
		return nil
	}
	boss := &Employee{Happy: myRand.Intn(maxHappy + 1)}
	generateSubordinates(boss, myRand, 1, maxHappy, maxLevel, maxSubordinates)
	return boss
}

// 生成员工，挂到boss下面
func generateSubordinates(boss *Employee, myRand *rand.Rand, curLevel, maxHappy, maxLevel, maxSubordinates int) {
	if curLevel > maxLevel {
		return
	}
	// 有多少个员工
	num := myRand.Intn(maxSubordinates + 1)
	if num == 0 {
		return
	}
	for i := 0; i < num; i++ {
		next := &Employee{Happy: myRand.Intn(maxHappy + 1)}
		generateSubordinates(next, myRand, curLevel+1, maxHappy, maxLevel, maxSubordinates)
		boss.Subordinates = append(boss.Subordinates, next)
	}
}
