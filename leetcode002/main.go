package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	mp, l, r := 0, 0, 1

	for r < len(prices) {
		if prices[l] < prices[r] {
			if prices[r]-prices[l] > mp {
				mp = prices[r] - prices[l]
			}
		} else {
			l = r
		}
		r += 1
	}
	return mp

}
