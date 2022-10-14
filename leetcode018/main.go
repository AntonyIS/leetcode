package main

import "fmt"

func main() {
	nums := []int{2,1, 1,2}
	res := rob(nums)
	fmt.Println(res)
}

func rob(nums []int) int {

	r1, r2 := 0, 0

	for i, n := range nums {
		if i%2 == 1 {
			r1 += n
		} else if i%2 == 0 {
			r2 += n
		}
	}
	if r1 > r2 {
		return r1
	} else {
		return r2
	}
}
