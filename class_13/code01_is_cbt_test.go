package class_13

import "testing"

// 判断一个二叉树是否是完全二叉树

func TestIsCBT(t *testing.T) {
	//
	root1 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 6},
		},
	}
	if !IsCBT(root1) {
		t.Error("出错了，当前树是完全二叉树")
	}

	root2 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 6},
		},
	}
	if IsCBT(root2) {
		t.Error("出错了，当前树不是完全二叉树")
	}

	t.Log("测试通过！！！")

}
