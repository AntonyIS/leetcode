from typing import List

def maxProduct(nums:List[int]) -> int :
    # Define the maximum value in array e.g [10]
    maxProd = max(nums)
    
    # Define max and min current values
    currentMax, currentMin = 1,1

    # Loop through values in nums and key track of the max, min
    for num in nums :
        # Ignore zero values in nums
        if num == 0 :
            currentMax, currentMin = 1,1
            continue

        tempMax = num * currentMax
        # Keep track of both currentMin and currentMax in cases where nums has negative values
        currentMax = max(num * currentMax, num * currentMin, num)
        currentMin = min(tempMax, num * currentMin, num)
        maxProd = max(currentMax, maxProd)
        
    return maxProd

if __name__ == "__main__":
    nums:List[int] = [-4, -3, -2]
    maxValue = maxProduct(nums)
    print(maxValue)