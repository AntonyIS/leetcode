package main

import "fmt"

func main() {
	a, b := 2, 2
	sum := getSum(a, b)

	fmt.Println(sum)
}

func getSum(a, b int) int {
	AND := a & b
	XOR := a ^ b
	if AND == 0 {
		return XOR
	}
	return getSum(XOR, AND<<1)
}
