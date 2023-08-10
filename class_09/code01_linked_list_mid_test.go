package class_09

import (
	"testing"

	"ZuoShenAlgorithmGo/utils"
)

// 测试1：输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func TestMidOrUpMidNode(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 100000; i++ {
		// 生成一个随机链表和伴随数组，且里面的值不可能是负数
		head, arr := generateRandomLinkListAndArr(1000, 0, 200)
		midOrUpMidNode := MidOrUpMidNode(head)
		midOrUpMidInArr := MidOrUpMidNodeArr(arr)
		// 比较
		if head == nil {
			if midOrUpMidInArr != -1 {
				t.Errorf("Fucking fucked!! 链表中点为空，数组中点不是-1")
				return
			}
		} else if midOrUpMidNode.Value != midOrUpMidInArr {
			t.Errorf("Fucking fucked!! 中点不一致，链表中点是：%d, 数组中点是：%d", midOrUpMidNode.Value, midOrUpMidInArr)
			return
		}
	}
	t.Log("Great!!!")
}

// 测试2：输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func TestMidOrDownMidNode(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 100000; i++ {
		// 生成一个随机链表和伴随数组，且里面的值不可能是负数
		head, arr := generateRandomLinkListAndArr(1000, 0, 200)
		midOrDownMidNode := MidOrDownMidNode(head)
		midOrDownMidInArr := MidOrDownMidNodeArr(arr)
		// 比较
		if head == nil {
			if midOrDownMidInArr != -1 {
				t.Errorf("Fucking fucked!! 链表中点为空，数组中点不是-1")
				return
			}
		} else if midOrDownMidNode.Value != midOrDownMidInArr {
			t.Errorf("Fucking fucked!! 中点不一致，链表中点是：%d, 数组中点是：%d", midOrDownMidNode.Value, midOrDownMidInArr)
			return
		}
	}
	t.Log("Great!!!")
}

// 测试3：输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func TestMidOrUpMidPreNode(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 100000; i++ {
		// 生成一个随机链表和伴随数组，且里面的值不可能是负数
		head, arr := generateRandomLinkListAndArr(1000, 0, 200)
		newHead := MidOrUpMidPreNode(head)
		mid := MidOrUpMidNodePreArr(arr)
		// 比较
		if newHead == nil {
			if mid != -1 {
				t.Errorf("Fucking fucked!! 链表中点前一个为空，数组中点前一个不是-1")
				return
			}
		} else if newHead.Value != mid {
			t.Errorf("Fucking fucked!! 中点不一致，链表中点前一个是：%d, 数组中点前一个是：%d", newHead.Value, mid)
			return
		}
	}
	t.Log("Great!!!")
}

// 测试4：输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func TestMidOrDownMidPreNode(t *testing.T) {
	t.Log("Start test...")
	for i := 0; i < 100000; i++ {
		// 生成一个随机链表和伴随数组，且里面的值不可能是负数
		head, arr := generateRandomLinkListAndArr(1000, 0, 200)
		newHead := MidOrDownMidPreNode(head)
		mid := MidOrDownMidPreArr(arr)
		// 比较
		if newHead == nil {
			if mid != -1 {
				t.Errorf("Fucking fucked!! 链表中点前一个为空，数组中点前一个不是-1")
				return
			}
		} else if newHead.Value != mid {
			t.Errorf("Fucking fucked!! 中点不一致，链表中点前一个是：%d, 数组中点前一个是：%d", newHead.Value, mid)
			return
		}
	}
	t.Log("Great!!!")
}

// 生成一个随机链表和伴随数组，伴随数组的值顺序和链表完全一致
// 如果链表头节点为空，则伴随数组的长度也为0
func generateRandomLinkListAndArr(maxLen, minNum, maxNum int, customizedValues ...int) (*Node, []int) {
	arr := utils.GenerateRandomSlice(maxLen, minNum, maxNum, customizedValues...)
	// 基于这个数组生成链表
	var head *Node
	var cur *Node
	if len(arr) > 0 {
		head = &Node{Value: arr[0]}
		cur = head
	}
	// 从1开始往后串
	for i := 1; i < len(arr); i++ {
		cur.Next = &Node{Value: arr[i]}
		cur = cur.Next
	}
	return head, arr
}

// 如果是奇数个，返回中点，偶数个，返回上中点，如果数组为空，返回-1
func MidOrUpMidNodeArr(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	// 找到上中点
	var mid int
	if len(arr)%2 == 0 {
		mid = len(arr)/2 - 1
	} else {
		mid = len(arr) / 2
	}
	return arr[mid]
}

// 如果是奇数个，返回中点，偶数个，返回下中点，如果数组为空，返回-1
func MidOrDownMidNodeArr(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	// 找到下中点，默认长度/2算出来就是下中点
	return arr[len(arr)/2]
}

// 如果是奇数个，返回中点前一个，偶数个，返回上中点前一个，如果找不到，返回-1
func MidOrUpMidNodePreArr(arr []int) int {
	if len(arr) <= 2 {
		return -1
	}
	// 找到上中点
	var mid int
	if len(arr)%2 == 0 {
		mid = len(arr)/2 - 1
	} else {
		mid = len(arr) / 2
	}
	// 返回上中点前一个数
	return arr[mid-1]
}

// 如果是奇数个，返回中点前一个，偶数个，返回下中点前一个，如果找不到，返回-1
func MidOrDownMidPreArr(arr []int) int {
	// 如果少于两个，则不可能有中点前一个，返回-1
	if len(arr) <= 1 {
		return -1
	}
	// 返回上中点前一个数
	return arr[len(arr)/2-1]
}
