from typing import List

def maxSubArray(nums:List[int]) -> int :
    # Get the total sum
    sum = nums[0]
    # Define current sum
    curr_sum = 0

    # Iterate through values and get the sum
    for num in nums :
        # if curr_sum sum is negative value
        if curr_sum < 0 :
            curr_sum = 0
        # Calculates the current sum
        curr_sum += num
        # Get the maximum value between curr_sum and sum
        sum = max(curr_sum, sum)
    return sum


if __name__ == "__main__":
    numbers = [-2,1,-3,4,-1,2,1,-5,4]
    results = maxSubArray(numbers)
    print(results)