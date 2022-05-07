package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	results := search(nums, 0)
	fmt.Println(results)
}

func search(nums []int, target int) int {
	// Define left and right portion of the array
	left, right := 0, len(nums)-1

	// Loop through the array until left <= right
	for left <= right {
		// Define the midpoint of the array
		midpoint := (left + right) / 2

		// Check if the value at midpoint is the search value
		if nums[midpoint] == target {
			return midpoint
		}
		// Check if search value is in the left portion of the array
		if nums[left] <= nums[midpoint] {
			if target > nums[midpoint] || target < nums[left] {
				// Move the left pointer at midpoint + 1
				left = midpoint + 1
			} else {
				// Move the right pointer at midpoint - 1
				right = midpoint - 1
			}
		} else {
			if target < nums[midpoint] || target > nums[right] {
				// Move the right pointer at midpoint - 1
				right = midpoint - 1
			} else {
				// Move the left pointer at midpoint + 1
				left = midpoint + 1
			}
		}
	}

	return -1

}
