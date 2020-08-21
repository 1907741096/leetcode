package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var depth int

func minDepth(root *TreeNode) int {
	depth = 0
	getDepth(root, 1)
	return depth
}

func getDepth(root *TreeNode, nowDepth int) {
	if depth != 0 && nowDepth > depth || root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		depth = nowDepth
		return
	}
	getDepth(root.Left, nowDepth + 1)
	getDepth(root.Right, nowDepth + 1)
}