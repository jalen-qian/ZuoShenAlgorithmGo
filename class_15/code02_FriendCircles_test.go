package class_15

import "testing"

func TestFriendCircles(t *testing.T) {
	arr := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
	ans := findCircleNum(arr)
	t.Log(ans)
}
