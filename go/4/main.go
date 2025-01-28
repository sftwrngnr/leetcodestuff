package main

import (
	"fmt"
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	bisectfind := func(inarr []int, fval int) (int, bool) {
		var rval int
		var rfound bool
		var high int = len(inarr) - 1
		var mid int
		low := 0
		for low <= high {
			mid = (low + high) / 2
			if inarr[mid] == fval {
				rfound = true
				rval = mid
				break
			}
			if inarr[mid] > fval {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		if !rfound {
			if inarr[mid] < fval {
				rval = len(inarr)
			} else {
				rval = mid
			}
		}
		return rval, rfound
	}

	var rval float64
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var midval = (l1 + l2) / 2
	var midcombined [2]int
	var cfidx int
	if (l1+l2)%2 == 0 {
		midcombined[0] = midval
		midcombined[1] = midval + 1
	} else {
		midcombined[0] = midval
		midcombined[1] = midval
	}
	hldArr := nums1

	//fmt.Printf("mid combined is %d\n", midcombined)
	//fmt.Printf("%v, %v, %d, %d\n", nums1, nums2, l1, l2)
	for (len(hldArr)-1 < midcombined[1]) && (cfidx < len(nums2)) {
		fmt.Printf("%v\n", hldArr)
		iloc, fnd := bisectfind(hldArr, nums2[cfidx])
		if iloc == len(hldArr) {
			hldArr = append(hldArr, nums2[cfidx])
		} else {
			hldArr = slices.Insert(hldArr, iloc, nums2[cfidx])
		}
		fmt.Printf("%d, %t\n", iloc, fnd)
		cfidx += 1
	}
	fmt.Printf("%v\n", hldArr)
	fmt.Printf("Midpoint values are %d, %d\n", hldArr[midcombined[0]], hldArr[midcombined[1]])
	if midcombined[1] > len(hldArr)-1 {
		rval = float64(hldArr[midcombined[0]])
	} else {
		rval = float64(hldArr[midcombined[0]]+hldArr[midcombined[1]]) / 2.0
	}
	return rval
}

func main() {
	//                 4,6,7
	var l1 = []int{0} //7
	//           5,8
	var l2 = []int{0} //4
	//5[4]
	//5,5[5]
	//5,5,5[6]
	fmt.Printf("%f\n", findMedianSortedArrays(l1, l2))
}
