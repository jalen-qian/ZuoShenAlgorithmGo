package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"time"
)

// 题目2：删除链表中给定的数
func deleteNumberInLinkedList(head *utils.Node, number int) *utils.Node {
	if head == nil {
		return nil
	}
	// 1. 找到第一个不等于number的节点作为头结点，并将前面的节点都释放
	cur := head
	for head != nil && head.Value == number {
		// head跳下一个
		head = head.Next
		cur.Next = nil
		cur = head
	}
	if head == nil {
		return nil
	}
	// 2. 接下来的链表中，如果遇到number，则删除
	pre := head
	cur = head.Next
	for cur != nil {
		if cur.Value == number {
			cur = cur.Next
			continue
		}
		// 找到了不是number的节点
		// pre的下一个直接指向这个节点
		pre.Next = cur
		// pre来到当前位置
		pre = cur
		// 往下一个条
		cur = cur.Next
	}
	return head
}

// testDeleteNumberInLinkedList 对数器，分配额外的空间来实现
func testDeleteNumberInLinkedList(head *utils.Node, number int) *utils.Node {
	// 1. 生成一个节点数组
	nodeArr := make([]*utils.Node, 0)
	// 2. 将链表中的节点填入数组中
	cur := head
	for cur != nil {
		nodeArr = append(nodeArr, cur)
		cur = cur.Next
	}
	// 3. 遍历数组，将是number的节点置为空
	for i, node := range nodeArr {
		if node.Value == number {
			nodeArr[i] = nil
		}
	}
	// 4. 将节点重新连接起来
	var foundHead bool
	var pre *utils.Node
	nilCount := 0
	for _, node := range nodeArr {
		// 值为空了，说明是number被删除掉了，则跳过
		if node == nil {
			nilCount++
			continue
		}
		// 如果还没指定过头节点，则指定，并将pre设定到此位置
		if !foundHead {
			head = node
			pre = head
			foundHead = true
		} else {
			// 往下遍历，遇到的节点都置为pre的next
			pre.Next = node
			pre = node
		}
	}
	// 全删除了，则返回nil
	if nilCount == len(nodeArr) {
		return nil
	}
	return head
}

func main() {
	fmt.Println("开始测试")
	// 使用特定例子测试
	head := utils.GenerateLinkedListBySlice([]int{2, 2, 2, 3, 4, 5, 2, 3})
	ans1 := testDeleteNumberInLinkedList(head, 2)
	fmt.Println(utils.SPrintLinkedList(ans1))
	// 测试，使用对数器测试一百万次
	testTimes := 1000000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		// 初始化一个随机链表
		head1 := utils.GenerateRandomLinkedList(100, -20, 20)
		head2 := utils.CopyLinkedList(head1)
		originList := utils.CopyLinkedList(head1)
		autoDelete := utils.GenNodeValue(-20, 20)
		head1 = deleteNumberInLinkedList(head1, autoDelete)
		head2 = testDeleteNumberInLinkedList(head2, autoDelete)
		fmt.Printf("第%d次测试\n", i+1)
		if !utils.CheckLinkedListEqual(head1, head2) {
			fmt.Printf("出错了\n原始链表：%s\nhead1:%s\nhead2:%s\n要删除的数:%d\n",
				utils.SPrintLinkedList(originList),
				utils.SPrintLinkedList(head1),
				utils.SPrintLinkedList(head2),
				autoDelete,
			)
			return
		}
	}
	fmt.Println("结束测试")
}
