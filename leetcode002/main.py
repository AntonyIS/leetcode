from typing import List


def maxProfit(prices:List[int]) -> int:
    mp, l, r = 0,0,1
    while r < len(prices) :
        if prices[l] < prices[r] :
            if prices[r] - prices[l] > mp :
                mp = prices[r] - prices[l]
        else:
            l = r
        r += 1

    return mp
   

if __name__ == "__main__":
    prices:List[int] = [7,1,5,3,6,4]
    results = maxProfit(prices)
    print(results)
