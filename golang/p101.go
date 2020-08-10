package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var leftArr []int
var rightArr []int
func isSymmetric(root *TreeNode) bool {
	leftArr = []int{}
	rightArr = []int{}
	leftTree(root)
	rightTree(root)
	length := len(leftArr)
	for i := 0; i < length; i ++ {
		if leftArr[i] != rightArr[i] {
			return false
		}
	}
	return true
}

func leftTree(node *TreeNode) {
	if node != nil {
		leftArr = append(leftArr, node.Val)
		if node.Left != nil || node.Right != nil {
			leftTree(node.Left)
			leftTree(node.Right)
		}
	} else {
		leftArr = append(leftArr, -1)
	}
}

func rightTree(node *TreeNode) {
	if node != nil {
		rightArr = append(rightArr, node.Val)
		if node.Left != nil || node.Right != nil {
			rightTree(node.Right)
			rightTree(node.Left)
		}
	} else {
		rightArr = append(rightArr, -1)
	}
}

