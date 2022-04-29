package main

import "fmt"

/*
2.Best Time to Buy and Sell Stock
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example :

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

*/

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	// present day 	and future day
	present_day, future_day := 0, 1

	// Store maximun profits in this variable maxPro
	var maxPro int

	// loop through days and compare values in present and future day
	for future_day < len(prices) {
		// As long as future day is not greater than the number of days
		// compare values in present_day and future_day
		if prices[present_day] < prices[future_day] {
			// check if the profit is greater than maxPro
			profit := prices[future_day] - prices[present_day]
			if profit > maxPro {
				maxPro = profit
			}

		} else {
			// if price in future is less than price in present, reassign the future value to present day to get the minimum price to buy stock
			present_day = future_day
		}
		// continue moving forward in the future
		future_day += 1
	}

	return maxPro

}
