package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const inf = math.MaxInt32 / 2 // 或 math.MaxInt64 / 2

func minCameraCover(root *TreeNode) int {
	var dfs func(*TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1					// 当前节点放置监控，摄像头数目
		b = min(a, min(la+rb, ra+lb))	// 整棵树需要的摄像头数目，无论节点本身是否被监控到
		c = min(a, lb+rb)				// 两棵子树需要的摄像头数目，无论节点本身是否被监控到
		return
	}
	_, ans, _ := dfs(root)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}