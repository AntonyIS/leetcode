package main

import "fmt"

func main() {
	res := hammingWeight(10)
	fmt.Println(res)
}

func hammingWeight(num uint32) int {
	var res int
	n := int(num)
	for n > 0 {
		res += n & 1
		n = n >> 1
	}

	return res

}
