package main

import "fmt"

func main() {
	a, b := 2, 43
	sum := GetSum(a, b)

	fmt.Println(sum)
}

func GetSum(a, b int) int {
	AND := a & b
	XOR := a ^ b
	if AND == 0 {
		return XOR
	}
	return GetSum(XOR, AND<<1)
}
