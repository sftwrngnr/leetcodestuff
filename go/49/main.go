package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	type Kmap struct {
		numoccur [26]int
	}
	gendict := func(instr string) *Kmap {
		rVal := new(Kmap)

		for _, c := range instr {
			rVal.numoccur[c-'a'] += 1
		}
		return rVal
	}

	myHmap := make(map[Kmap][]string)
	for _, a := range strs {
		myK := gendict(a)
		if myHmap[*myK] == nil {
			myHmap[*myK] = make([]string, 0)
		}
		myHmap[*myK] = append(myHmap[*myK], a)
	}
	retVal := make([][]string, 0)
	for a := range myHmap {
		retVal = append(retVal, myHmap[a])
	}
	return retVal
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("%v", groupAnagrams(strs))
}
