package main

import (
	"fmt"
)

func main() {

	nums := []int{-4, -3, -2}
	maxProduct := maxProduct(nums)

	fmt.Println(maxProduct)
}

func maxProduct(nums []int) int {
	// Define maxsub array to be the first item in the array
	response := max(nums)

	// Define the currentMax, currentMin
	currentMax, currentMin := 1, 1

	// Loop through item in the number slice
	for _, num := range nums {
		// Check if currentProduct is less than 0
		if num == 0 {

			currentMax, currentMin = 1, 1
		}

		tempMax := num * currentMax
		currentMax = max([]int{num * currentMax, num * currentMin, num})
		currentMin = min([]int{tempMax, num * currentMin, num})
		response = max([]int{response, currentMax})

	}

	return response
}
func max(nums []int) int {
	var maxValue int = -1000000
	if len(nums) == 1 {
		maxValue = nums[0]
		return maxValue
	}
	for _, value := range nums {

		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
func min(nums []int) int {
	var minValue int = 1000000
	if len(nums) == 1 {
		minValue = nums[0]
		return minValue
	}
	for _, value := range nums {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}
