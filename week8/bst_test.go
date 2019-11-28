package main

import (
	"testing"
)

func TestInsertAndFind(t *testing.T) {
	root := &TreeNode{
		Val: 4,
	}

	root.Left = &TreeNode{
		Val: 1,
	}
	root.Left.Left = &TreeNode{
		Val: 0,
	}
	root.Left.Right = &TreeNode{
		Val: 2,
	}
	root.Left.Right.Right = &TreeNode{
		Val: 3,
	}


	root.Right = &TreeNode{
		Val: 6,
	}
	root.Right.Left = &TreeNode{
		Val: 5,
	}
	root.Right.Right = &TreeNode{
		Val: 7,
	}
	root.Right.Right.Right = &TreeNode{
		Val: 8,
	}

	ToGst(root)
}
