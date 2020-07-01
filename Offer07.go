package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var inHashList map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := &TreeNode{}
	head.Val = preorder[0]
	inHashList = make(map[int]int)
	for i, v := range inorder {
		inHashList[v] = i
	}
	head, _ = getTree(head, preorder[1:], -1)
	return head
}

func getTree(node *TreeNode, list []int, index int) (*TreeNode, []int) {
	if len(list) == 0 || index > 0 && index < inHashList[list[0]] {
		return node, list
	}
	if inHashList[node.Val] > inHashList[list[0]] {
		nowNode :=  &TreeNode{
			Val: list[0],
			Left: nil,
			Right: nil,
		}
		node.Left, list = getTree(nowNode, list[1:], inHashList[node.Val])
	}
	if len(list) == 0 || index > 0 && index < inHashList[list[0]] {
		return node, list
	}
	if inHashList[node.Val] < inHashList[list[0]] {
		nowNode := &TreeNode{
			Val:   list[0],
			Left:  nil,
			Right: nil,
		}
		node.Right, list = getTree(nowNode, list[1:], index)
	}
	return node, list
}