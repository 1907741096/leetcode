package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var treePaths []string

func binaryTreePaths(root *TreeNode) []string {
	treePaths = []string{}
	if root != nil {
		dfsTreePaths(root, "")
	}
	return treePaths
}

func dfsTreePaths(node *TreeNode, path string) {
	if node.Left == nil && node.Right == nil {
		treePaths = append(treePaths, path + strconv.Itoa(node.Val))
	}
	if node.Left != nil {
		dfsTreePaths(node.Left, path + strconv.Itoa(node.Val) + "->")
	}
	if node.Right != nil  {
		dfsTreePaths(node.Right, path + strconv.Itoa(node.Val) + "->")
	}
}