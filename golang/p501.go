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
	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}
		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}
	return res
}