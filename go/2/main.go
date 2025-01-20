package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rVal := &ListNode{}
	hPtr := rVal
	crPtr := rVal
	var carry int
	cn1 := l1
	cn2 := l2
	done := false
	for done != true {
		if (cn1 == nil) && (cn2 == nil) {
			done = true
			if carry != 0 {
				hPtr.Next.Val = carry
			} else {
				hPtr.Next = nil
			}
			continue
		}
		hPtr = crPtr
		tval := carry
		fmt.Printf("Carry is %d, cn1 is %v, cn2 is %v\n", carry, cn1, cn2)
		if cn1 != nil {
			tval += cn1.Val
		}
		if cn2 != nil {
			tval += cn2.Val
		}
		carry = 0
		if tval > 9 {
			carry = 1
			tval -= 10
		}
		crPtr.Val = tval
		crPtr.Next = &ListNode{}
		crPtr = crPtr.Next
		if cn1 != nil {
			cn1 = cn1.Next
		}
		if cn2 != nil {
			cn2 = cn2.Next
		}

	}
	return rVal
}

func printem(ln *ListNode) []int {
	cn := ln
	rval := make([]int, 0)
	for cn != nil {
		rval = append(rval, cn.Val)
		cn = cn.Next
	}
	return rval
}

func main() {
	genList := func(invals []int) *ListNode {
		var rval *ListNode = new(ListNode)
		var cnode *ListNode = rval
		for i, v := range invals {
			if i == 0 {
				cnode.Val = v
				continue
			}
			cnode.Next = &ListNode{Val: v, Next: nil}
			cnode = cnode.Next

		}
		return rval
	}

	l1 := genList([]int{9, 9, 9, 9, 9, 9, 9})
	printem(l1)
	l2 := genList([]int{9, 9, 9, 9})
	printem(l2)

	fmt.Printf("%v\n", printem(addTwoNumbers(l1, l2)))

}
