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
	fmt.Println(containsDuplicate(nums))
}
func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)
	for i, v := range nums {
		if _, found := mp[v]; found {
			return true
		}
		mp[v] = i
	}
	return false
}
