package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var queue []nodeAndTier
	queue = []nodeAndTier{}
	queue = append(queue, nodeAndTier{
		node:root,
		tier:0,
	})
	var arr [][]int
	arr = [][]int{}
	for len(queue) != 0 && queue[0].node != nil {
		queueTop := queue[0]
		if queueTop.node.Left != nil {
			queue = append(queue, nodeAndTier{
				node: queueTop.node.Left,
				tier: queueTop.tier + 1,
			})
		}
		if queueTop.node.Right != nil {
			queue = append(queue, nodeAndTier{
				node: queueTop.node.Right,
				tier: queueTop.tier + 1,
			})
		}
		queue = queue[1:]
		if len(arr) <= queueTop.tier {
			arr = append(arr, []int{})
		}
		if queueTop.tier % 2 == 0 {
			arr[queueTop.tier] = append(arr[queueTop.tier], queueTop.node.Val)
		} else {
			arr[queueTop.tier] = append([]int{queueTop.node.Val}, arr[queueTop.tier]...)
		}
	}
	return arr
}

type nodeAndTier struct {
	node *TreeNode
	tier int
}