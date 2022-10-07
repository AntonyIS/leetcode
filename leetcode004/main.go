package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maximumSubarray(nums))
}

func maximumSubarray(nums []int) int {
	// Define maxSub to store sum of values in a sub array
	maxSub, currentSum := nums[0], 0
	// Loop through all values in the nums array
	for _, num := range nums {
		// If currentSum is negative, ignore it, it will reduce the current sum
		if currentSum < 0 {
			currentSum = 0
		}
		currentSum += num
		// Get the max sum
		maxSub = max(maxSub, currentSum)
	}

	return maxSub
}

func max(num1, num2 int) int {
	// Returns the max value between two values
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
