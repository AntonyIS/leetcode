package main

import "fmt"

func main() {
	res := hammingWeight(10)
	fmt.Println(res)
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num&1 == 1 {
			res += 1
		}
		num = num >> 1

	}
	return res

}
