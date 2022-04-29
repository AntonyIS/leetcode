package main

import "fmt"

/*
1. Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

*/

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	// Solution : 2
	// A map to store visited numbers
	visitedValues := make(map[int]int)
	for index, value := range nums {
		// Get the difference between target value and the current value
		difference := target - value

		// Check if the difference exists in the visited value, if yes the second value has been found
		if _, ok := visitedValues[difference]; ok {
			// The complimenting value has been found, return indexes
			return []int{visitedValues[difference], index}
		}
		// Store value as key and index as value
		visitedValues[value] = index
	}
	// Return and empty array if the values that add to target are not found
	return []int{}
}
