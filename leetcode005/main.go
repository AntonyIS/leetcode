package main

import "fmt"

/*## 5. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
A subarray is a contiguous part of an array.
Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
func main() {
	numbers := []int{5, 4, -1, 7, 8}
	res := maxSub(numbers)
	fmt.Println(res)
}

func maxSub(numbers []int) int {
	// Define maxSub , currentSum to store the max sum of subarray and current sum in a loop
	maxSub, currentSum := 0, 0

	for i := 0; i < len(numbers); i++ {
		// Assign the current sum to the currentSum variable
		currentSum = currentSum + numbers[i]

		// Compare currentSum with maxSub
		if currentSum > maxSub {
			maxSub = currentSum
		}
		// Reset currentSum to zero because negative value reduces the max sum
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSub
}
