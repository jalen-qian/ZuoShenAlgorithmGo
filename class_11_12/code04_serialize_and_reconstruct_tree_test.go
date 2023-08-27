package class_11_12

import (
	"fmt"
	"testing"
)

// 序列化和反序列化二叉树测试

func TestSerializeAndReconstructBT_PreSerialize(t *testing.T) {
	// 测试先序序列化
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
	}
	//root = nil
	t.Log("打印二叉树：")
	PrintBT(root)
	ser := &SerializeAndReconstructBT{}
	serStr := ser.PreSerialize(root)
	t.Logf("先序方式序列化：%s \n", serStr)
	if serStr != "1,2,4,#,#,5,#,#,3,6,#,#,7,#,#" {
		t.Error("先序方式序列化不正确！！！")
	}

	inStr := ser.InSerialize(root)
	t.Logf("中序方式序列化：%s \n", inStr)
	if inStr != "#,4,#,2,#,5,#,1,#,6,#,3,#,7,#" {
		t.Error("中序方式序列化不正确！！！")
	}

	posStr := ser.PosSerialize(root)
	t.Logf("后序方式序列化：%s \n", posStr)
	if posStr != "#,#,4,#,#,5,2,#,#,6,#,#,7,3,1" {
		t.Error("后序方式序列化不正确！！！")
	}

	levelStr := ser.LevelSerialize(root)
	t.Logf("按层遍历方式序列化：%s \n", levelStr)
	if levelStr != "1,2,3,4,5,6,7,#,#,#,#,#,#,#,#" {
		t.Error("按层方式序列化不正确！！！")
	}

	fmt.Println("===========反序列化==============")
	newRoot := ser.BuildByPreSerialize("1,2,4,#,#,5,#,#,3,6,#,#,7,#,#")
	fmt.Println("先序遍历的方式反序列化：")
	PrintBT(newRoot)

	newRoot = ser.BuildByPosSerialize("#,#,4,#,#,5,2,#,#,6,#,#,7,3,1")
	fmt.Println("后序遍历的方式反序列化：")
	PrintBT(newRoot)

	newRoot = ser.BuildByLevelSerialize("1,2,3,4,5,6,7,#,#,#,#,#,#,#,#")
	//newRoot = ser.BuildByLevelSerialize("25,#,47,61,#,#,#")
	fmt.Println("按层遍历的方式反序列化：")
	PrintBT(newRoot)

}

// 使用对数器测试全部功能
func TestSerializeAndReconstructBT(t *testing.T) {
	t.Log("开始测试...")
	ser := &SerializeAndReconstructBT{}
	for i := 0; i < 100000; i++ {
		root := generateRandomBT(0, 100, 10)
		// 分别用3种方式序列化
		preSer := ser.PreSerialize(root)
		posSer := ser.PosSerialize(root)
		levelSer := ser.LevelSerialize(root)

		// 再分别用各自的方式反序列化重建
		rootPre := ser.BuildByPreSerialize(preSer)
		rootPos := ser.BuildByPosSerialize(posSer)
		rootLevel := ser.BuildByLevelSerialize(levelSer)

		// 比较重建后的树是否完全一致，且与原始树一致
		if !isBTEqual(rootPre, root) || !isBTEqual(rootPos, root) || !isBTEqual(rootLevel, root) {
			t.Errorf("测试失败，反序列化后的值不相等\n")
			//PrintBT(rootPre)
			//PrintBT(rootPos)
			//PrintBT(rootLevel)
			return
		}
	}
	t.Log("测试成功！！！")
}
