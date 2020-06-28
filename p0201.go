package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hashSet := make(map[int]int)
	hashSet[head.Val] = 0
	node := head
	for node.Next != nil {
		if _, ok := hashSet[node.Next.Val]; ok {
			node.Next = node.Next.Next
		} else {
			node = node.Next
			hashSet[node.Val] = 0
		}
	}
	node.Next = nil
	return head
}
