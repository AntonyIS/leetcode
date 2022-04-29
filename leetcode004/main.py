from typing import List



def maxSubArray(numbers:List[int]) -> int :
    maxSub:int = numbers[0]
    currentSum:int = 0

    for i in numbers:
        if currentSum < 0 :
            currentSum = 0
        
        currentSum += i
        maxSub = max(maxSub, currentSum)

    return maxSub




if __name__ == "__main__":
    numbers = [-2,1,-3,4,-1,2,1,-5,4]
    results = maxSubArray(numbers)
    print(results)