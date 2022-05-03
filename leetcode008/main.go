package main

import "fmt"

func main() {
	nums := []int{2, 1}
	minValue := findMin(nums)
	fmt.Println(minValue)
}

func findMin(nums []int) int {
	minValue := nums[0]
	left, right := 0, len(nums)-1

	for left <= right {
		// Check if item on the left is less that item on the right
		if nums[left] < nums[right] { // Sorted
			// Get the min value between minValue and item on the left
			minValue = min(minValue, nums[left])
			break
		}
		// Get the mid point of the array
		mid := (left + right) / 2

		// Get the min value between minValue and value at mid point
		minValue = min(minValue, nums[mid])

		if nums[left] > nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return minValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
