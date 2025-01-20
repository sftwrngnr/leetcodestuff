package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenTree(input []*int) *TreeNode {
	var head *TreeNode = nil
	var cnode *TreeNode = head
	genleaf := func(inval int) *TreeNode {
		rval := new(TreeNode)
		rval.Val = inval
		return rval
	}
	for v := 0; v < len(input); v += 3 {
		if head == nil {
			head = new(TreeNode)
			cnode = head
		}
		cnode.Val = *input[v]
		if input[v+1] != nil {
			cnode.Left = genleaf(*input[v+1])
		}
		if input[v+2] != nil {
			cnode.Right = genleaf(*input[v+1])
		}
		if cnode.Left != nil {
			cnode = cnode.Left
		} else if cnode.Right != nil {
			cnode = cnode.Right
		}

	}
	fmt.Printf("%v\n", head)
	return head
}

func main() {
	ival := func(i int) *int { return &i }
	rList := []*int{ival(3), ival(9), ival(20), nil, nil, ival(15), ival(7)}
	mytree := GenTree(rList)
	fmt.Printf("%v\n", mytree)

}
