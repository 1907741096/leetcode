package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sortArr []int
func kthLargest(root *TreeNode, k int) int {
	sortArr = []int{}
	midTree(root, k)
	return sortArr[k - 1]
}

func midTree(node *TreeNode, k int) {
	if node == nil || len(sortArr) >= k {
		return
	}
	midTree(node.Right, k)
	sortArr = append(sortArr, node.Val)
	midTree(node.Left, k)
}