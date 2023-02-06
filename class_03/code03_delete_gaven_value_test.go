package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDeleteGavenValue(t *testing.T) {
	fmt.Println("开始测试")
	// 使用特定例子测试
	head := utils.GenerateLinkedListBySlice([]int{2, 2, 2, 3, 4, 5, 2, 3})
	ans1 := testDeleteNumberInLinkedList(head, 2)
	fmt.Println(utils.SPrintLinkedList(ans1))
	// 测试，使用对数器测试一百万次
	testTimes := 200000
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
