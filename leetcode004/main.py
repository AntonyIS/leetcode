from typing import List



def maxSubArray(numbers:List[int]) -> int :
    mx,currentSum= numbers[0], 0
 

    for i in numbers:
        if currentSum < 0 :
            currentSum = 0
        currentSum += i
        mx = max(mx, currentSum)
    return mx

if __name__ == "__main__":
    numbers = [-2,1,-3,4,-1,2,1,-5,4]
    results = maxSubArray(numbers)
    print(results)