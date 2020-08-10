package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA := headA
	nodeB := headB
	for nodeA != nodeB {
		if nodeA == nil && nodeB == nil {
			return nil
		}

		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}

		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}