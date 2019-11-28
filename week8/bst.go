package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var sum int = 0

// To Greater Sum Tree
func ToGst(root *TreeNode) {
	if root != nil {
		ToGst(root.Right)
		sum += root.Val
		root.Val = sum
		fmt.Printf("Node:%v\n",root)
		ToGst(root.Left)
	}
}

// 中序遍历
func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("Node:%v\n", root)
		InOrder(root.Right)
	}
}

// 前序遍历
func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Printf("Node:%v\n", root)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}

}

// Insert
func Insert(root *TreeNode, data int) {
	if root == nil {
		root = &TreeNode{
			Val: data,
		}
		return
	}
	currentNode := root
	for ;currentNode != nil; {
		if data > currentNode.Val {
			if currentNode.Right == nil {
				currentNode.Right = &TreeNode{
					Val: data,
				}
				return
			}
			currentNode = currentNode.Right
		} else {
			if currentNode.Left == nil {
				currentNode.Left = &TreeNode{
					Val: data,
				}
				return
			}
			currentNode = currentNode.Left
		}
	}
}

// Search
func Find(root *TreeNode, data int) *TreeNode {
	currentNode := root
	for ;currentNode!= nil; {
		fmt.Printf("currentNode:%v\n", currentNode)
		if data < currentNode.Val {
			currentNode = currentNode.Left
		} else if data > currentNode.Val {
			currentNode = currentNode.Right
		} else {
			return currentNode
		}
	}

	return nil
}

func (n *TreeNode) String() string {
	str := fmt.Sprintf("{val=%d}", n.Val)
	return str
}