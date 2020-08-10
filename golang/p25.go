package main

type ListNode struct {
	Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var resultNode *ListNode
	var stack []int
	length := 0
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
		length ++
	}
	var num int
	num = length / k
	for i := length - 1; i >= num * k; i -- {
		resultNode = &ListNode{
			Val:  stack[i],
			Next: resultNode,
		}
	}
	for i := num - 1 ; i >= 0; i -- {
		for j := 0; j < k; j ++ {
			resultNode = &ListNode{
				Val:  stack[i * k + j],
				Next: resultNode,
			}
		}
	}
	return resultNode
}