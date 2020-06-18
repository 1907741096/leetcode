package main

import (
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverFromPreorder(S string) *TreeNode {
	strArr := []byte(S)
	length := len(strArr)
	root := new(TreeNode)
	str := ""
	var i int
	for i = 0; i < length; i ++ {
		if strArr[i] != '-' {
			str += string(strArr[i])
		} else {
			break
		}
	}
	root.Val, _ = strconv.Atoi(str)
	str = ""
	lastCen := 0
	cen := 0
	stack := []*TreeNode{root}
	for ; i < length; i ++ {
		if strArr[i] == '-' {
			cen ++
		} else {
			if i == length - 1 || strArr[i + 1] == '-' {
				str += string(strArr[i])
				if lastCen + 1 == cen {
					node := stack[len(stack) - 1]
					newNode := new(TreeNode)
					num, _ := strconv.Atoi(str)
					newNode.Val = num
					node.Left = newNode

					stack = append(stack, newNode)
					lastCen = cen
					cen = 0
				} else {
					node := stack[len(stack) - lastCen + cen - 2]
					stack = stack[0: len(stack) - lastCen + cen - 1]

					newNode := new(TreeNode)
					num, _ := strconv.Atoi(str)
					newNode.Val = num
					node.Right = newNode

					stack = append(stack, newNode)
					lastCen = cen
					cen = 0
				}
				str = ""
			} else {
				str += string(strArr[i])
			}
		}
	}
	return root
}