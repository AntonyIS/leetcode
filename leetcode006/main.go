package main

import "fmt"

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	results := productExceptSelf(nums)
	fmt.Println(results)
}

func productExceptSelf(nums []int) []int {
	// create an array with items == 1 if length size of nums
	size := len(nums) // Array size
	results := make([]int, size)

	// prefix : stores product of all items before them
	prefix := 1
	// Loop through items in nums and get the prefixes
	for i := 0; i < len(nums); i++ {
		results[i] = prefix
		prefix *= nums[i]
	}
	// prefix : stores product of all items after them
	postfix := 1
	// Loop through items in nums and get the postfixes

	for i := len(nums) - 1; i >= 0; i-- {
		results[i] = postfix
		postfix *= nums[i]
	}
	return results
}
