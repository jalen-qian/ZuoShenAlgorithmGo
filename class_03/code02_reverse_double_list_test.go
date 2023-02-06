package main

import (
	"ZuoShenAlgorithmGo/utils"
	"fmt"
	"testing"
)

func TestReverseDoubleList(t *testing.T) {
	head := utils.GenerateRandomDoubleList(10, -100, 100)
	head2 := utils.CopyDoubleList(head)
	head3 := reverseDoubleList(head)
	fmt.Println(head, head2, head3)
}
