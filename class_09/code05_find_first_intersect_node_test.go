package class_09

import "testing"

type testCase struct {
	head1 *Node
	head2 *Node
	want  *Node
	name  string
}

func TestGetIntersectNode(t *testing.T) {
	cases := []testCase{
		generateTestCase(
			"测试用例1：两个无环，也不相交",
			[]int{1, 2, 3, 4, 5, 6, 7, 8}, false, 0,
			[]int{101, 102, 103, 104, 105, 106}, false, 0,
			-1,
		),
		generateTestCase(
			"测试用例2：两个无环，但是相交",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false, 0,
			[]int{11, 12, 13, 14, 15, 4, 5, 6, 7, 8, 9, 10}, false, 0,
			4,
		),
		generateTestCase(
			"测试用例3：head1有环，head2无环",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, true, 4,
			[]int{20, 21, 22, 23, 24, 25, 26, 27, 28}, false, 0,
			-1,
		),
		generateTestCase(
			"测试用例4：两个都有环，且相交于入环前的点",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, true, 4,
			[]int{20, 21, 22, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, true, 4,
			3,
		),
		generateTestCase(
			"测试用例5：两个都有环，且相交于环内的两个点",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, true, 4,
			[]int{20, 21, 22, 24, 25, 6, 7, 8, 9, 10, 11, 12, 13, 4, 5}, true, 6,
			4,
		),
	}
	for _, tc := range cases {
		got := GetIntersectNode(tc.head1, tc.head2)
		if got != tc.want {
			t.Errorf("Fucking Fucked!! %sfailed, want:%v, got:%v\n", tc.name, tc.want, got)
			return
		}
	}
	t.Log("Great!!!")
}

func generateTestCase(
	name string,
	arr1 []int,
	isArr1Ring bool,
	arr1RingValue int,
	arr2 []int,
	isArr2Ring bool,
	arr2RingValue int,
	wantIndex int,
) testCase {
	generateNodeMap := make(map[int]*Node)
	head1 := generateRingLinkedList(arr1, isArr1Ring, arr1RingValue, generateNodeMap)
	head2 := generateRingLinkedList(arr2, isArr2Ring, arr2RingValue, generateNodeMap)
	var want *Node
	if wantIndex >= 0 {
		want = generateNodeMap[wantIndex]
	}
	return testCase{
		head1: head1,
		head2: head2,
		want:  want,
		name:  name,
	}
}

// 根据数组生成链表，如果数组中的值在全局数组中已存在，就会复用这个节点
//
//	arr 数组
//	isRing 是否是环
//	ringValue 入环节点
//	要求：arr数组必须是不重复的，并且如果要求是环， 则入环的值必须在arr中存在
func generateRingLinkedList(arr []int, isRing bool, ringValue int, generateNodeMap map[int]*Node) *Node {
	if len(arr) == 0 {
		return nil
	}
	for _, val := range arr {
		n, ok := generateNodeMap[val]
		if !ok {
			n = &Node{Value: val}
			generateNodeMap[val] = n
		}
	}
	// 一个个串起来
	for i := 0; i < len(arr)-1; i++ {
		generateNodeMap[arr[i]].Next = generateNodeMap[arr[i+1]]
	}
	head := generateNodeMap[arr[0]]
	// 最后，如果是环，则将最后一个节点的Next指向入环节点
	if isRing {
		n, ok := generateNodeMap[ringValue]
		if !ok {
			return head
		}
		// 最后一个节点的Next指向n
		generateNodeMap[arr[len(arr)-1]].Next = n
	}
	return head
}
