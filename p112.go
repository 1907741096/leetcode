package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return checkPathSum(root, sum)
}

func checkPathSum(root *TreeNode, sum int) bool {
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	} else if root.Left == nil && root.Right == nil && sum != 0 {
		return false
	} else {
		if root.Left == nil {
			return checkPathSum(root.Right, sum)
		} else if root.Right == nil {
			return checkPathSum(root.Left, sum)
		} else {
			return checkPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
		}
	}
}