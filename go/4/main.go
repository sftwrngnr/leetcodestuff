package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	bisectfind := func(inarr []int, fval int) int {
		var rval int
		var cend int = len(inarr)
		var cmid int = cend / 2
		fmt.Printf("%d, %d\n", inarr[cmid], fval)
		if inarr[cmid] > fval {
			fmt.Printf("middle is > input value. Bisecting to right\n")

		} else if inarr[cmid] < fval {
			fmt.Printf("middle is < input value. Bisecting to left.\n")
			if inarr[cmid+1] >= fval {
				fmt.Printf("Found entry point. %d\n", cmid)
			}
		} else if inarr[cmid] == fval {
			fmt.Printf("Middle == input value")
		}
		return rval
	}
	var rval float64
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	var midval = (l1 + l2) / 2
	var midcombined [2]int
	if (l1+l2)%2 == 0 {
		fmt.Printf("Even number of elements. %d\n", l1+l2)
		midcombined[0] = midval
		midcombined[1] = midval + 1
	} else {
		fmt.Printf("Odd number of elements %d\n", l1+l2)
		midcombined[0] = midval
		midcombined[1] = midval
	}

	//var inds int
	fmt.Printf("mid combined is %d\n", midcombined)
	fmt.Printf("%v, %v, %d, %d\n", nums1, nums2, l1, l2)
	if l1 > l2 {
		if nums1[l1-1] > nums2[0] {
			fmt.Printf("End of array 1 is > beginning of array 2\n")
		} else {
			fmt.Printf("End of array 1 is <= beginning of array 2")
		}
		bisectfind(nums1, nums2[0])
	} else {
		fmt.Printf("End of array 1 is <= beginning of array 2")

	}
	return rval
}

func main() {
	//                 4,6,7
	var l1 = []int{1, 3, 5, 5, 5, 7} //7
	//           5,8
	var l2 = []int{5, 6, 7, 11} //4
	//5[4]
	//5,5[5]
	//5,5,5[6]
	fmt.Printf("%f\n", findMedianSortedArrays(l1, l2))
}
