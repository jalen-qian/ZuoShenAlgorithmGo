package main

import (
	"fmt"
	"testing"
)

func TestReverseDoubleList(t *testing.T) {
	head := GenerateRandomDoubleList(10, -100, 100)
	head2 := CopyDoubleList(head)
	head3 := reverseDoubleList(head)
	fmt.Println(head, head2, head3)
}
