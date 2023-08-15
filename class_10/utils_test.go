package class_10

import "testing"

func TestPrintBT(t *testing.T) {
	root := &Node{
		Value: 1,
		Left:  &Node{Value: 2, Left: &Node{Value: 4}, Right: &Node{Value: 5}},
		Right: &Node{Value: 3, Left: &Node{Value: 6}, Right: &Node{Value: 7}},
	}
	PrintBT(root)
}
