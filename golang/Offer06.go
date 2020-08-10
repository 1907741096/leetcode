package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	list := []int{}
	for head != nil {
		list = append([]int{head.Val}, list...)
		head = head.Next
	}
	return list
}