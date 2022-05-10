package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	results := threeSum(nums, target)
	fmt.Println(results)
}

func threeSum(nums []int, target int) [][]int {
	// Final results appended here
	results := [][]int{}
	// Sort the nums slice
	sort.Ints(nums)

	// Loop through sorted int slice
	for i, num := range nums {
		// Check if we are processing the same item twice
		if i > 0 && num == nums[i-1] {
			continue
		}
		// Define left and right pointers
		left, right := i+1, len(nums)-1

		// Loop while left < right
		for left < right {
			threeSum := num + nums[left] + nums[right]
			if threeSum > 0 {
				// Move the right pointer back
				right -= 1
			} else if threeSum < 0 {
				// Move left pointer forward
				left += 1
			} else {
				// Add an array of num, nums[left], nums[right] to final results
				results = append(results, []int{num, nums[left], nums[right]})
				// Move the left pointer forward
				left += 1
				// Ensure that the nums[left] does not repeat itself
				for nums[left] == nums[i-1] && left < right {
					left += 1
				}
			}
		}

	}

	return results

}
