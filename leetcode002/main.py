from typing import List


def bestTimeToBuySellStock(prices:List[int]) -> int:
    maxProfit = 0 # Stores the maximum profit as we treverse the prices
    today, tomorrow = 0, 1 #Defines stock prices today and tomorrow

    # Loop through the prices array
    while tomorrow < len(prices):

        if prices[today] < prices[tomorrow] :
            profit = prices[tomorrow] - prices[today] # Get price difference
            # Check if profit is greater than maxProfit
            if profit > maxProfit :
                # Assign the profit to be maxProfit
                maxProfit = profit
        else:
            # Reassign today to tomorrow
            today = tomorrow
        # For each loop move to the next day
        tomorrow += 1

    return maxProfit

if __name__ == "__main__":
    prices:List[int] = [7,1,5,3,6,4]
    results = bestTimeToBuySellStock(prices)
    print(results)
