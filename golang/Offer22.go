package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	length := 0
	node := head
	for node != nil {
		length ++
		node = node.Next
	}
	length -= k
	for i := 0; i < length; i ++ {
		head = head.Next
	}
	return head
}