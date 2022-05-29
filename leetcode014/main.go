package main

import "fmt"

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	XOR_1 := 0
	XOR_2 := 0

	for _, v := range nums {
		XOR_1 ^= v
	}
	for i := 0; i <= len(nums); i++ {
		XOR_2 ^= i
	}

	return XOR_1 ^ XOR_2
}
