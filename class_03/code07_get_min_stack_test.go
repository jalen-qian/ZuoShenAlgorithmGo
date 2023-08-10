package class_03

import (
	"fmt"
	"testing"
)

func TestGetMinStack(t *testing.T) {
	stack := NewGetMinStack[int]()
	stack.Push(10)
	stack.Push(5)
	stack.Push(8)
	stack.Push(7)
	stack.Push(12)
	fmt.Println("最小值：", stack.GetMin())
	fmt.Println("弹出：", stack.Pop())
	stack.Push(4)
	stack.Push(20)
	for !stack.IsEmpty() {
		fmt.Println("最小值：", stack.GetMin())
		fmt.Println("弹出：", stack.Pop())
	}

}
