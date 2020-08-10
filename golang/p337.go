package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	return max(dfs(root))
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l1, l2 := dfs(root.Left)
	r1, r2 := dfs(root.Right)
	lMax, rMax := max(l1, l2), max(r1, r2)
	return root.Val + l2 + r2, lMax + rMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}