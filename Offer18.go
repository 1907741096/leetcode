package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	node := head
	if node.Val == val {
		return node.Next
	}
	for {
		if node.Next.Val == val {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
	return head
}