package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	num, count, maxCount := -1, 0, 0
	var update func(int)
	update = func(n int) {
		if num == n {
			count ++
		} else {
			count = 1
		}

		if count > maxCount {
			maxCount = count
			res = []int{n}
		} else if count == maxCount {
			res = append(res, n)
		}
		num = n
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		nowQueue := queue[0]
		queue = queue[1:]
		update(nowQueue.Val)
		if nowQueue.Left != nil {
			if nowQueue.Val == nowQueue.Left.Val {
				queue = append([]*TreeNode{nowQueue.Left}, queue...)
			} else {
				queue = append(queue, nowQueue.Left)
			}
		}
		if nowQueue.Right != nil {
			if nowQueue.Val == nowQueue.Right.Val {
				queue = append([]*TreeNode{nowQueue.Right}, queue...)
			} else {
				queue = append(queue, nowQueue.Right)
			}
		}
	}
	return res
}