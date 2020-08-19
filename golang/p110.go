package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, res := getDepth(root)
	return res
}

func getDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, ok1 := getDepth(root.Left)
	right, ok2 := getDepth(root.Right)
	if ok1 && ok2 && abs(left - right) <= 1 {
		return max(left, right) + 1, true
	} else {
		return 0, false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}