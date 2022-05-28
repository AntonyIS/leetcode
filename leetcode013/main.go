package main

import "fmt"

func main() {
	res := countBits(10)
	fmt.Println(res)
}

func countBits(n int) []int {
	ans, offset := make([]int, n+1), 1

	for i := 1; i <= n; i++ {
		if offset*2 == i {
			offset = i
		}
		ans[i] = 1 + ans[i-offset]

	}
	return ans
}
