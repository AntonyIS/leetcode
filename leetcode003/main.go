/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
Input: nums = [1,2,3,1]
Output: true
*/

package main

import "fmt"

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
Input: nums = [1,2,3,1]
Output: true
*/

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicates(nums))
}
func containsDuplicates(nums []int) bool {
	// Takes a slice as an argument and returns a boolean
	// Create a map to store values that have been visited
	// visitedValues stores values in array as key and number of duplicates as value
	visitedValues := make(map[int]int)

	// Loop through the slice values
	for _, value := range nums {
		// Check if value exists in visitedValues
		if _, ok := visitedValues[value]; ok {
			// This means this is the second time value appears in the slice
			return true
		}
		// Store values in visitedValues
		visitedValues[value] = 1
	}
	// No value apprear more than once
	return false

}
