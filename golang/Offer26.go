package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	return checkStructure(A, B, false)
}

func checkStructure(A *TreeNode, B *TreeNode, flag bool) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val && checkStructure(A.Left, B.Left, true) && checkStructure(A.Right, B.Right, true) {
		return true
	}
	if flag {
		return false
	}
	return checkStructure(A.Left, B, false) || checkStructure(A.Right, B, false)
}