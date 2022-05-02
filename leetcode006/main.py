from typing import List


def productExceptSelf(nums: List[int]) -> List[int]:
    # create an array results with size of number of items in nums : all items will be 1 i.e [1,1,1, ..... len(nums)]
    result = [1] * len(nums)
    # prefix : product of individual items before them
    prefix = 1
    # Loop through items from left to right : prefixing
    for i in range(len(nums)):
        result[i] = prefix
        prefix *= nums[i]

    # postfix : product of individual items after them
    postfix = 1
    # Loop through items from right to left : postfixing
    for i in range(len(nums)-1,-1,-1):
        result[i] = postfix
        postfix *= nums[i]

    return result


if __name__ == "__main__":
    nums = [-1,1,0,-3,3]
    results = productExceptSelf(nums)
    print(results)

    


