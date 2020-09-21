package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	rightConvert(root, 0)
	return root
}

func rightConvert(root *TreeNode, addNum int) int {
	if root == nil {
		return 0
	}
	rightNum := rightConvert(root.Right, addNum)
	root.Val += rightNum + addNum
	leftNum := rightConvert(root.Left, root.Val)
	return leftNum + root.Val - addNum
}