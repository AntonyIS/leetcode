package main

import (
	"strconv"
)

func main() {
	hammingWeight(123)
}

func hammingWeight(num uint32) int {
	res := 0
	s := strconv.FormatUint(uint64(num), 2)
	for _, v := range s {
		if v == '1' {
			res++
		}
	}
	return res

}
