package main

import "fmt"

func main() {
	var num uint32 = 3
	res := reverseBits(num)
	fmt.Println(res)
}

func reverseBits(num uint32) uint32 {
	// Define the results
	var res, i uint32
	for i < 32 {
		// Get the bit at i
		bit := (num >> i) & 1
		// Shift the bit to its righful place
		res = res | (bit << (31 - i))
		i += 1
	}
	return res
}
