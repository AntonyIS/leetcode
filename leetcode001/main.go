package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum_1(nums, target))
}

func twoSum_1(nums []int, target int) []int {

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
func twoSum_2(nums []int, target int) []int {
	// Define a map to store values and indexes with key and value being int
	m := make(map[int]int)
	// loop through key and values in nums | index and value
	for i, num := range nums {
		// Check if index i is in map m
		if v, found := m[target-num]; found {
			// Return int array if i(v) is in map
			return []int{v, i}
		}
		// Add num as key and i as value in map m
		m[num] = i
	}
	// Return nil if no corresponding value found
	return nil

}
