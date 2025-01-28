package main

import (
	"fmt"
	"strconv"
)

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

func testBinaryAdd(ln1 *ListNode, ln2 *ListNode) []int {
	lookup := []int64{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	var myVal1 int64 //Can handle up to 128 nybbles
	var myVal2 int64
	var rVal []int
	var nshft int
	cptr1 := ln1
	cptr2 := ln2
	for {
		myVal1 = myVal1 << 4
		myVal2 = myVal2 << 4
		nshft++
		//fmt.Printf("%s\n", strconv.FormatInt(myVal1, 2))
		//fmt.Printf("%s\n", strconv.FormatInt(myVal2, 2))
		if cptr1 != nil {
			myVal1 |= int64(cptr1.Val)
			cptr1 = cptr1.Next
		}
		if cptr2 != nil {
			myVal2 |= int64(cptr2.Val)
			cptr2 = cptr2.Next
		}
		if (cptr1 == nil) && (cptr2 == nil) {
			break
		}
	}
	smask := int64(0x0F) << ((nshft) * 4)
	fmt.Printf("%s\n", strconv.FormatInt(myVal1, 2))
	fmt.Printf("%s\n", strconv.FormatInt(myVal2, 2))
	cryflg := false
	for i := 0; i < nshft+1; i++ {
		v1 := (myVal1 & smask) >> ((nshft - i) * 4)
		v2 := (myVal2 & smask) >> ((nshft - i) * 4)
		tidx := v1 + v2
		if cryflg {
			tidx += 1
		}
		tsum := lookup[tidx]
		cryflg = tidx > 9
		fmt.Printf("%d\n", i)
		//fmt.Printf("%s\n", strconv.FormatInt(smask, 2))
		fmt.Printf("%s\n", strconv.FormatInt(v1, 2))
		fmt.Printf("%s\n", strconv.FormatInt(v2, 2))
		fmt.Printf("%s\n", strconv.FormatInt(tsum, 10))
		smask = smask >> 4
	}
	if cryflg {
		fmt.Printf("Extra carry\n")
	}
	fmt.Printf("\n%v\n%v\n", myVal1, myVal2)
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
	l2 := genList([]int{9, 8, 7, 6})
	printem(l1)
	printem(l2)
	testBinaryAdd(l1, l2)

	fmt.Printf("%v\n", printem(addTwoNumbers(l1, l2)))

}
