package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, -2, -9, -6}
	maxProduct := maxProduct(nums)
	fmt.Println(maxProduct)
}

func maxProduct(nums []int) int {
	currentMax, currentMin, response := nums[0], nums[0], nums[0]
	// Loop through all the items in nums starting at index 1, index 0 already covered
	for i := 1; i < len(nums); i++ {
		// Define a temperary current max
		tmpMax := currentMax

		// Get currentMax from the current item in the loop (nums[i]), (nums[i]*tmpMax) and currentMin * nums[i]
		currentMax = max(nums[i], max(tmpMax*nums[i], currentMin*nums[i]))
		currentMin = min(nums[i], min(tmpMax*nums[i], currentMin*nums[i]))
		if currentMax > response {
			response = currentMax
		}
	}
	return response
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
