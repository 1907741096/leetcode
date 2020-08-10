package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var pathSumList [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	pathSumList = [][]int{}
	list := []int{}
	addPathSum(root, sum, list)
	return pathSumList
}

func addPathSum(root *TreeNode, sum int, list []int) {
	if root == nil {
		return
	}
	sum -= root.Val
	list = append(list, root.Val)
	l := make([]int, len(list))
	copy(list, l)
	if sum == 0 && root.Left == nil && root.Right == nil {
		pathSumList = append(pathSumList, l)
		return
	}
	addPathSum(root.Right, sum, l)
	addPathSum(root.Left, sum, l)
}