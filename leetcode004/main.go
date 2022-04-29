/*
4.Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
A subarray is a contiguous part of an array.
Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
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
	maxSub := nums[0]
	// Define currentSum to store sum as we traverse the nums array
	currentSum := 0
	// Loop through all values in the nums array
	for _, i := range nums {
		// If currentSum is negative, ignore it, it will reduce the current sum
		if currentSum < 0 {
			// Reassign currentSum to 0
			currentSum = 0
		}
		// Get the currentSum
		currentSum += i
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
