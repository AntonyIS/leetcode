from typing import List

def maxProduct(nums:List[int]) -> int :
    # Define the maximum value in array e.g [10]
    maxProd = max(nums)
    # Define max and min current values
    currentMax, currentMin = 1,1

    # Loop through values in nums and key track of the max, min
    for num in nums :
        # Ignore zero values in nums
        tempMax = currentMax
        # Keep track of both currentMin and currentMax in cases where nums has negative values
        currentMax = max(num,num* tempMax, num * currentMin)
        currentMin = min(num,num* tempMax,num * currentMin)
        maxProd = max(currentMax, maxProd)
        
    return maxProd

if __name__ == "__main__":
    nums:List[int] = [-1, -2, -9, -6]
    maxValue = maxProduct(nums)
    print(maxValue)