package class_03

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// TestTwoQueueStack 测试两个队列实现的栈
func TestTwoQueueStack(t *testing.T) {
	fmt.Println("测试开始")
	// 测试500000次
	testTimes := 100000
	// 每次测试操作的次数（push或者poll)100次
	oneTestOperatorNum := 1000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		// 创建一个栈，确定泛型类型为int类型，将使用两个列表实现的，和双向队列实现的做对比
		stack1 := NewTwoQueueStack[int]()
		stack2 := NewMyStack[int]()
		for j := 0; j < oneTestOperatorNum; j++ {
			// 随机生成一个数
			value := rand.Intn(1001) - rand.Intn(1001) // [-1000, 1000]
			// 如果栈是空的，则必然入栈
			if stack1.IsEmpty() {
				if !stack2.IsEmpty() {
					t.Fatal("出错了栈1是空的，栈2不是空的\n")
				}
				// 同时入栈
				stack1.Push(value)
				stack2.Push(value)
			} else {
				// 栈不是空的，则50%的概率决定是入栈还是弹出
				p := rand.Float32()
				// 25 % 的概率执行入栈操作
				if p < 0.25 {
					// 同时入栈
					stack1.Push(value)
					stack2.Push(value)
				} else if p < 0.5 {
					// 25%的概率执行Peek操作
					// 同时Peek看是否相等
					ans1 := stack1.Peek()
					ans2 := stack2.Peek()
					if ans1 != ans2 {
						t.Fatalf("出错了，Peek值不相等，栈1%d,栈2%d\n", ans1, ans2)
						return
					}
				} else if p < 0.75 {
					// 25%的概率弹出操作
					// 同时弹出，并判断弹出的数是否相同
					ans1 := stack1.Pop()
					ans2 := stack2.Pop()
					if ans1 != ans2 {
						t.Fatalf("出错了，弹出值不相等，栈1弹出%d,栈2弹出%d\n", ans1, ans2)
						return
					}
				} else {
					// 剩下25%的概率测试Empty()函数
					if stack1.IsEmpty() != stack1.IsEmpty() {
						t.Fatalf("出错了，是否相等，栈1%v,栈2%v\n", stack1.IsEmpty(), stack2.IsEmpty())
						return
					}
				}
			}
		}
	}
}
