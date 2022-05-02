package main

import "fmt"

func main() {
	a := 2
	b := 10

	sum := getSum(a, b)

	fmt.Println(sum)
}

func getSum(a, b int) int {
	XOR := a ^ b
	AND := a & b

	if AND == 0 {
		return XOR
	}
	return getSum(XOR, AND<<1)
}
