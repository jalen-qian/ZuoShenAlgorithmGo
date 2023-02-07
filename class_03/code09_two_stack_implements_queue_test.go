package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// TestTwoStackQueue 测试两个栈实现的队列
func TestTwoStackQueue(t *testing.T) {
	fmt.Println("测试开始")
	// 测试500000次
	testTimes := 100000
	// 每次测试操作的次数（push或者poll)100次
	oneTestOperatorNum := 1000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		// 创建一个栈，确定泛型类型为int类型，将使用两个列表实现的，和双向队列实现的做对比
		queue1 := NewTwoStackQueue[int]()
		queue2 := NewMyQueue[int]()
		for j := 0; j < oneTestOperatorNum; j++ {
			// 随机生成一个数
			value := rand.Intn(1001) - rand.Intn(1001) // [-1000, 1000]
			// 如果队列是空的，则必然加入
			if queue1.IsEmpty() {
				if !queue2.IsEmpty() {
					t.Fatal("出错了，队列1是空的，队列2不是空的\n")
				}
				// 同时加入
				queue1.Push(value)
				queue2.Push(value)
			} else {
				// 队列不是空的，则50%的概率决定是加入还是弹出
				p := rand.Float32()
				// 25 % 的概率执行加入操作
				if p < 0.25 {
					// 同时加入
					queue1.Push(value)
					queue2.Push(value)
				} else if p < 0.5 {
					// 25%的概率执行Peek操作
					// 同时Peek看是否相等
					ans1 := queue1.Peek()
					ans2 := queue2.Peek()
					if ans1 != ans2 {
						t.Fatalf("出错了，Peek()值不相等，队列1：%d,队列2：%d\n", ans1, ans2)
						return
					}
				} else if p < 0.75 {
					// 25%的概率测试出队操作
					// 同时出队，并判断出队的数是否相同
					ans1 := queue1.Poll()
					ans2 := queue2.Poll()
					if ans1 != ans2 {
						t.Fatalf("出错了，出队不相等，队列1出队%d,队列2出队%d\n", ans1, ans2)
						return
					}
				} else {
					// 剩下25%的概率测试Empty()函数
					if queue1.IsEmpty() != queue1.IsEmpty() {
						t.Fatalf("出错了，是否相等，队列1%v,队列2%v\n", queue1.IsEmpty(), queue2.IsEmpty())
						return
					}
				}
			}
		}
	}
}
