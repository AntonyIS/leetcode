package main

import "fmt"

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
