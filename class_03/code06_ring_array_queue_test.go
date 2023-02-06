package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRingArrayQueue(t *testing.T) {
	fmt.Println("测试开始")
	// 测试500000次
	testTimes := 500000
	// 每次测试操作的次数（push或者poll)100次
	oneTestOperatorNum := 1000
	queueLimit := 10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		// 创建一个队列，确定泛型类型为int类型，将使用双向队列实现的，和单链表实现的做对比
		queue1 := NewRingArrayQueue[int](10)
		queue2 := &MyQueueWithLinkedList[int]{}
		for j := 0; j < oneTestOperatorNum; j++ {
			// 随机生成一个数
			value := rand.Intn(1001) - rand.Intn(1001) // [-1000, 1000]
			// 如果队列是空的，则必然加入队列
			if queue1.IsEmpty() {
				// 同时加入队列
				err := queue1.Push(value)
				if err != nil {
					t.Fatalf("出错了，队列是空的，加入数据报错！")
				}
				queue2.Push(value)
			} else {
				// 队列不是空的，则50%的概率决定是入队还是出队
				p := rand.Float32()
				if p < 0.5 {
					// 同时加入队列
					err := queue1.Push(value)
					if err != nil {
						// 如果出错了，一定是队列满了，判断队列2是否满了，如果没有满，则报错
						if queue2.size != queueLimit {
							t.Fatalf("队列1满了，队列2没满")
						}
					} else {
						queue2.Push(value)
					}
				} else {
					// 同时弹出，并判断弹出的数是否相同
					ans1, err := queue1.Poll()
					if err != nil {
						if queue2.size != 0 {
							t.Fatalf("出错了，队列1弹出报错，说明队列1空了，队列2却没空")
						}
					} else {
						ans2 := queue2.Poll()
						if ans1 != ans2 {
							t.Fatalf("出错了，弹出值不相等，队列1弹出%d,队列2弹出%d\n", ans1, ans2)
							return
						}
					}
				}
			}
		}
	}
}
