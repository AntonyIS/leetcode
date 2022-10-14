package main

import "fmt"

func main() {
	nums := []int{0, 1}
	res := missingNumber(nums)
	fmt.Println(res)
}

func missingNumber(nums []int) int {
	res := len(nums)

	for i, _ := range nums {
		res += (i - nums[i])
	}
	return res
}
