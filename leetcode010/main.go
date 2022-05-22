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
	results := [][]int{}
	sort.Ints(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1

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
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}

	}

	return results

}
