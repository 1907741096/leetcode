package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	list := []int{}
	for ; head != nil; head = head.Next {
		list = append(list, head.Val)
	}
	return getListToBST(list)
}

func getListToBST(list []int) *TreeNode {
	if len(list) == 0 {
		return nil
	}
	if len(list) == 1 {
		return &TreeNode{
			Val: list[0],
		}
	}
	mid := len(list) / 2
	return &TreeNode{
		Val: list[mid],
		Left: getListToBST(list[:mid]),
		Right: getListToBST(list[mid + 1:]),
	}
}