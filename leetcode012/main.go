package main

import "fmt"

func main() {
	res := hammingWeight(1111111111)
	fmt.Println(res)

	var t uint32 = 2
	fmt.Println(t % 2)
}

func hammingWeight(num uint32) int {
	var bits int
	for num != 0 {
		num = num & (num - 1)
		bits += 1
	}
	return bits
}
func hammingWeight2(num uint32) int {
	var bits uint32

	for num != 0 {
		bits += num % 2
		num = num >> 1
	}

	return int(bits)
}
