package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(nums)
	fmt.Println(res)
}

func maxArea(nums []int) int {
	// Define the maxarea
	maxArea := nums[0]

	left, right := 0, len(nums)-1

	for left < right {
		length := right - left
		height := min(nums[left], nums[right])
		area := length * height
		maxArea = max(maxArea, area)

		if nums[left] < nums[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return maxArea

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
