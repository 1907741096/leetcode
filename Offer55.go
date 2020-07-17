package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	n := 0
	return depthTree(root, n)
}

func depthTree(node *TreeNode, n int) int {
	if node == nil {
		return n
	}
	n ++
	return max(depthTree(node.Left, n), depthTree(node.Right, n))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}