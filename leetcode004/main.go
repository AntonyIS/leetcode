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

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maximumSubarray(nums))
}

func maximumSubarray(nums []int) int {
	// A variable to  store the max largest sum
	var largestSum int
	// Loop through the array and get the largest sum
	left, right := 0, 1
	for right < len(nums) {
		// Get the sum of values in index left and right
		sum := nums[left] + nums[right]
		fmt.Println(sum)
		// Update largestsum if it is less than sum

		if largestSum < (sum + largestSum) {
			largestSum = sum + largestSum
			right += 1
		} else {
			left = right
			right += 1
		}

	}
	return largestSum
}
