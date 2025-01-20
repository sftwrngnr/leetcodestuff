package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildRoot(invals []int) *TreeNode {
	myTree := new(TreeNode)
	cNode := myTree
	for i := 0; i < len(invals)-1; i += 3 {
		cNode.Val = invals[i]
		if invals[i+1] != 0 {
			cNode.Left = new(TreeNode)
			cNode.Left.Val = invals[i+1]
		} else {
			cNode.Left = nil
		}
		if invals[i+2] != 0 {
			cNode.Right = new(TreeNode)
			cNode.Right.Val = invals[i+2]
		}
		fmt.Printf("%v %v %v\n", *cNode, cNode.Left, cNode.Right)
		cNode = cNode.Left
	}
	return myTree
}

//func maxPathSum(root *TreeNode) int {
//
//}

func main() {
	myVals := []int{-10, 9, 20, 0, 0, 15, 7}
	BuildRoot(myVals)
}
