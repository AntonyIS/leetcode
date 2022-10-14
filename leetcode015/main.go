package main

import "fmt"

func main() {
	res := fib(2)
	fmt.Println(res)
}

func fib(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		temp := one
		one = one + two
		two = temp
	}

	return one
}
