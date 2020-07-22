package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var treeList []*TreeNode
func generateTrees(n int) []*TreeNode {
	treeList := []*TreeNode{}
	for i := 1; i <= n; i ++ {
		tree := &TreeNode{
			Val:   i,
			Left:  dgTree(1, i - 1),
			Right: dgTree(i + 1, n),
		}
		treeList = append(treeList, tree)
	}
	return treeList
}

func dgTree(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	for i := start; i <= end; i ++ {

	}
}