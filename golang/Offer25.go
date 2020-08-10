package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var node = new(ListNode)
	head := node
	for {
		if l1 == nil {
			node.Val = l2.Val
			node.Next = l2.Next
			break
		} else if l2 == nil {
			node.Val = l1.Val
			node.Next = l1.Next
			break
		} else {
			if l1.Val < l2.Val {
				node.Val = l1.Val
				node.Next = new(ListNode)
				node = node.Next
				l1 = l1.Next
			} else {
				node.Val = l2.Val
				node.Next = new(ListNode)
				node = node.Next
				l2 = l2.Next
			}
		}
	}
	return head
}