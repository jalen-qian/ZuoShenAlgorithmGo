package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMyQueueAndStack(t *testing.T) {
	fmt.Println("测试开始")
	// 测试500000次
	testTimes := 500000
	// 每次测试操作的次数（push或者poll)100次
	oneTestOperatorNum := 100
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		// 创建一个队列，确定泛型类型为int类型，将使用双向队列实现的，和单链表实现的做对比
		queue1 := NewMyQueue[int]()
		queue2 := &MyQueueWithLinkedList[int]{}
		for j := 0; j < oneTestOperatorNum; j++ {
			// 随机生成一个数
			value := rand.Intn(1001) - rand.Intn(1001) // [-1000, 1000]
			// 如果队列是空的，则必然加入队列
			if queue1.IsEmpty() {
				// 同时加入队列
				queue1.Push(value)
				queue2.Push(value)
			} else {
				// 队列不是空的，则50%的概率决定是入队还是出队
				p := rand.Float32()
				if p < 0.5 {
					// 同时加入队列
					queue1.Push(value)
					queue2.Push(value)
				} else {
					// 同时弹出，并判断弹出的数是否相同
					ans1 := queue1.Poll()
					ans2 := queue2.Poll()
					if ans1 != ans2 {
						fmt.Printf("出错了，队列1弹出%d,队列2弹出%d\n", ans1, ans2)
						return
					}
				}
			}
		}

		// 创建栈，使用int类型
		//stack1 := &MyStackWithLinkedList[int]{}
		//stack2 := &MyStackWithSlice[int]{}
		//for j := 0; j < oneTestOperatorNum; j++ {
		//	// 随机生成一个数
		//	value := rand.Intn(1001) - rand.Intn(1001) // [-1000, 1000]
		//	// 如果栈是空的，则必然入栈
		//	if stack1.IsEmpty() {
		//		// 同时入栈
		//		stack1.Push(value)
		//		stack2.Push(value)
		//	} else {
		//		// 栈不是空的，则50%的概率决定是入栈还是弹出
		//		p := rand.Float32()
		//		if p < 0.5 {
		//			// 同时入栈
		//			stack1.Push(value)
		//			stack2.Push(value)
		//		} else {
		//			// 同时弹出，并判断弹出的数是否相同
		//			ans1 := stack1.Pop()
		//			ans2 := stack2.Pop()
		//			if ans1 != ans2 {
		//				fmt.Printf("出错了，栈1弹出%d,栈2弹出%d\n", ans1, ans2)
		//				return
		//			}
		//		}
		//	}
		//}
	}
	fmt.Println("测试结束")
}
