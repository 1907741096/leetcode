package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return mergeTwoLists(lists[0], nil)
	} else if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else {
		mid := len(lists) / 2
		return mergeTwoLists(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:]))
	}
}

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
